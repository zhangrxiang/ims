package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"simple-ims/models"
	"strconv"
	"strings"
	"time"
)

//添加资源分类
func ResourceTypeAdd(ctx iris.Context) {
	name := ctx.FormValue("name")
	desc := ctx.FormValue("desc")
	if name == "" || desc == "" {
		response(ctx, false, "资源分类名和描述不能为空", nil)
		return
	}
	rtm := &models.ResourceTypeModel{
		Name: name,
		Desc: desc,
	}

	if err := rtm.Insert(); err != nil {
		response(ctx, false, "保存资源分类失败:"+err.Error(), nil)
		return
	}

	log(ctx, fmt.Sprintf("添加资源分类:[ %s ],描述:[ %s ]", name, desc))
	response(ctx, true, "", iris.Map{
		"resource_type": rtm,
		"timestamp":     time.Now().Unix(),
	})
	return

}

//删除资源分类
func ResourceTypeDelete(ctx iris.Context) {
	var (
		names = ""
		id    []int
	)
	ids := ctx.FormValue("id")
	if ids == "" {
		response(ctx, false, "资源分类ID不能为空", nil)
		return
	}
	split := strings.Split(ids, ",")
	for _, v := range split {
		i, err := strconv.Atoi(v)
		if err != nil {
			response(ctx, false, "资源分类ID非法:"+err.Error(), nil)
			return
		}
		rt := &models.ResourceTypeModel{Id: i}
		if rt, err := rt.First(); rt != nil && err == nil {
			names += rt.Name
		}
		rm := &models.ResourceModel{Type: i}
		rm, err = rm.First()
		if err == nil {
			response(ctx, false, "当前资源分类已经被占用,不能删除", nil)
			return
		}
		id = append(id, i)
	}

	rt := &models.ResourceTypeModel{}
	if err := rt.DeleteByIds(id); err != nil {
		response(ctx, false, "删除资源分类失败:"+err.Error(), nil)
		return
	}
	log(ctx, fmt.Sprintf("删除资源分类:[ %s ]", names))
	response(ctx, true, "", nil)
}

//更新资源分类
func ResourceTypeUpdate(ctx iris.Context) {
	id, err := ctx.PostValueInt("id")
	if err != nil {
		response(ctx, false, "资源分类ID非法:"+err.Error(), nil)
		return
	}
	name := ctx.FormValue("name")
	desc := ctx.FormValue("desc")
	if name == "" || desc == "" {
		response(ctx, false, "资源分类名和描述不能为空", nil)
		return
	}
	rt := &models.ResourceTypeModel{Id: id}
	oldName := rt.Name
	oldDesc := rt.Desc
	rt.Name = name
	rt.Desc = desc
	if err = rt.Update(); err != nil {
		response(ctx, false, "更新资源分类失败:"+err.Error(), nil)
		return
	}
	log(ctx, fmt.Sprintf("更新资源分类名:[ %s ] -> [ %s ],分类描述:[ %s ] -> [ %s ]", oldName, rt.Name, oldDesc, desc))
	response(ctx, true, "", nil)
}

//资源分类列表
func ResourceTypeLists(ctx iris.Context) {
	rtm := &models.ResourceTypeModel{}
	list, err := rtm.Find()
	if err != nil {
		response(ctx, false, "获取资源分类列表失败:"+err.Error(), nil)
		return
	}
	response(ctx, true, "", iris.Map{
		"item":      list,
		"timestamp": time.Now().Unix(),
	})
	return

}
