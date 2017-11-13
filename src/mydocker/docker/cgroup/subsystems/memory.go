package subsystems

import (
	"io/ioutil"
	"path"
	"fmt"
)

type MemorySubSystem struct {}

func(s *MemorySubSystem) Set(cgroupPath string, res *ResourceConfig) error {

	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {

		if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644); err != nil{
			return fmt.Errorf("set cgroup memory fail %v", err)
		}
		return nil
	} else {
		return err
	}
}


func (s *MemorySubSystem) Name() string {
	return "memory"
}
