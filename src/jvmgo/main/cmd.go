/*
 * 在ch01的基础上搜索class文件
 */
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"jvmgo/classpath"
	"jvmgo/classfile"
)


func startJVM(cmd *Cmd){
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

	fmt.Printf("classpath: %s class: %s args: %v\n", cmd.cpOption, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
}

func printClassInfo(cf *classfile.ClassFile)  {
	fmt.Printf("version:%v.%v\n",cf.MajorVersion(),cf.MinorVersion())

}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile  {
	classData, _, err := cp.ReadClass(className)
	if err != nil{
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
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
