package container

import "github.com/sirupsen/logrus"

func RunContainerInitProcess(command string, args []string) error {

	logrus.Infof("command %s", command)

	return nil;

}