package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

type Reader struct {
	io.Reader
	Total   string
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	if err != nil && err != io.EOF {
		log.Log.Info(err.Error())
	}
	r.Current += int64(n)
	fmt.Printf("\r进度 %d Mib, 总共 %s", r.Current/1024/1024, r.Total)
	return
}

func downloadFileProgress(url, filename string) error {
	log.Log.Info(filename, " is Downloading ...")
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() { _ = r.Body.Close() }()

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	reader := &Reader{
		Reader: r.Body,
		Total:  dObjects.getTotal(path.Base(filename)),
	}

	_, err = io.Copy(f, reader)
	fmt.Println()
	return err
}
