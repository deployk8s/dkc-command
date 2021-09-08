package download

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/TwinProduction/go-color"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func Delete(name string) error {
	var repo DownloadObjects
	if b, err := ioutil.ReadFile("repo.json"); err != nil {
		log.Log.Errorln(color.Ize(color.Red, err.Error()))
		return err
	} else {
		json.Unmarshal(b, &repo)
	}
	for _, v := range repo {
		if v.Name == name {
			dirName := strings.Split(name, ".tar")[0]
			if _, err := os.Stat(dirName); err == nil {
				return os.RemoveAll(dirName)
			}
		}
	}
	return nil
}
