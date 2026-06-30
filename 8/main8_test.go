package main

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Базовый сценарий: несколько воркеров, Wait дожидается всех.
func TestCustomerWaitGroup_Basic(t *testing.T) {
	wg := NewCustomerWaitGroup()

	var counter int32
	const workers = 10

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()

	if got := atomic.LoadInt32(&counter); got != workers {
		t.Fatalf("ожидали %d завершённых воркеров, получили %d", workers, got)
	}
}

// Wait не должен возвращаться раньше времени, пока есть незавершённые задачи.
func TestCustomerWaitGroup_WaitBlocksUntilDone(t *testing.T) {
	wg := NewCustomerWaitGroup()
	wg.Add(1)

	done := make(chan struct{})

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		t.Fatal("Wait() вернулся до вызова Done()")
	case <-time.After(100 * time.Millisecond):
		// ожидаемое поведение — Wait всё ещё блокирует
	}

	wg.Done()

	select {
	case <-done:
		// ок, разблокировался после Done()
	case <-time.After(time.Second):
		t.Fatal("Wait() не разблокировался после Done()")
	}
}

// Несколько горутин вызывают Wait одновременно — все должны разблокироваться.
// Это ключевой тест на семантику close() vs send().
func TestCustomerWaitGroup_MultipleWaiters(t *testing.T) {
	wg := NewCustomerWaitGroup()
	wg.Add(1)

	const waiters = 5
	var wgTest sync.WaitGroup
	results := make([]bool, waiters)

	for i := 0; i < waiters; i++ {
		wgTest.Add(1)
		go func(idx int) {
			defer wgTest.Done()
			wg.Wait()
			results[idx] = true
		}(i)
	}

	time.Sleep(100 * time.Millisecond) // даём горутинам встать в Wait()
	wg.Done()

	finished := make(chan struct{})
	go func() {
		wgTest.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case <-time.After(time.Second):
		t.Fatal("не все ожидающие горутины разблокировались")
	}

	for i, ok := range results {
		if !ok {
			t.Errorf("waiter %d не разблокировался", i)
		}
	}
}

// Add с суммарным значением > 1 за один вызов.
func TestCustomerWaitGroup_AddMultiple(t *testing.T) {
	wg := NewCustomerWaitGroup()
	wg.Add(3)

	var counter int32
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()

	if got := atomic.LoadInt32(&counter); got != 3 {
		t.Fatalf("ожидали 3, получили %d", got)
	}
}

// Done() без предварительного Add() должен паниковать (idx уходит в -1).
func TestCustomerWaitGroup_DoneWithoutAdd_Panics(t *testing.T) {
	wg := NewCustomerWaitGroup()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("ожидали панику при Done() без Add(), паники не было")
		}
	}()

	wg.Done()
}

// Лишний Done() (больше вызовов Done, чем Add) тоже должен паниковать.
func TestCustomerWaitGroup_ExtraDone_Panics(t *testing.T) {
	wg := NewCustomerWaitGroup()
	wg.Add(1)
	wg.Done() // idx = 0, sem закрывается, ок

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("ожидали панику при лишнем Done(), паники не было")
		}
	}()

	wg.Done() // idx = -1, должна быть паника
}

// Проверка гонки: конкурентные Add/Done не должны портить счётчик.
// Запускать с go test -race.
func TestCustomerWaitGroup_ConcurrentAddDone_Race(t *testing.T) {
	wg := NewCustomerWaitGroup()

	const n = 100
	wg.Add(n)

	for i := 0; i < n; i++ {
		go wg.Done()
	}

	select {
	case <-func() chan struct{} {
		ch := make(chan struct{})
		go func() {
			wg.Wait()
			close(ch)
		}()
		return ch
	}():
	case <-time.After(2 * time.Second):
		t.Fatal("Wait() завис при конкурентных Done()")
	}
}

// ВАЖНОЕ ОГРАНИЧЕНИЕ: эта WaitGroup одноразовая.
// После закрытия sem повторный цикл Add -> Wait сломается,
// потому что Wait() читает из уже закрытого канала и возвращается мгновенно,
// даже если реальная работа ещё не завершена.
func TestCustomerWaitGroup_NotReusable_DemonstratesLimitation(t *testing.T) {
	wg := NewCustomerWaitGroup()

	wg.Add(1)
	wg.Done() // idx = 0 -> close(sem)

	// Второй "раунд" использования той же wg
	wg.Add(1) // idx = 1, но sem уже закрыт навсегда

	waitReturned := make(chan struct{})
	go func() {
		wg.Wait() // должен был бы блокироваться, но НЕ блокируется
		close(waitReturned)
	}()

	select {
	case <-waitReturned:
		// Подтверждаем баг: Wait() вернулся, хотя Done() для второго Add ещё не вызван
		t.Log("подтверждено: Wait() вернулся преждевременно — WaitGroup не переиспользуема")
	case <-time.After(200 * time.Millisecond):
		t.Fatal("неожиданно: Wait() заблокировался — поведение изменилось?")
	}
}
