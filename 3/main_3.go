package main

import "fmt"

type StringIntMap map[string]int

func main() {
	var stringMap StringIntMap
	stringMap = make(StringIntMap)

	stringMap.Add("key1", 1)
	stringMap.Add("key2", 2)
	stringMap.Remove("key1")

	stringMapCopy := stringMap.Copy()
	fmt.Printf("stringMapCopy is %v\n", stringMapCopy)

	stringMapExists := stringMap.Exists("key2")
	fmt.Println(stringMapExists)

	stringMapValue, stringMapExists := stringMap.Get("key2")
	fmt.Println(stringMapValue, stringMapExists)
}

func (s StringIntMap) Add(key string, value int) {
	s[key] = value
}

func (s StringIntMap) Remove(key string) {
	delete(s, key)
}

func (s StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int, len(s))
	for k, v := range s {
		copyMap[k] = v
	}
	return copyMap
}

func (s StringIntMap) Exists(key string) bool {
	_, ok := s[key]
	return ok
}

func (s StringIntMap) Get(key string) (int, bool) {
	value, ok := s[key]
	return value, ok
}
