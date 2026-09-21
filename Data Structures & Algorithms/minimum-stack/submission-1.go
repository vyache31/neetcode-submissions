type MinStack struct {
	stack  []int
	minMap map[int]int
}

func Constructor() MinStack {
	return MinStack{
		stack:  []int{},
		minMap: map[int]int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.minMap[0] = val
	} else {
		if val < this.minMap[len(this.stack)-1] {
			this.minMap[len(this.stack)] = val
		} else {
			this.minMap[len(this.stack)] = this.minMap[len(this.stack)-1]
		}
	}
	this.stack = append(this.stack, val)

}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minMap[len(this.stack)-1]
}