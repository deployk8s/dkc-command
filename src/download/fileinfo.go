package download

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/TwinProduction/go-color"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

type FileInfoResp struct {
	Name  string `json:"name"`
	State string `json:"state"`
	Desc  string `json:"desc"`
	Local string `json:"local"`
}

func GetFileInfo() []FileInfoResp {
	var fis []FileInfoResp
	var repo DownloadObjects
	if b, err := ioutil.ReadFile("repo.json"); err != nil {
		log.Log.Errorln(color.Ize(color.Red, err.Error()))
		return fis
	} else {
		json.Unmarshal(b, &repo)
	}
	for _, v := range repo {
		dirName := strings.Split(v.Name, ".")[0]
		if _, err := os.Stat(dirName); err == nil {
			fn, fs := statStart(dirName)
			fis = append(fis, FileInfoResp{Name: v.Name, Desc: v.Desc, State: fmt.Sprintf("Size: %s, Md5: %s", v.Total, v.Md5), Local: fmt.Sprintf("已下载, 解压之后Size: %d Mib. 文件数: %d", fs/1024/1024, fn)})
		} else {
			fis = append(fis, FileInfoResp{Name: v.Name, Desc: v.Desc, State: fmt.Sprintf("Size: %s, Md5: %s", v.Total, v.Md5), Local: fmt.Sprintf("未下载")})
		}
	}
	return fis
}

//获取目录dir下的文件大小
func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() { //目录
			wg.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir(subDir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

//sema is a counting semaphore for limiting concurrency in dirents
var sema = make(chan struct{}, 20)

//读取目录dir下的文件信息
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func statStart(root string) (int64, int64) {

	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	wg.Add(1)
	go walkDir(root, &wg, fileSizes)
	go func() {
		wg.Wait() //等待goroutine结束
		close(fileSizes)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		}
	}
	return nfiles, nbytes
}
