package basic

import "fmt"

func main()  {

	var b int = 15
	var a int

	numbers :=  [6]int{1, 3, 5, 7}

	for a := 0; a < 10; a++ {
		fmt.Printf("a的值为：%d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a value is %d\n", a)
	}

	for i, x := range numbers{
		fmt.Printf("index %d key x value is %d\n", i, x)
	}

	//嵌套循环
	var i, j int
	for i = 2; i < 100; i++{
		for j = 2; j < (i/j); j++{
			if i % j == 0 {
				break;
			}
		}

		if j > (i / j ){
			fmt.Printf("%d 是素数\n", i )
		}
	}

}
