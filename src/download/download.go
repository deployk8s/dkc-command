package download

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/TwinProduction/go-color"

	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func Download() {
	config.InitDownload()
	if err := checkMirror(); err != nil {
		log.Log.Error(err.Error())
		os.Exit(1)
	}
	log.Log.Infoln(config.Kconfig.Mirror, " Connect OK. Use this mirror.")
	createDownloadDir()
	downloadFromRepo()
}

func createDownloadDir() {
	if d, ok := os.Stat(config.Kconfig.DownloadDir); ok != nil {
		log.Log.Debugln(config.Kconfig.DownloadDir, " not exist. it will be created")
		os.MkdirAll(config.Kconfig.DownloadDir, os.ModePerm)
	} else {
		if d.IsDir() == false {
			log.Log.Errorln(config.Kconfig.DownloadDir, " is not directory.")
			os.Exit(1)
		}
	}
}

func downloadFromRepo() {
	for _, fo := range dObjects {

		//是否只下载一个离线文件
		if config.Kconfig.OnlyOne != "" && !strings.Contains(fo.Name, config.Kconfig.OnlyOne) {
			log.Log.Debugln(fo.Name, "not contains", config.Kconfig.OnlyOne)
			continue
		}
		fileUrl, _ := url.Parse(config.Kconfig.Mirror)
		fileUrl.Path = fileUrl.Path + "/" + fo.Name
		filePath := path.Join(config.Kconfig.DownloadDir, fo.Name)
		log.Log.Debug("Process ", filePath)
		if err := Downloadfile(fileUrl.String(), filePath); err != nil {
			log.Log.Errorln(color.Ize(color.Red, err.Error()))
			os.Exit(1)
		}
	}
}

func Downloadfile(url, filePath string) error {
	//是否检查md5
	fileName := path.Base(filePath)
	if config.Kconfig.CheckMD5 {
		if !checksum(fileName) {
			//检查md5没通过,开始下载
			if err := downloadFileProgress(url, filePath); err != nil {
				return err
			}
			//下载完,再次进行md5校验
			if !checksum(fileName) {
				return fmt.Errorf(filePath, "md5 check Failed.")
			}
		} else {
			log.Log.Println(filePath, "is latest")
		}
	} else {
		//不检查md5 直接下载
		if err := downloadFileProgress(url, filePath); err != nil {
			return err
		}
	}

	return nil
}

func checksum(name string) bool {
	localPath := path.Join(config.Kconfig.DownloadDir, name)
	if strings.Compare(dObjects.getMd5(name), "") == 0 {
		//如果repo里没有记录md5,并且文件已存在则不进行比较
		if _, err := os.Stat(localPath); err != nil {
			return false
		}
		return true
	}
	if d, err := os.ReadFile(localPath); err != nil {
		log.Log.Debugln(err.Error())
		return false
	} else {
		var downloadMd5sum string
		downloadMd5sum = fmt.Sprintf("%x", md5.Sum(d))
		log.Log.Debug("Local md5: ", downloadMd5sum, "; latest md5: ", dObjects.getMd5(name))
		return downloadMd5sum == dObjects.getMd5(name)
	}

}
