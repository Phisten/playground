package leetcode

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
