package basic

import "fmt"

/*
 * 条件表达式
 */
func main() {
	var a = 10
	var b = 11
	if a < 10 {
		println("a <= 10")
	} else {
		println("a >= 10")
	}

	if a == 10 {
		if a < b  {
			println("a < b")
		}
	}


	var grade string = "B"
	var mark int = 90

	switch mark {
		case 90 : grade = "A"
		case 80 : grade = "B"
		case 50, 60, 70 : grade = "C"
		default:
			grade = "D"
	}

	switch  {
	case grade == "A":
		fmt.Println("优秀")
	case grade == "B":
		fmt.Println("良好")
	case grade == "D":
		fmt.Println("及格")
	case grade == "F":
		fmt.Println("不及格")
	default:
		fmt.Println("差")
	}

	//switch语句可以判断某个interface变量中实际存储的变量类型
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Println("x type of %t", i)
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("undefined")
	}

}
