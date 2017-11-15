package constants

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type ACONST_NULL struct {
	instruction.NoOperandsInstruction
}

type DCONST_0 struct {
	instruction.NoOperandsInstruction
}

type DCONST_1 struct {
	instruction.NoOperandsInstruction
}

type FCONST_0 struct {
	instruction.NoOperandsInstruction
}

type FCONST_1 struct {
	instruction.NoOperandsInstruction
}

type FCONST_2 struct {
	instruction.NoOperandsInstruction
}

type ICONST_M1 struct {
	instruction.NoOperandsInstruction
}

type ICONST_0 struct {
	instruction.NoOperandsInstruction
}

type ICONST_1 struct {
	instruction.NoOperandsInstruction
}

type ICONST_2 struct {
	instruction.NoOperandsInstruction
}

type ICONST_3 struct {
	instruction.NoOperandsInstruction
}

type ICONST_4 struct {
	instruction.NoOperandsInstruction
}

type ICONST_5 struct {
	instruction.NoOperandsInstruction
}

type LCONST_0 struct {
	instruction.NoOperandsInstruction
}

type LCONST_1 struct {
	instruction.NoOperandsInstruction
}

func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushNull()
}

func (self *DCONST_0) Execute(frame rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

//iconst_m1指令把int型-1推入操作数栈顶
func (self *ICONST_M1) Execute(frame rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}
