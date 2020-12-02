package controller

import (
	"fmt"
	"github.com/elliotchance/pie/pie"
	"github.com/kataras/iris"
	"path"
	"simple-ims/models"
	"simple-ims/utils"
	"strconv"
	"strings"
	"time"
)

//添加资源
func ResourceAdd(ctx iris.Context) {
	name := ctx.PostValue("name")
	desc := ctx.PostValue("desc")
	if name == "" || desc == "" {
		response(ctx, false, "请输入资源名称和描述", nil)
		return
	}
	t, err := ctx.PostValueInt("type")
	if err != nil {
		response(ctx, false, "资源类型非法", nil)
		return
	}
	user := auth(ctx)
	rm := &models.ResourceModel{
		UserId: user.ID,
		Name:   name,
		Type:   t,
		Desc:   desc,
	}
	if err = rm.Insert(); err != nil {
		response(ctx, false, "保存资源失败:"+err.Error(), nil)
		return
	}
	log(ctx, fmt.Sprintf("添加资源:[ %s ], 描述:[ %s ]", name, desc))
	response(ctx, true, "添加资源成功", nil)
}

//更新资源
func ResourceUpdate(ctx iris.Context) {
	id, err := ctx.PostValueInt("id")
	if err != nil {
		response(ctx, false, "资源ID不合法:"+err.Error(), nil)
		return
	}
	name := ctx.PostValue("name")
	desc := ctx.PostValue("desc")
	if name == "" || desc == "" {
		response(ctx, false, "请输入资源名称和描述", nil)
		return
	}
	t, err := ctx.PostValueInt("type")
	if err != nil {
		response(ctx, false, "资源类型非法:"+err.Error(), nil)
		return
	}
	user := auth(ctx)
	rm := models.ResourceModel{
		Id:     id,
		UserId: user.ID,
		Name:   name,
		Type:   t,
		Desc:   desc,
	}
	if err = rm.Update(); err != nil {
		response(ctx, false, "更新资源失败:"+err.Error(), nil)
		return
	}
	log(ctx, fmt.Sprintf("更新资源:[ %s ], 描述:[ %s ]", name, desc))
	response(ctx, true, "更新资源成功", nil)
}

//添加资源版本
func ResourceUpgrade(ctx iris.Context) {
	resId, err := ctx.PostValueInt("resource_id")
	if err != nil {
		response(ctx, false, "资源ID非法", nil)
		return
	}
	version := ctx.PostValue("version")
	l := ctx.PostValue("log")
	if version == "" || l == "" {
		response(ctx, false, "请填写版本号或日志", nil)
		return
	}
	rhm := &models.ResourceHistoryModel{ResourceId: resId}
	oldVersion := ""
	tmp, err := rhm.First()
	if tmp != nil && err == nil {
		oldVersion = rhm.Version
		if utils.VersionCompare(version, tmp.Version) < 1 {
			response(ctx, false, "当前版本必须高于最新版本: "+tmp.Version, nil)
			return
		}
	}
	file, info, err := ctx.FormFile("file")
	if file == nil {
		response(ctx, false, "请上传文件", nil)
		return
	}
	if err != nil {
		response(ctx, false, "获取上传文件失败:"+err.Error(), nil)
		return
	}

	defer func() {
		_ = file.Close()
	}()

	uploadDir := "uploads/" + time.Now().Format("2006/01/")
	if !utils.Mkdir(uploadDir) {
		response(ctx, false, "创建文件夹失败", nil)
		return
	}

	rhm.Hash, err = utils.Md5File(file)
	if err != nil {
		response(ctx, false, "获取文件MD5失败:"+err.Error(), nil)
		return
	}
	tmp, err = rhm.First()
	if tmp != nil && err == nil {
		response(ctx, false, "相同的文件hash已存在:"+tmp.Hash, iris.Map{
			"resource": tmp,
		})
		return
	}
	rhm.Version = version
	rhm.Log = l
	rhm.File = info.Filename
	rhm.Path = uploadDir + utils.FileName(info.Filename, version)
	err = utils.CopyFile(rhm.Path, file)
	if err != nil {
		response(ctx, false, "保存文件失败:"+err.Error(), nil)
		return
	}

	user := auth(ctx)
	rhm.UserId = user.ID
	if err = rhm.Insert(); err != nil {
		response(ctx, false, "添加资源版本失败:"+err.Error(), nil)
		return
	}

	rm := &models.ResourceModel{
		Id:   resId,
		RHId: rhm.Id,
	}
	if err = rm.Update(); err != nil {
		response(ctx, false, "添加资源失败:"+err.Error(), nil)
		return
	}

	log(ctx, fmt.Sprintf("添加资源版本:[ %s ], 版本[ %s ] -> [ %s ], 日志[ %s ]", rhm.File, oldVersion, rhm.Version, rhm.Log))
	response(ctx, true, "添加资源版本成功", iris.Map{
		"resource": rhm,
	})
}

//删除资源
func ResourceDelete(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if err != nil {
		response(ctx, false, "资源ID不合法", nil)
		return
	}

	rm := &models.ResourceModel{Id: id}
	rm, err = rm.First()
	if err != nil {
		response(ctx, false, "查找要删除的资源失败:"+err.Error(), nil)
		return
	}

	rh := &models.ResourceHistoryModel{ResourceId: rm.Id}
	var ids pie.Ints
	ids, err = rh.FindIDBy()
	if err != nil {
		response(ctx, false, "查找所有要删除的历史资源版本失败:"+err.Error(), nil)
		return
	}
	ph := &models.ProjectHistoryModel{}
	var rhIds pie.Strings
	rhIds, err = ph.FindRHIDs()
	if err != nil {
		response(ctx, false, "查找项目失败:"+err.Error(), nil)
		return
	}
	rhIds = strings.Split(strings.Join(rhIds, ","), ",")
	compIDs := rhIds.Unique().Ints().Sort()
	added, removed := ids.Diff(compIDs)
	if len(removed) != len(ids) || len(added) != len(compIDs) {
		response(ctx, false, "当前资源已经被项目占用,禁止删除", nil)
		return
	}
	err = rh.DeleteBy(ids)
	if err != nil {
		response(ctx, false, "删除所有资源版本失败", nil)
		return
	}
	err = rm.Delete()
	if err != nil {
		response(ctx, false, "删除资源失败", nil)
		return
	}
	log(ctx, fmt.Sprintf("删除资源[ %s ], 描述[ %s ]", rm.Name, rm.Desc))
	response(ctx, true, "删除资源成功", nil)
	//go func(resources *[]models.ResourceModel) {
	//	for _, resource := range *resources {
	//		err := os.Remove(resource.Path)
	//		if err != nil {
	//			log.Println("remove file", resource.Name, err)
	//		} else {
	//			log.Println("remove file", resource.Name)
	//		}
	//	}
	//}(resources)

}

//资源列表
func ResourceLists(ctx iris.Context) {
	type item struct {
		models.ResourceModel
		File     string `json:"file"`
		Version  string `json:"version"`
		Download int    `json:"download"`
		Log      string `json:"log"`
	}
	var list []item
	rm := models.ResourceModel{}

	if ctx.URLParamExists("resource_type") {
		resourceType, err := ctx.URLParamInt("resource_type")
		if err != nil {
			response(ctx, false, "资源分类ID不合法:"+err.Error(), nil)
			return
		}
		rm = models.ResourceModel{
			Type: resourceType,
		}

	}
	resources, err := rm.Find()
	if err != nil {
		response(ctx, false, "获取资源失败:"+err.Error(), nil)
		return
	}
	for _, v := range resources {
		if v.RHId != 0 {
			rhm := models.ResourceHistoryModel{
				Id: v.RHId,
			}
			resource, err := rhm.First()
			if err != nil && err != models.NoRecordExists {
				response(ctx, false, "获取资源版本失败:"+err.Error(), nil)
				return
			}
			if resource != nil {
				list = append(list, item{
					ResourceModel: v,
					Version:       resource.Version,
					Download:      resource.Download,
					File:          resource.File,
					Log:           resource.Log,
				})
			}
		} else {
			list = append(list, item{
				ResourceModel: v,
				Version:       "",
				Download:      0,
				File:          "",
				Log:           "",
			})
		}
	}
	response(ctx, true, "获取资源列表成功", list)
}

//资源列表
func ResourceGroupLists(ctx iris.Context) {
	typeModel := &models.ResourceTypeModel{}
	allType, err := typeModel.Find()
	if err != nil {
		response(ctx, false, "获取资源类型列表失败:"+err.Error(), nil)
		return
	}
	if len(allType) > 0 {
		resourceModel := &models.ResourceModel{}
		type item struct {
			ID        int       `json:"id"`
			Name      string    `json:"name"`
			Desc      string    `json:"desc"`
			Version   string    `json:"version"`
			File      string    `json:"file"`
			Log       string    `json:"log"`
			Download  int       `json:"download"`
			UpdatedAt time.Time `json:"updated_at"`
		}
		var data []map[string]interface{}
		for _, t := range allType {
			resourceModel.Type = t.Id
			resources, err := resourceModel.Find()
			if err != nil {
				response(ctx, false, "获取资源失败:"+err.Error(), nil)
				return
			}
			if len(resources) > 0 {
				var lists []item
				for _, v := range resources {
					if v.RHId != 0 {
						rhm := models.ResourceHistoryModel{Id: v.RHId}
						result, err := rhm.First()
						if err != nil {
							response(ctx, false, "获取历史资源失败:"+err.Error(), nil)
							return
						}
						if result != nil {
							lists = append(lists, item{
								ID:        v.Id,
								Name:      v.Name,
								Desc:      v.Desc,
								Version:   result.Version,
								Download:  result.Download,
								Log:       result.Log,
								File:      utils.FileName(result.File, result.Version),
								UpdatedAt: result.CreatedAt,
							})
						} else {
							lists = append(lists, item{
								Name: v.Name,
								Desc: v.Desc,
							})
						}
					} else {
						lists = append(lists, item{
							Name:      v.Name,
							Desc:      v.Desc,
							UpdatedAt: v.UpdatedAt,
						})
					}
				}
				resource := make(map[string]interface{})
				resource["name"] = t.Name
				resource["desc"] = t.Desc
				resource["lists"] = lists
				data = append(data, resource)
			} else {
				data = append(data, map[string]interface{}{
					"name":  t.Name,
					"desc":  t.Desc,
					"lists": nil,
				})
			}
		}
		response(ctx, true, "", iris.Map{
			"resources": data,
			"timestamp": time.Now().Unix(),
		})
		return
	}
	response(ctx, true, "暂无资源", nil)
}

//下载文件
func ResourceDownload(ctx iris.Context) {

	idStr := ctx.URLParam("id")
	version := ctx.URLParam("version")

	id, err := strconv.Atoi(idStr)
	if err != nil || version == "" {
		response(ctx, false, "文件ID和版本不能为空", nil)
		return
	}

	hm := &models.ResourceHistoryModel{
		ResourceId: id,
		Version:    version,
	}
	hm, err = hm.First()
	if err != nil {
		response(ctx, false, "当前资源不存在", nil)
		return
	}

	if err = hm.Increment(); err != nil {
		response(ctx, false, "更新资源下载量失败", nil)
		return
	}

	dm := models.DownloadModel{
		RHId:   hm.Id,
		UserId: auth(ctx).ID,
	}
	if err = dm.Insert(); err != nil {
		response(ctx, false, "添加下载资源记录失败", nil)
		return
	}

	log(ctx, fmt.Sprintf("下载[ %s ], Id[ %d ], 版本[ %s ], 日志[ %s ]", hm.File, hm.Id, hm.Version, hm.Log))
	err = ctx.SendFile(hm.Path, path.Base(hm.Path))
	if err != nil {
		response(ctx, false, "文件不存在"+err.Error(), nil)
	}
}
