package main

import "container/list"

const base = 769

func (s *MyHashSet) hash(key int) int {
	return key % base
}

type MyHashSet struct {
	data []list.List
}

func Constructor() MyHashSet {
	return MyHashSet{make([]list.List, base)}
}

func (s *MyHashSet) Add(key int) {
	if !s.Contains(key) {
		h := s.hash(key)
		s.data[h].PushBack(key)
	}
}

func (s *MyHashSet) Remove(key int) {
	h := s.hash(key)
	for elem := s.data[h].Front(); elem != nil; elem = elem.Next() {
		if elem.Value.(int) == key {
			s.data[h].Remove(elem)
		}
	}
}

func (s *MyHashSet) Contains(key int) bool {
	h := s.hash(key)
	for elem := s.data[h].Front(); elem != nil; elem = elem.Next() {
		if elem.Value.(int) == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

func main() {
	hashSet := Constructor()
	hashSet.Add(1)
	hashSet.Add(2)
	println(hashSet.Contains(1))
	println(hashSet.Contains(2))
	hashSet.Remove(1)
	hashSet.Remove(2)
	println(hashSet.Contains(1))
	println(hashSet.Contains(2))

}
