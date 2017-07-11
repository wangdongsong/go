package main

import (
	"flag"
	"fmt"
	"os"
)

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

func startJVM(cmd *Cmd){
	fmt.Printf("classpath: %s class: %s args: %v\n", cmd.cpOption, cmd.class, cmd.args)
}

type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOption string
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
