package math

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

//布尔运算指令

type IAND struct {
	instruction.NoOperandsInstruction
}

type LAND struct {
	instruction.NoOperandsInstruction
}

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

func (self *LAND) Execute(frmae *rtda.Frame) {
	stack := frmae.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
