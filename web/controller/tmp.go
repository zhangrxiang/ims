package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"io/ioutil"
	"os"
	"simple-ims/utils"
	"time"
)

type TmpFile struct {
	Key  string    `json:"key"`
	Name string    `json:"name"`
	Path string    `json:"path"`
	Time time.Time `json:"time"`
}

var TmpFiles = map[string]TmpFile{}

func TmpUpLists(ctx iris.Context) {
	response(ctx, true, "", TmpFiles)
}

func init() {
	day := time.Hour * 24
	week := day * 7
	go func() {
		t := time.NewTicker(day)
		for {
			select {
			case <-t.C:
				for k, f := range TmpFiles {
					if time.Now().Sub(f.Time) > week {
						err := os.Remove(f.Path)
						if err != nil {
							utils.Error(fmt.Sprintf("删除 %s 失败: %s", f.Name, err))
						}
						delete(TmpFiles, k)
					}
				}

				dir, err := ioutil.ReadDir("uploads/tmp/")
				if err != nil {
					utils.Error(fmt.Sprintf("ReadDir uploads/tmp/ 失败: %s", err))
					break
				}
				for _, f := range dir {
					if !f.IsDir() && time.Now().Sub(f.ModTime()) > week*4 {
						err := os.Remove("uploads/tmp/" + f.Name())
						if err != nil {
							utils.Error(fmt.Sprintf("删除 %s 失败: %s", f.Name(), err))
						}
					} else {
						utils.Info(f.Name(), "未删除", !f.IsDir(), time.Now().Sub(f.ModTime()) > week*4)
					}
				}
			}
		}
	}()
}

func TmpUpload(ctx iris.Context) {
	file, info, err := ctx.FormFile("file")
	if err != nil {
		response(ctx, false, "获取上传文件失败:"+err.Error(), nil)
		return
	}
	if file != nil {
		uploadDir := "uploads/tmp/"
		if !utils.Mkdir(uploadDir) {
			response(ctx, false, "创建文件夹失败", nil)
			return
		}
		md5Str, err := utils.Md5File(file)
		if err != nil {
			response(ctx, false, "获取文件MD5失败:"+err.Error(), nil)
			return
		}
		err = utils.CopyFile(uploadDir+info.Filename, file)
		if err != nil {
			response(ctx, false, "保存文件失败:"+err.Error(), nil)
			return
		}
		key := md5Str[:6]
		TmpFiles[key] = TmpFile{
			Key:  key,
			Name: info.Filename,
			Path: uploadDir + info.Filename,
			Time: time.Now(),
		}
		response(ctx, true, "success", key)
	} else {
		response(ctx, false, "请上传文件", nil)
	}
}

func TmpDownload(ctx iris.Context) {
	key := ctx.URLParam("key")
	if file, ok := TmpFiles[key]; ok {
		err := ctx.SendFile(file.Path, file.Name)
		if err != nil {
			response(ctx, false, "文件不存在"+err.Error(), nil)
		}
		return
	}
	response(ctx, false, "文件不存在", nil)
}
