package comparisons

import (
	"jvmgo/rtda"
	"jvmgo/instruction"
)

// Branch if reference comparison succeeds
type IF_ACMPEQ struct{ instruction.BranchInstruction }
type IF_ACMPNE struct{ instruction.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		instruction.Branch(frame, self.Offset)
	}
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		instruction.Branch(frame, self.Offset)
	}
}

func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}

