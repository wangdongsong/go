package math

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"math"
)

//算术运算——求余指令

//double
type DREM struct {
	instruction.NoOperandsInstruction
}

//float
type FREM struct {
	instruction.NoOperandsInstruction
}

//int
type IREM struct {
	instruction.NoOperandsInstruction
}

//long
type LREM struct {
	instruction.NoOperandsInstruction
}

//先从操作数栈中弹出两个int变量，求余，然后将结果推入操作数栈，对int和long变量做除法和求余运算时，是有可能抛出ArithmeticException
func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	if v2 == 0 {
		panic("java.lang.ArithemticException: /by zero")
	}

	result := v1 % v2
	stack.PushInt(result)
}

//double
func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2) // todo
	stack.PushDouble(result)
}

//float
func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2))) // todo
	stack.PushFloat(result)
}

//long
func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v2 == 0 {
		frame.Thread().ThrowDivByZero()
	} else {
		result := v1 % v2
		stack.PushLong(result)
	}
}