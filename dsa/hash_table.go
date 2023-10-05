package dsa

// DAY 4
// URL: https://leetcode.com/problems/design-hashmap/

// Title: Design HashMap

const MAP_CAPACITY = 1000000

type Element struct {
	value int
}

type HashMap struct {
	elements []*Element
}

func Hash(key int) int {
	return key % MAP_CAPACITY
}

func Constructor() HashMap {
	hashMap := HashMap{
		elements: make([]*Element, MAP_CAPACITY),
	}
	return hashMap
}

func (hashMap *HashMap) Put(key int, value int) {
	key = Hash(key)
	hashMap.elements[key] = &Element{value: value}
}

func (hashMap *HashMap) Get(key int) int {
	key = Hash(key)
	element := hashMap.elements[key]
	if element == nil {
		return -1
	}
	return element.value
}

func (hashMap *HashMap) Remove(key int) {
	key = Hash(key)
	hashMap.elements[key] = nil
}
