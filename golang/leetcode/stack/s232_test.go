package leetcode

import (
	"testing"
)

func Test_s232(t *testing.T) {

	myQueue := Constructor()
	myQueue.Push(1) // queue is: [1]
	myQueue.Push(2) // queue is: [1, 2] (leftmost is front of the queue)
	myQueue.Peek()  // return 1
	myQueue.Pop()   // return 1, queue is [2]
	myQueue.Empty() // return false
}

type Stack struct {
	data []int
}

func (this *Stack) Push(x int) {
	this.data = append(this.data, x)
}
func (this *Stack) Pop() int {
	pop := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return pop
}
func (this *Stack) Peek() int {
	return this.data[len(this.data)-1]
}
func (this *Stack) Empty() bool {
	return len(this.data) == 0
}

type MyQueue struct {
	normal  Stack
	reverse Stack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	for !this.reverse.Empty() {
		this.normal.Push(this.reverse.Pop())
	}
	this.normal.Push(x)
}

func (this *MyQueue) Pop() int {
	for !this.normal.Empty() {
		this.reverse.Push(this.normal.Pop())
	}
	return this.reverse.Pop()
}

func (this *MyQueue) Peek() int {
	for !this.normal.Empty() {
		this.reverse.Push(this.normal.Pop())
	}
	return this.reverse.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.normal.Empty() && this.reverse.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
