package download

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/TwinProduction/go-color"

	"github.com/deployKubernetesInCHINA/dkc-command/src/config"
	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

type downloadObject struct {
	Name  string `json:"name"`
	Md5   string `json:"md5"`
	Total string `json:"total"`
	Desc  string `json:"desc"`
}

type DownloadObjects []downloadObject

var dObjects DownloadObjects

func (do DownloadObjects) getTotal(name string) string {
	for _, r := range do {
		if r.Name == name {
			return r.Total
		}
	}
	return ""
}
func (do DownloadObjects) getMd5(name string) string {
	for _, r := range do {
		if r.Name == name {
			return r.Md5
		}
	}
	return ""
}
func UpdateMirror() error {
	return checkMirror()
}
func checkMirror() error {
	var err error
	_check := func(m string) error {
		if !strings.HasPrefix(m, "http") && !strings.HasPrefix(m, "https") {
			m = "http://" + m
		}
		log.Log.Debugln(m)
		if mUrl, err := url.Parse(m); err != nil {
			log.Log.Debug(err.Error())
			return err
		} else {
			if mUrl.Scheme == "" {
				mUrl.Scheme = "http"
			}
			//mUrl.User = url.UserPassword(src.UserInfo["user"], src.UserInfo["password"])
			mUrl.Path = mUrl.Path + "/repo.json"
			config.Kconfig.Mirror = m
			client := http.Client{
				Timeout: 5 * time.Second,
			}
			if res, err := client.Get(mUrl.String()); err != nil {
				log.Log.Error(err.Error())
				return err
			} else if res.StatusCode == 200 {
				defer res.Body.Close()
				b, _ := ioutil.ReadAll(res.Body)
				if err := json.Unmarshal(b, &dObjects); err != nil {
					log.Log.Fatal(color.Ize(color.Red, "解析repo.json失败, 退出下载"), err.Error())
				}
				if err := ioutil.WriteFile("repo.json", b, 0644); err != nil {
					log.Log.Fatal(color.Ize(color.Red, "保存repo.json失败,退出下载"), err.Error())
				}
			}
			return nil
		}
	}

	if err = _check(config.Kconfig.Mirror); err != nil {
		for _, m := range config.Mirrors {
			if err = _check(m); err == nil {
				embackMd5()
				return nil
			}
		}
	} else {
		embackMd5()
	}
	return err
}
func embackMd5() {

	for i, v := range dObjects {
		mUrl, _ := url.Parse(config.Kconfig.Mirror)
		mUrl.Path = mUrl.Path + filepath.Base(v.Name) + ".md5"
		if res, err := http.Get(mUrl.String()); err == nil && res.StatusCode == 200 {
			defer res.Body.Close()
			if d, err := ioutil.ReadAll(res.Body); err != nil {
				dObjects[i].Md5 = ""
			} else {
				dObjects[i].Md5 = strings.TrimSuffix(string(d), "\n")
			}
		}
	}
}
