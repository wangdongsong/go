package jvmgo

import "fmt"

func main(){
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == ""{
		printUsage()
	} else {
		startJVM(cmd)
	}
}
