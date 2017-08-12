package references

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

// Check whether object is of given type
type CHECK_CAST struct {
	instruction.Index16Instruction
	class *heap.Class
}

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	if self.class == nil {
		cp := frame.Method().Class().ConstantPool()
		kClass := cp.GetConstant(self.Index).(*heap.ConstantClass)
		self.class = kClass.Class()
	}

	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)

	if ref == nil {
		return
	}

	if !ref.IsInstanceOf(self.class) {
		frame.Thread().ThrowClassCastException(ref.Class(), self.class)
	}
}
