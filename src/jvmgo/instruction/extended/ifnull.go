package extended

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

// Branch if reference is null
type IFNULL struct{ instruction.BranchInstruction }

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		instruction.Branch(frame, self.Offset)
	}
}

// Branch if reference not null
type IFNONNULL struct{ instruction.BranchInstruction }

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		instruction.Branch(frame, self.Offset)
	}
}
