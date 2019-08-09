package v1

import (
	"github.com/kataras/iris"
	"simple-ims/models"
	"strconv"
	"strings"
	"time"
)

//添加资源分类
func ResourceTypeAdd(ctx iris.Context) {

	name := ctx.FormValue("name")
	desc := ctx.FormValue("desc")

	if name == "" {
		response(ctx, false, "资源分类名不能为空", nil)
		return
	}
	resourceTypeModel := &models.ResourceTypeModel{
		Name:     name,
		Desc:     desc,
		CreateAt: time.Now(),
	}
	model, err := resourceTypeModel.Insert()

	if err != nil {
		response(ctx, false, "保存资源分类失败:"+err.Error(), nil)
		return
	}

	response(ctx, true, "", iris.Map{
		"resource_type": model,
		"timestamp":     time.Now().Unix(),
	})
	return

}

//删除资源分类
func ResourceTypeDelete(ctx iris.Context) {

	ids := ctx.FormValue("id")

	if ids == "" {
		response(ctx, false, "资源分类ID不能为空", nil)
		return
	}

	var id []int
	split := strings.Split(ids, ",")
	for _, v := range split {
		i, err := strconv.Atoi(v)
		if err != nil {
			response(ctx, false, "资源分类ID非法:"+err.Error(), nil)
			return
		}
		id = append(id, i)
	}

	resourceTypeModel := &models.ResourceTypeModel{}
	_, err := resourceTypeModel.DeleteByIds(id)

	if err != nil {
		response(ctx, false, "删除资源分类失败:"+err.Error(), nil)
		return
	}

	response(ctx, true, "", nil)
	return
}

//更新资源分类
func ResourceTypeUpdate(ctx iris.Context) {

	_id := ctx.FormValue("id")
	name := ctx.FormValue("name")
	desc := ctx.FormValue("desc")

	if _id == "" {
		response(ctx, false, "资源分类ID不能为空", nil)
		return
	}

	id, err := strconv.Atoi(_id)

	if err != nil {
		response(ctx, false, "资源分类ID非法:"+err.Error(), nil)
		return
	}

	resourceTypeModel := &models.ResourceTypeModel{
		ID:   id,
		Name: name,
		Desc: desc,
	}
	_, err = resourceTypeModel.Update()

	if err != nil {
		response(ctx, false, "更新资源分类失败:"+err.Error(), nil)
		return
	}

	response(ctx, true, "", nil)
	return

}

//资源分类列表
func ResourceTypeLists(ctx iris.Context) {

	resourceTypeModel := &models.ResourceTypeModel{}
	model, err := resourceTypeModel.All()

	if err != nil {
		response(ctx, false, "获取资源分类列表失败:"+err.Error(), nil)
		return
	}

	response(ctx, true, "", iris.Map{
		"resource_types": model,
		"timestamp":      time.Now().Unix(),
	})
	return

}
