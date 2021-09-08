package inventory

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func CopyInventory() {

	//传参path
	varPath, err := filepath.Abs(config.Kconfig.InventoryFile)
	if err != nil {
		log.Log.Fatal(err.Error())
	}
	//使用path
	usePath, err := filepath.Abs(config.InventoryPath)
	if err != nil {
		log.Log.Fatal(err.Error())
	}
	if strings.Compare(varPath, usePath) != 0 {
		//copy invetory file

		//backup inventory
		if err = BackUp(config.InventoryPath); err != nil {
			log.Log.Fatal(err.Error())
		}

		//copy from varpath to usepath
		data, err := os.ReadFile(varPath)
		if err != nil {
			log.Log.Fatal(err.Error())
		}
		err = os.WriteFile(config.InventoryPath, data, 0644)
		if err != nil {
			log.Log.Fatal(err.Error())
		}
	}
}

func BackUp(fpath string) error {
	if f, ok := os.Stat(fpath); ok == nil {
		if f.IsDir() {
			return fmt.Errorf("%s is not file", fpath)
		}
		newFileName := fpath + ".bak-" + time.Now().Format("2006-01.02-15:04:05")
		data, err := os.ReadFile(fpath)
		if err != nil {
			return fmt.Errorf("read from %s error.", fpath)
		}
		err = os.WriteFile(newFileName, data, 0644)
		if err != nil {
			return fmt.Errorf("write to %s error.", newFileName)
		}
	}
	return nil
}
