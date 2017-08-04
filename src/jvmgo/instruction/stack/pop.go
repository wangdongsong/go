package stack

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

//pop， pop2指令

type POP struct {
	instruction.NoOperandsInstruction
}

type POP2 struct {
	instruction.NoOperandsInstruction
}

//pop指令把栈顶变量弹出
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

//pop指令只能弹出int、float等占用一个操作数栈位置的变量，double和long变量在操作数栈中占据两个位置
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}