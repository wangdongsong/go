package rtda

//Java虚拟机栈
type Stack struct {
	//栈最大容量
	maxSize uint
	size uint
	_top *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

//压栈
func (self *Stack) push(frame *Frame)  {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if self._top != nil {
		frame.lower = self._top
	}

	self._top = frame
	self.size++
}

//出栈
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}

	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--

	return top
}

//返回栈顶
func (self *Stack) top()  *Frame{
	if self._top == nil {
		panic("jvm stack is empty")
	}

	return self._top

}

