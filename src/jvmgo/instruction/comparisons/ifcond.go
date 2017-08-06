package comparisons

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type IFEQ struct {	instruction.BranchInstruction}

type IFLT struct {	instruction.BranchInstruction}

type IFNE struct{ instruction.BranchInstruction }

type IFGT struct{ instruction.BranchInstruction }

type IFGE struct{ instruction.BranchInstruction }

type IFLE struct{ instruction.BranchInstruction }

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		instruction.Branch(frame, self.Offset)
	}
}


func (self *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		instruction.Branch(frame, self.Offset)
	}
}


func (self *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		instruction.Branch(frame, self.Offset)
	}
}



func (self *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}



func (self *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}



func (self *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}

