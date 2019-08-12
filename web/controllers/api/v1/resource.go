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
	"strings"
	"time"
)

//添加资源
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
			err := os.MkdirAll(uploadDir, 0777)
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
		out, err := os.OpenFile(resourceModel.Path, os.O_RDWR|os.O_CREATE, 0777)

		if err != nil {
			response(ctx, false, "打开文件失败:"+err.Error(), nil)
			return
		}
		defer out.Close()

		_, _ = file.Seek(0, io.SeekStart)
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

//删除资源
func ResourceDelete(ctx iris.Context) {
	ids := ctx.FormValue("id")

	if ids == "" {
		response(ctx, false, "资源ID不能为空", nil)
		return
	}

	var id []int
	split := strings.Split(ids, ",")
	for _, v := range split {
		i, err := strconv.Atoi(v)
		if err != nil {
			response(ctx, false, "资源ID非法:"+err.Error(), nil)
			return
		}
		id = append(id, i)
	}

	resourceModel := &models.ResourceModel{}
	_, err := resourceModel.DeleteByIds(id)

	if err != nil {
		response(ctx, false, "删除资源失败:"+err.Error(), nil)
		return
	}

	response(ctx, true, "", nil)
	return
}

//更新资源
func ResourceUpdate(ctx iris.Context) {
	_id := ctx.FormValue("id")
	name := ctx.FormValue("name")
	_type := ctx.FormValue("type")
	version := ctx.FormValue("version")
	desc := ctx.FormValue("desc")

	if name == "" || _type == "" {
		response(ctx, false, "请输入资源名称,选择资源类型", nil)
		return
	}

	id, err := strconv.Atoi(_id)
	if err != nil {
		response(ctx, false, "资源ID格式不合法:"+err.Error(), nil)
		return
	}
	t, err := strconv.Atoi(_type)
	if err != nil {
		response(ctx, false, "资源类型格式不合法:"+err.Error(), nil)
		return
	}
	var resourceModel = &models.ResourceModel{
		ID:      id,
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
			//response(ctx, true, "保存文件成功:", iris.Map{
			//	"resource": model,
			//})
			//return
		}
		resourceModel.File = info.Filename
		resourceModel.Path = uploadDir + resourceModel.Hash + path.Ext(info.Filename)
		resourceModel.CreateAt = time.Now()
		out, err := os.OpenFile(resourceModel.Path, os.O_RDWR|os.O_CREATE, 0777)

		if err != nil {
			response(ctx, false, "打开文件失败:"+err.Error(), nil)
			return
		}
		defer out.Close()

		_, _ = file.Seek(0, io.SeekStart)
		_, err = io.Copy(out, file)
		if err != nil {
			response(ctx, false, "保存文件失败:"+err.Error(), nil)
			return
		}
	}

	model, err := resourceModel.Update()

	if err != nil {
		response(ctx, false, "添加资源失败:", nil)
		return
	}
	response(ctx, true, "保存文件成功:", iris.Map{
		"resource": model,
	})
}

//资源列表
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

//下载文件
func ResourceDownload(ctx iris.Context) {
	p := ctx.URLParam("path")
	f := ctx.URLParam("file")

	if p == "" || f == "" {
		response(ctx, false, "文件路径不能为空", nil)
		return
	}

	err := ctx.SendFile(p, f)

	if err != nil {
		response(ctx, false, "文件不存在", nil)
	}

}
