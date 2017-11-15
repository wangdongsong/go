package basic

import (
	"fmt"
	"strconv"
)

func main() {

	a := 21
	var b int = 10
	var c int

	//operator(c, a, b)

	fmt.Println(strconv.FormatBool(a == b) + " a ")
	fmt.Println(a == b, "b ")
	println(c)
	println(a < b)
	println(a > b)
	println(a <= b)
	println(a >= b)

	println("--------------")

	var boolA, boolB bool
	println("line one boolA && boolB = ", boolA && boolB)

	boolA = true
	boolB = true
	println("line two boolA && boolB = ", boolA && boolB)

	boolB = false
	println("line three boolA || boolB = ", boolA || boolB)

	println("line four !boolB = ", !boolB)

	println("---------位运算--------")
	println("0 & 0 = ", 0&0)
	println("0 | 0 = ", 0|0)
	println("0 ^ 0 = ", 0^0)

	println("0 & 1 = ", 0&1)
	println("0 | 1 = ", 0|1)
	println("0 ^ 1 = ", 0^1)

	println("1 & 0 = ", 1&0)
	println("1 | 0 = ", 1|0)
	println("1 ^ 0 = ", 1^0)

	println("1 & 1 = ", 1&1)
	println("1 | 1 = ", 1|1)
	println("1 ^ 1 = ", 1^1)

	a = 10
	var ptr *int
	println(&a)
	ptr = &a
	println(*ptr)

}

func operator(c int, a int, b int) {
	c = a + b
	fmt.Printf("first line - c value = %d\n", c)
	c = a - b
	fmt.Printf("second line -c value = %d\n", c)
	c = a * b
	fmt.Printf("third line -c value = %d\n", c)
	c = a / b
	fmt.Printf("fouth line -c value = %d\n", c)
	c = a % b
	fmt.Printf("fifth line -c value = %d\n", c)
	a++
	fmt.Printf("sixth line -c value= %d\n", a)
	a--
	fmt.Printf("seventh line -c value= %d\n", a)
}
