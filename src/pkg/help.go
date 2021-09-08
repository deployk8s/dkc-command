package pkg

import (
	"os"
	"os/exec"
)

func Help(commands ...string) {
	args := []string{}
	args = append(args, "help")
	args = append(args, commands...)
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
}
