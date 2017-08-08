package control

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

// Jump subroutine
type JSR struct{ instruction.BranchInstruction }

func (self *JSR) Execute(frame *rtda.Frame) {
	panic("todo")
}

// Jump subroutine (wide index)
type JSR_W struct {
	offset int
}

func (self *JSR_W) FetchOperands(reader *instruction.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *JSR_W) Execute(frame *rtda.Frame) {
	panic("todo")
}

// Return from subroutine
type RET struct{ instruction.Index8Instruction }

func (self *RET) Execute(frame *rtda.Frame) {
	panic("todo")
}
