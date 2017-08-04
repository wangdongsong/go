package stores

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

//与加载指令相反，存储指令是把变量从操作数栈顶弹出，然后存入局部变量表。

type LSTORE struct {
	instruction.Index8Instruction
}

type LSTORE_0 struct {
	instruction.NoOperandsInstruction
}

type LSTORE_1 struct {
	instruction.NoOperandsInstruction
}

type LSTORE_2 struct {
	instruction.NoOperandsInstruction
}

type LSTORE_3 struct {
	instruction.NoOperandsInstruction
}

func _lstore(frame *rtda.Frame, index uint){
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (self *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(self.Index))
}

func (self *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}