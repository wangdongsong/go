package conversions

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

//d2x系统指令把double变量强制转换成其它类型

type D2F struct {
	instruction.NoOperandsInstruction
}

type D2I struct {
	instruction.NoOperandsInstruction
}

type D2L struct {
	instruction.NoOperandsInstruction
}

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}

func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}
