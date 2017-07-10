package main

import "fmt"

func main()  {
	var numbers = make([] int, 3, 5)
	fmt.Printf("len = %d cap = %d slice=%v\n", len(numbers), cap(numbers), numbers)

	var numbers1 []int
	fmt.Printf("len = %d cap = %d slice=%v\n", len(numbers1), cap(numbers1), numbers1)

	if numbers1 == nil {
		fmt.Printf("Slice is nil")
	}

	subSlice()
}

func subSlice()  {

	//创建切片
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	//打印原始切片
	fmt.Println("numbers == ", numbers)

	//打印子切片从索引1到索引4
	fmt.Println("number[1:4] == ", numbers[1:4])

	fmt.Println("number[:3] ==", numbers[:3] )

	fmt.Println("numbers[4:] == ", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	numbers2 := numbers[:2]
	printSlice(numbers2)

	numbers3 := numbers[2:5]
	printSlice(numbers3)

}

func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice=%v\n", len(x), cap(x), x)
}
