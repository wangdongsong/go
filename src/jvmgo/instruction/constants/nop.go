package constants

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type NOP struct {
	instruction.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	//nothing todo
}
