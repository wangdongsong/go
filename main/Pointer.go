package main

import (
	"fmt"
)

//定义常量
const MAX  int = 3

func main()  {

	//basePointerFunc()

	//指针数组
	pointArray()

}

func pointArray()  {
	a := []int{10, 100, 200}

	var i int

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i,  a[i])
	}

	/*
	 * 使用指针访问数组
	 */
	var ptr [MAX] *int
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

func basePointerFunc() {
	var c = 0

	fmt.Printf("变量地址：%x\n", &c)

	var a int = 20
	var ip *int

	ip = &a

	//a变量地址
	fmt.Printf("a变量的地址：%x\n", &a)
	//指针ip的值
	fmt.Printf("ip存储的指针地址：%x\n", ip)
	//获取指针指向的内容
	fmt.Printf("使用指针访问变量值：%d\n", *ip)
	//空指针
	var nilPtr *int
	fmt.Printf("ptr的值为：%x\n", nilPtr)
	//空指针判断
	fmt.Println(nilPtr != nil)
}
