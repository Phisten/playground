package main

import (
	"fmt"
	"math/rand"
)

func main() {
	obj := Constructor()
	obj.Insert(1)
	obj.Insert(2)
	obj.Insert(3)
	obj.Remove(2)
	obj.Remove(2)
	obj.PRINTFF()
	fmt.Printf("----\n")

	obj.GetRandom()
	obj.GetRandom()
	obj.GetRandom()
	obj.GetRandom()
	obj.GetRandom()
}

// 爬文看到golang取出map的keys時是隨機的, 但是leetcode似乎有指定rand種子, 不使用rand輸出就會錯？

type RandomizedSet struct {
	ht map[int]bool
	n  int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		ht: map[int]bool{},
		n:  0,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if this.ht[val] {
		return false
	}
	this.n += 1
	this.ht[val] = true
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if this.ht[val] {
		delete(this.ht, val)
		this.n -= 1
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	rnd := rand.Intn(this.n)
	count := 0
	for key := range this.ht {
		if count == rnd {
			return key
		}
		count++
	}
	return -1
}

// --

func (this *RandomizedSet) PRINTFF() {
	for key, value := range this.ht {
		fmt.Println(key, " ", value)
	}
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
