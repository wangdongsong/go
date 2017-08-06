package comparisons

import (
	"jvmgo/rtda"
	"jvmgo/instruction"
)

// Branch if int comparison succeeds
type IF_ICMPEQ struct{ instruction.BranchInstruction }
type IF_ICMPNE struct{ instruction.BranchInstruction }
type IF_ICMPGE struct{ instruction.BranchInstruction }
type IF_ICMPLE struct{ instruction.BranchInstruction }
type IF_ICMPLT struct{ instruction.BranchInstruction }
type IF_ICMPGT struct{ instruction.BranchInstruction }

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		instruction.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		instruction.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		instruction.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		instruction.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		instruction.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		instruction.Branch(frame, self.Offset)
	}
}

func _icmpPop(frame *rtda.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}
