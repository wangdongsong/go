package constants

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

//加载指令
//加载指令从变量表获取变量，然后推入操作数栈顶，按所操作变量的类型可分为6类
//1、aload加载操作引用类型
//2、dload加载double类型
//3、fload操作float变量
//4、iload系统操作int变量
//5、lload操作long变量
//6、xaload操作数组

type ILOAD struct {
	instruction.Index8Instruction
}

type ILOAD_0 struct {
	instruction.NoOperandsInstruction
}

type ILOAD_1 struct {
	instruction.NoOperandsInstruction
}

type ILOAD_2 struct {
	instruction.NoOperandsInstruction
}

type ILOAD_3 struct {
	instruction.NoOperandsInstruction
}

func _iload(frame *rtda.Frame, index uint) {
	val :=frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(self.Index))
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}