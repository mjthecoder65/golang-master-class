package dsa

// URL: https://leetcode.com/problems/design-hashmap/

// Title: Design HashMap

const MAP_CAPACITY = 1000000

type HashMap struct {
	elements []int
}

func Hash(key int) int {
	return key % MAP_CAPACITY
}

func (hashMap *HashMap) Put(key int, value int) {
	key = Hash(key)
	hashMap.elements[key] = value
}

func (hashMap *HashMap) Get(key int) int {
	key = Hash(key)
	return hashMap.elements[key]
}

func (hashMap *HashMap) Remove(key int) {
	key = Hash(key)
	hashMap.elements[key] = -1
}

func Constructor() HashMap {
	hashMap := HashMap{make([]int, MAP_CAPACITY)}
	for index := range hashMap.elements {
		hashMap.elements[index] = -1
	}
	return hashMap
}
