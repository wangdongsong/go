package constants

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

//bipush指令从操作数中获取一个byte型整数，扩展成int型，然后推入栈顶。
//sipush指令从操作数中获取一个short型整数，扩展成int型，然后推入栈顶

type BIPUSH struct {
	val int8
}

type SIPUSH struct {
	val int16
}

func (self *BIPUSH) FetchOperands(reader *instruction.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
