package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path/filepath"
)

func GetFileMd5(filename string) string {
	path, err := filepath.Abs(filename)
	if err != nil {
		panic("Convert file absolute path error: " + path)
	}

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(h.Sum(nil))
}

func GetStringMd5(s string) string {
	return GetByteMd5([]byte(s))
}

func GetByteMd5(s []byte) string {
	md5 := md5.New()
	md5.Write(s)
	return hex.EncodeToString(md5.Sum(nil))
}
