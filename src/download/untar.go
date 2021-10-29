package download

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/deployKubernetesInCHINA/dkc-command/src/pkg/log"
)

func Untar(tarball, target string) error {
	var tarReader *tar.Reader
	reader, err := os.Open(tarball)
	if err != nil {
		return err
	}
	defer reader.Close()
	log.Log.Info(filepath.Base(tarball), " 解压中...")
	if strings.HasSuffix(tarball, ".tar.gz") {
		uncompressedStream, err := gzip.NewReader(reader)
		defer uncompressedStream.Close()
		if err != nil {
			log.Log.Errorln("ExtractTarGz: NewReader failed")
		}
		tarReader = tar.NewReader(uncompressedStream)
	} else if strings.HasSuffix(tarball, ".tar") {
		tarReader = tar.NewReader(reader)
	} else {
		log.Log.Errorln("not supported file type")
	}

	for {
		reader, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(target, reader.Name)
		info := reader.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		_, err = io.Copy(file, tarReader)
		file.Close()
		if err != nil {
			return err
		}
	}
	log.Log.Info("解压成功. 删除 ", tarball)
	os.Remove(tarball)
	return nil
}
