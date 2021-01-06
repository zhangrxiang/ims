package controller

import (
	"github.com/kataras/iris/v12"
	"path"
	"simple-ims/models"
)

func VersionDownload(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		response(ctx, false, "软件Id非法", nil)
		return
	}
	rh := &models.ResourceHistoryModel{Id: id}
	rh, err = rh.First()
	if err != nil {
		response(ctx, false, "文件不存在", nil)
		return
	}
	err = ctx.SendFile(rh.Path, path.Base(rh.Path))
	if err != nil {
		response(ctx, false, "文件不存在", nil)
	}
}

func VersionLists(ctx iris.Context) {
	name := ctx.Params().Get("name")
	r := &models.ResourceModel{Name: name}
	r, err := r.First()
	if err != nil {
		response(ctx, false, "无", nil)
		return
	}
	rh := models.ResourceHistoryModel{ResourceId: r.Id}
	list, err := rh.FindBy()
	if err != nil {
		response(ctx, false, "无", nil)
		return
	}
	response(ctx, true, "", list)
}
