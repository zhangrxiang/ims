package utils

import (
	"archive/zip"
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
)

func StrToIntAlice(str, sep string) []int {
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
		dst.Close()
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
	return strings.TrimSuffix(path.Base(p), path.Ext(p)) + "-" + version + path.Ext(p)
}

func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = prefix + "/" + header.Name
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
