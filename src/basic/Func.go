package basic

import (
	"fmt"
	"math"
)

var globalVar = 100

func main() {
	fmt.Printf("global var is %d \n", globalVar)
	baseFunc()

	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
}

type Circle struct {
	radius float64
}

/**
 * 方法就是一个包含了接受者的函数
 */
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

/**
 * 闭包，函数返回匿名函数
 */
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/**
 * 基本函数，函数作为值
 */
func baseFunc() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
}
