package stores

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

// Store reference into local variable
type ASTORE struct{ instruction.Index8Instruction }

func (self *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, uint(self.Index))
}

type ASTORE_0 struct{ instruction.NoOperandsInstruction }

func (self *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

type ASTORE_1 struct{ instruction.NoOperandsInstruction }

func (self *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

type ASTORE_2 struct{ instruction.NoOperandsInstruction }

func (self *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

type ASTORE_3 struct{ instruction.NoOperandsInstruction }

func (self *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}

func _astore(frame *rtda.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}
