package control

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

// Return void from method
type RETURN struct{ instruction.NoOperandsInstruction }

// Return reference from method
type ARETURN struct{ instruction.NoOperandsInstruction }

// Return double from method
type DRETURN struct{ instruction.NoOperandsInstruction }

// Return float from method
type FRETURN struct{ instruction.NoOperandsInstruction }

// Return double from method
type LRETURN struct{ instruction.NoOperandsInstruction }

func (self *RETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	thread.PopFrame()
}


func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}


func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}



func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

// Return int from method
type IRETURN struct{ instruction.NoOperandsInstruction }

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}


func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}
