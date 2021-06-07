package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"simple-ims/models"
	"simple-ims/utils"
	"strings"
	"sync"
)

const (
	remoteUploadDir = "uploads/remote/"
)

func init() {
	utils.Mkdir(remoteUploadDir)
}

type Remote struct {
	sync.Mutex

	files []models.File
}

// Exec 执行下载远程文件到当前的主机上保存
func (r *Remote) Exec(ctx iris.Context) {
	path := ctx.URLParamEscape("path")
	if path == "" {
		response(ctx, false, "远程文件路径非法", nil)
		return
	}

	address, err := url.Parse(path)
	if err != nil {
		response(ctx, false, "远程文件路径非法: "+err.Error(), nil)
		return
	}

	resp, err := http.Get(path)
	if err != nil {
		response(ctx, false, "下载上传文件失败"+err.Error(), nil)
		return
	}

	if resp.StatusCode != 200 {
		response(ctx, false, "下载上传文件失败:!200", nil)
		return
	}
	disposition := resp.Header.Get(context.ContentDispositionHeaderKey)

	filename := strings.ReplaceAll(disposition, "attachment;filename=", "")
	if filename == "" {
		filename = filepath.Base(address.Path)
	}
	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		response(ctx, false, "获取远程文件内容失败:"+err.Error(), nil)
		return
	}

	file, err := os.Create(remoteUploadDir + filename)
	if err != nil {
		response(ctx, false, "创建保存远程文件失败:"+err.Error(), nil)
		return
	}

	_, err = file.Write(data)
	defer file.Close()

	if err != nil {
		response(ctx, false, "写入远程文件失败:"+err.Error(), nil)
		return
	}

	response(ctx, true, "success", nil)

}

// List 文件列表
func (r *Remote) List(ctx iris.Context) {
	dir, err := os.ReadDir(remoteUploadDir)
	if err != nil {
		response(ctx, false, fmt.Sprintf("读目录 %s 失败: %s", remoteUploadDir, err), nil)
		return
	}

	files := make([]models.File, len(dir))
	for i, entry := range dir {
		info, _ := entry.Info()
		files[i] = models.File{
			Id:        i,
			Name:      entry.Name(),
			Path:      remoteUploadDir + entry.Name(),
			Size:      getFriendlyFileSize(info.Size()),
			CreatedAt: info.ModTime(),
		}
	}
	r.Lock()
	r.files = files
	r.Unlock()
	response(ctx, true, "success", files)
}

// Download 下载文件
func (r *Remote) Download(ctx iris.Context) {
	name := ctx.URLParam("name")
	r.Lock()
	defer r.Unlock()
	for _, f := range r.files {
		if f.Name == name {
			_ = ctx.SendFile(f.Path, f.Name)
			return
		}
	}
	response(ctx, false, "下载失败", nil)
}

func getFriendlyFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}
