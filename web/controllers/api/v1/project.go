package v1

import (
	"archive/zip"
	"github.com/kataras/iris"
	"io/ioutil"
	"log"
	"os"
	"path"
	"simple-ims/models"
	"simple-ims/utils"
	"time"
)

//项目删除
func ProjectDelete(ctx iris.Context) {
	id := ctx.FormValue("id")
	ids := utils.StrToIntAlice(id, ",")
	if ids == nil {
		response(ctx, false, "项目ID非法", nil)
		return
	}
	pm := models.ProjectModel{}
	_, err := pm.DeleteByIds(ids)
	if err != nil {
		response(ctx, false, "删除项目失败:"+err.Error(), nil)
		return
	}

	phm := models.ProjectHistoryModel{}
	_, err = phm.DeleteByProjectId(pm.ID)
	if err != nil {
		response(ctx, false, "删除当前项目历史所有版本失败:"+err.Error(), nil)
		return
	}
	response(ctx, true, "删除项目成功", nil)
}

//项目列表
func ProjectLists(ctx iris.Context) {
	type item struct {
		models.ProjectModel
		models.ProjectHistoryModel
	}
	var list []item
	pm := models.ProjectModel{}
	model, err := pm.FindBy()
	if err != nil {
		response(ctx, true, "获取项目列表失败:"+err.Error(), nil)
		return
	}
	for _, v := range *model {
		phm := models.ProjectHistoryModel{
			ProjectId: v.ID,
		}
		first, err := phm.First()
		if err != nil {
			response(ctx, true, "获取项目版本失败:"+err.Error(), nil)
			return
		}
		list = append(list, item{v, *first})
	}
	response(ctx, true, "获取项目列表成功", list)
}

//添加项目版本
func ProjectUpgrade(ctx iris.Context) {
	projectId, err := ctx.PostValueInt("project_id")
	if err != nil {
		response(ctx, false, "项目ID非法", nil)
		return
	}
	version := ctx.FormValue("version")
	RHIds := ctx.FormValue("rh_ids")
	logStr := ctx.FormValue("log")
	if version == "" || RHIds == "" || logStr == "" {
		response(ctx, false, "请输入版本号和日志,选择对应资源", nil)
		return
	}
	pm := models.ProjectModel{
		ID: projectId,
	}
	projectModel, err := pm.FirstBy()
	if err != nil {
		response(ctx, false, "获取当前项目详情失败:"+err.Error(), nil)
		return
	}
	uploadDir := "uploads/" + time.Now().Format("2006/01/")
	if !utils.Mkdir(uploadDir) {
		response(ctx, false, "创建文件夹失败", nil)
		return
	}
	zipDir := uploadDir + utils.FileName(projectModel.Name, version) + ".zip"
	phm := models.ProjectHistoryModel{
		ProjectId: projectId,
		Version:   version,
		Log:       logStr,
		RHIds:     RHIds,
		Path:      zipDir,
		Hash:      utils.Md5Str(string(projectId) + version + logStr + RHIds),
	}
	model, err := phm.Insert()
	if err != nil {
		response(ctx, false, "保存项目版本失败:"+err.Error(), nil)
		return
	}
	fZip, err := os.Create(zipDir)
	if err != nil {
		response(ctx, false, "创建zip文件失败"+err.Error(), nil)
		return
	}
	w := zip.NewWriter(fZip)
	defer w.Close()
	for _, id := range utils.StrToIntAlice(model.RHIds, ",") {
		rhm := models.ResourceHistoryModel{
			ID: id,
		}
		resourceHistoryModel, err := rhm.FirstBy()
		if err != nil {
			response(ctx, false, "获取资源版本失败"+err.Error(), nil)
			return
		}
		fw, err := w.Create(path.Base(resourceHistoryModel.Path))
		if err != nil {
			response(ctx, false, "创建打包文件失败"+err.Error(), nil)
			return
		}
		fileContent, err := ioutil.ReadFile(resourceHistoryModel.Path)
		if err != nil {
			response(ctx, false, "读取要打包的文件内容失败"+err.Error(), nil)
			return
		}
		_, err = fw.Write(fileContent)
		if err != nil {
			response(ctx, false, "将文件内容写入压缩包失败"+err.Error(), nil)
			return
		}
		err = w.SetComment(logStr)
		if err != nil {
			log.Println("向压缩包写入注释失败")
		}
	}

	pm.PHId = model.ID
	_, err = pm.Update()
	if err != nil {
		response(ctx, false, "更新当前项目失败:"+err.Error(), nil)
		return
	}
	response(ctx, true, "保存项目版本成功", model)
}

//添加项目
func ProjectAdd(ctx iris.Context) {
	name := ctx.FormValue("name")
	desc := ctx.FormValue("desc")
	if name == "" || desc == "" {
		response(ctx, false, "请输入项目名称和简介", nil)
		return
	}
	pm := &models.ProjectModel{
		Name:   name,
		Desc:   desc,
		UserId: auth(ctx).ID,
	}
	model, err := pm.Insert()
	if err != nil {
		response(ctx, false, "保存项目失败:"+err.Error(), nil)
		return
	}
	response(ctx, false, "保存项目成功", model)
}

//项目下载
func ProjectDownload(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	version := ctx.URLParam("version")
	if err != nil || version == "" {
		response(ctx, false, "项目ID和版本不能为空", nil)
		return
	}
	phm := models.ProjectHistoryModel{
		ProjectId: id,
		Version:   version,
	}
	model, err := phm.First()
	if err != nil {
		response(ctx, false, "当前项目版本不存在", nil)
		return
	}
	model.Download += 1
	_, err = model.Update()
	if err != nil {
		log.Println("更新项目下载量失败:", err)
	}
	err = ctx.SendFile(model.Path, path.Base(model.Path))
	if err != nil {
		log.Println("下载项目失败:", err)
	}
}
