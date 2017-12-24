package stack

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type SWAP struct {
	instruction.NoOperandsInstruction
}

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)

}
