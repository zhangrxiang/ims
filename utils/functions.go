package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	TarGz       = ".tar.gz"
	Zip         = ".zip"
	LocalFormat = ".20060102"
)

func Encode(str string) string {
	return strings.Map(func(r rune) rune {
		return r + 10
	}, str)
}

func Decode(str string) string {
	return strings.Map(func(r rune) rune {
		return r - 10
	}, str)
}

func StrToIntSlice(str, sep string) []int {
	var intStr []int
	split := strings.Split(str, sep)
	for _, v := range split {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil
		}
		intStr = append(intStr, i)
	}

	return intStr
}

func Mkdir(p string) bool {
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		err := os.MkdirAll(p, 0777)
		if err != nil {
			return false
		}
	}
	return true
}

func CopyFile(p string, src multipart.File) error {
	dst, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE, 0777)
	defer func() {
		_ = dst.Close()
	}()
	if err != nil {
		return err
	}
	_, _ = src.Seek(0, io.SeekStart)
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil
}

func Md5File(file io.Reader) (string, error) {
	instance := md5.New()
	_, err := io.Copy(instance, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(instance.Sum(nil)), nil
}

func Md5Str(value string) string {
	instance := md5.New()
	instance.Write([]byte(value))
	return hex.EncodeToString(instance.Sum(nil))
}

func FileName(p string, version string) string {
	version2 := time.Now().Format(LocalFormat) + "." + version
	if strings.HasSuffix(p, TarGz) {
		if strings.Contains(p, version) {
			return strings.ReplaceAll(p, version, version+version2)
		}
		return strings.TrimSuffix(path.Base(p), TarGz) + version2 + TarGz
	}
	return strings.TrimSuffix(path.Base(p), path.Ext(p)) + version2 + path.Ext(p)
}

//版本比较 0.0.0
func VersionCompare(v1, v2 string) int8 {
	sv1 := strings.Split(v1, ".")
	sv2 := strings.Split(v2, ".")
	for k, v := range sv1 {
		v1, _ := strconv.Atoi(v)
		v2, _ := strconv.Atoi(sv2[k])
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}
	return 0
}
