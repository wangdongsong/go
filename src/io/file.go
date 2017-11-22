package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"path"
)

//IO
func main() {

	//path := os.Args[1]

	//fmt.Println(os.Args[1])

	fileInfo, err := os.Stat(os.Args[1])

	//fmt.Println(fileInfo.IsDir())
	//fmt.Println(err)

	/*
	if err != nil {
		error := os.Mkdir(path, 0755)
		fmt.Println(error)
	}*/

	if err == nil && fileInfo.IsDir() {
		fmt.Println(fileInfo.IsDir())
	}

	ioutil.WriteFile(path.Join(os.Args[1], "text.txt"), []byte("Hello"), 0644)

	//获取当前执行程序的路径
	//file, _ := exec.LookPath(os.Args[0])

}

//判断文件夹或文件是否存存
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, nil
}