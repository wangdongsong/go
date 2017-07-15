/*
 * 在ch01的基础上搜索class文件
 */
package jvmgo

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

import "jvm/classpath"


func startJVM(cmd *Cmd){
	cp := classpath.parse(cmd.XjreOption, cmd.cpOption)

	fmt.Printf("classpath: %s class: %s args: %v\n", cmd.cpOption, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil{
		fmt.Printf("Could not find or load basic class %s\n", cmd.class)
		return
	}

	fmt.Printf("class data:%v\n", classData)
}

type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	XjreOption string
	class string
	args []string
}

func parseCmd()  *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "claspath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()

	if(len(args) > 0){
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage()  {
	fmt.Printf("Usage: %s [options] class [args...]\n", os.Args[0])
}
