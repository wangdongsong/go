package main

import (
	"fmt"
	"os"
)

//IO
func main() {

	fmt.Println(os.Args[1])

	fileInfo, err := os.Stat(os.Args[1])

	fmt.Println(fileInfo.IsDir())
	fmt.Println(err)

	if err != nil && fileInfo.IsDir() {
		fmt.Println(fileInfo.IsDir())
	}


	//获取当前执行程序的路径
	//file, _ := exec.LookPath(os.Args[0])

}