package src

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/TwinProduction/go-color"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func CheckYes(echoMsg string) bool {
	var guess string
	const yes = "yes"
	const no = "no"
	for {
		fmt.Print(echoMsg, " Please type Y/n : ")
		if _, err := fmt.Scanf("%s", &guess); err != nil {
			continue
		}
		if yes == strings.ToLower(guess) || "y" == strings.ToLower(guess) || "" == strings.ToLower(guess) {
			return true
		} else if no == strings.ToLower(guess) || "n" == strings.ToLower(guess) {
			return false
		}
	}
}

func IsSupported() bool {
	if runtime.GOOS == "linux" && (Release == "7" || Release == "8") {
		return true
	}
	log.Log.Debug(color.Ize(color.Red, "only centos7/centos8 is supported"))
	return false
}
