package main

import (
	"mydocker/docker/cgroup/subsystems"
	"mydocker/docker/container"
	"os"
	log "github.com/sirupsen/logrus"
	"mydocker/docker/cgroup"
	"strings"
)

func Run(tty bool, comArray []string, res *subsystems.ResourceConfig) {
	parent, writePipe := container.NewParentProcess(tty)
	if err := parent.Start(); err != nil {
		log.Errorf("New parent process erro")
		return
	}

	if err := parent.Start(); err != nil {
		log.Error(err)
	}

	// use mydocker-cgroup as cgroup name
	// create cgroup manager, call the Set and Apply method to set resources limit and run on docker
	cgroupManager := cgroup.NewCgroupManager("mydocker-cgroup")
	defer cgroupManager.Destroy()

	//set resources limit
	cgroupManager.Set(res)

	//将容器进程加入到各个subsystem挂载对应的cgroup中
	cgroupManager.Apply(parent.Process.Pid)

	//对容器设置完成限制后，启动容器
	sendInitCommand(comArray, writePipe)

	parent.Wait()

	os.Exit(0)
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}