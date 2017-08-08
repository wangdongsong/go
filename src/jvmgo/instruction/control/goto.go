package control

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type GOTO struct {
	instruction.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	instruction.Branch(frame, self.Offset)
}