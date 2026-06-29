package main

import (
	"testing"
	"time"
)

func Test_generator_ReturnsValue(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	gen := generator(done)

	select {
	case v, ok := <-gen:
		if !ok {
			t.Fatal("channel closed unexpectedly")
		}
		if v < 0 || v >= 100 {
			t.Errorf("value %d out of range [0, 100)", v)
		}
	case <-time.After(time.Second):
		t.Fatal("timeout: no value received")
	}
}

func Test_generator_ReturnsMultipleValues(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	gen := generator(done)

	for i := 0; i < 10; i++ {
		select {
		case v, ok := <-gen:
			if !ok {
				t.Fatalf("channel closed at iteration %d", i)
			}
			if v < 0 || v >= 100 {
				t.Errorf("iteration %d: value %d out of range [0, 100)", i, v)
			}
		case <-time.After(time.Second):
			t.Fatalf("timeout at iteration %d", i)
		}
	}
}

func Test_generator_ClosesChannelAfterDone(t *testing.T) {
	done := make(chan struct{})
	gen := generator(done)

	// читаем одно значение чтобы горутина запустилась
	<-gen

	close(done)

	// ждём пока канал закроется
	select {
	case _, ok := <-gen:
		if ok {
			// канал ещё открыт — читаем до закрытия
			for range gen {
			}
		}
	case <-time.After(time.Second):
		t.Fatal("timeout: channel was not closed after done signal")
	}
}

func Test_generator_ValuesInRange(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	gen := generator(done)

	for i := 0; i < 1000; i++ {
		select {
		case v := <-gen:
			if v < 0 || v >= 100 {
				t.Errorf("value %d out of range [0, 100)", v)
			}
		case <-time.After(time.Second):
			t.Fatal("timeout")
		}
	}
}

func Test_generator_IsUnbuffered(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	gen := generator(done)

	if cap(gen) != 0 {
		t.Errorf("expected unbuffered channel (cap=0), got cap=%d", cap(gen))
	}
}

func Test_generator_MultipleInstances(t *testing.T) {
	done1 := make(chan struct{})
	done2 := make(chan struct{})
	defer close(done1)
	defer close(done2)

	gen1 := generator(done1)
	gen2 := generator(done2)

	v1 := <-gen1
	v2 := <-gen2

	if v1 < 0 || v1 >= 100 {
		t.Errorf("gen1: value %d out of range", v1)
	}
	if v2 < 0 || v2 >= 100 {
		t.Errorf("gen2: value %d out of range", v2)
	}
}

func Test_generator_StopsAfterDone(t *testing.T) {
	done := make(chan struct{})
	gen := generator(done)

	<-gen
	close(done)

	// даём горутине время завершиться
	time.Sleep(10 * time.Millisecond)

	// канал должен быть закрыт — дальнейшее чтение вернёт zero value и ok=false
	select {
	case v, ok := <-gen:
		if ok {
			t.Logf("got value %d before channel closed", v)
		}
		// ok=false — канал закрыт, всё корректно
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for channel to close")
	}
}
