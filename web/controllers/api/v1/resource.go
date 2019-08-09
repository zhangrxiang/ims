package v1

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/kataras/iris"
	"io"
	"os"
	"path"
	"simple-ims/models"
	"strconv"
	"time"
)

func ResourceAdd(ctx iris.Context) {

	name := ctx.FormValue("name")
	_type := ctx.FormValue("type")
	version := ctx.FormValue("version")
	desc := ctx.FormValue("desc")

	if name == "" || _type == "" {
		response(ctx, false, "请输入资源名称,选择资源类型", nil)
		return
	}

	t, err := strconv.Atoi(_type)
	if err != nil {
		response(ctx, false, "资源类型不存在:"+err.Error(), nil)
		return
	}
	var resourceModel = &models.ResourceModel{
		Name:    name,
		Type:    t,
		Version: version,
		Desc:    desc,
	}

	file, info, err := ctx.FormFile("file")
	if file != nil {
		if err != nil {
			response(ctx, false, "获取上传文件失败:"+err.Error(), nil)
			return
		}

		defer file.Close()

		uploadDir := "uploads/" + time.Now().Format("2006/01/")
		_, err = os.Stat(uploadDir)

		if os.IsNotExist(err) {
			err := os.MkdirAll(uploadDir, 0666)
			if err != nil {
				response(ctx, false, "创建文件夹失败:"+err.Error(), nil)
				return
			}
		}

		h := md5.New()
		_, err = io.Copy(h, file)
		if err != nil {
			response(ctx, false, "读取文件失败:"+err.Error(), nil)
			return
		}

		resourceModel.Hash = hex.EncodeToString(h.Sum(nil))
		model, err := resourceModel.FindByHash(resourceModel.Hash)
		if model != nil {
			response(ctx, true, "保存文件成功:", iris.Map{
				"resource": model,
			})
			return
		}
		resourceModel.File = info.Filename
		resourceModel.Path = uploadDir + resourceModel.Hash + path.Ext(info.Filename)
		resourceModel.CreateAt = time.Now()
		out, err := os.OpenFile(resourceModel.Path, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			response(ctx, false, "打开文件失败:"+err.Error(), nil)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			response(ctx, false, "保存文件失败:"+err.Error(), nil)
			return
		}
	}

	model, err := resourceModel.Insert()

	if err != nil {
		response(ctx, false, "添加资源失败:", nil)
		return
	}
	response(ctx, true, "保存文件成功:", iris.Map{
		"resource": model,
	})
}

func ResourceLists(ctx iris.Context) {
	resourceModel := &models.ResourceModel{}
	model, err := resourceModel.All()

	if err != nil {
		response(ctx, false, "获取资源列表失败:"+err.Error(), nil)
		return
	}

	response(ctx, true, "", iris.Map{
		"resources": model,
		"timestamp": time.Now().Unix(),
	})
	return
}
