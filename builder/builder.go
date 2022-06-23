package builder

import (
	"os/exec"
)

func BuildClipboard(c *Config) {

	SetConfig(*c)
	
	cmd := exec.Command("go","build","-o","dist/"+ c.Filename + ".exe","-ldflags","-H=windowsgui","./build")

	err := cmd.Run()

	if err != nil {
		panic(err)
	}

}
