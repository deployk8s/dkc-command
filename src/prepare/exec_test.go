package prepare

import (
	"io"
	"os"
	"os/exec"
	"testing"
)

func Test_Install(t *testing.T) {

	cmd := exec.Command("sudo", "ls", "/")
	cmd.Stdout = os.Stdout
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Log(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "123456\n")
	}()

	err = cmd.Run()
	if err != nil {
		t.Log(err)
	}

}
