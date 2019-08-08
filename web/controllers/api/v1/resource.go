package v1

import (
	"github.com/kataras/iris"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func ResourceAdd(ctx iris.Context) {

	file, info, err := ctx.FormFile("file")
	if err != nil {
		response(ctx, false, "获取上传文件失败:"+err.Error(), nil)
		return
	}

	defer file.Close()
	name := info.Filename

	dir := "./uploads/" + time.Now().Format("2006/01/")
	_, err = os.Stat(dir)

	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			response(ctx, false, "创建文件夹失败:"+err.Error(), nil)
			return
		}
	}
	out, err := os.OpenFile(dir+name, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		response(ctx, false, "打开文件失败:"+err.Error(), nil)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		response(ctx, false, "保存文件失败:"+err.Error(), nil)
	}

	response(ctx, false, "保存文件成功:", nil)

}
func MultipartForm(ctx iris.Context) {

	// Get the max post value size passed via iris.WithPostMaxMemory.
	maxSize := ctx.Application().ConfigurationReadOnly().GetPostMaxMemory()

	err := ctx.Request().ParseMultipartForm(maxSize)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.WriteString(err.Error())
		return
	}

	form := ctx.Request().MultipartForm

	files := form.File["files[]"]
	failures := 0
	for _, file := range files {
		_, err = saveUploadedFile(file, "./uploads")
		if err != nil {
			failures++
			_, _ = ctx.Writef("failed to upload: %s\n", file.Filename)
		}
	}
	_, _ = ctx.Writef("%d files uploaded", len(files)-failures)
}

func saveUploadedFile(fh *multipart.FileHeader, destDirectory string) (int64, error) {
	src, err := fh.Open()
	if err != nil {
		return 0, err
	}
	defer src.Close()

	out, err := os.OpenFile(filepath.Join(destDirectory, fh.Filename),
		os.O_WRONLY|os.O_CREATE, os.FileMode(0666))

	if err != nil {
		return 0, err
	}
	defer out.Close()

	return io.Copy(out, src)
}

func ResourceLists(ctx iris.Context) {}
