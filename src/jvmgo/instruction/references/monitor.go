package references

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

// Enter monitor for object
type MONITOR_ENTER struct{ instruction.NoOperandsInstruction }

func (self *MONITOR_ENTER) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		frame.RevertNextPC()
		thread.ThrowNPE()
	} else {
		ref.Monitor().Enter(thread)
	}
}

// Exit monitor for object
type MONITOR_EXIT struct{ instruction.NoOperandsInstruction }

func (self *MONITOR_EXIT) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		frame.RevertNextPC()
		thread.ThrowNPE()
	} else {
		ref.Monitor().Exit(thread)
	}
}
