package controller

import (
	"archive/zip"
	"fmt"
	"github.com/elliotchance/pie/pie"
	"github.com/kataras/iris"
	"io/ioutil"
	"os"
	"path"
	"simple-ims/models"
	"simple-ims/utils"
	"strings"
	"time"
)

//项目删除
func ProjectDelete(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if err != nil {
		response(ctx, false, "项目ID非法"+err.Error(), nil)
		return
	}
	phm := models.ProjectHistoryModel{}
	if err = phm.DeleteByProjectId(id); err != nil {
		response(ctx, false, "删除当前项目历史所有版本失败:"+err.Error(), nil)
		return
	}
	pm := &models.ProjectModel{Id: id}
	if err := pm.Delete(); err != nil {
		response(ctx, false, "删除项目失败:"+err.Error(), nil)
		return
	}
	log(ctx, fmt.Sprintf("删除项目[ %s ],描述[ %s ]", pm.Name, pm.Desc))
	response(ctx, true, "删除项目成功", nil)
}

//项目列表
func ProjectLists(ctx iris.Context) {
	type item struct {
		models.ProjectModel
		Version  string `json:"version"`
		Download int    `json:"download"`
		Log      string `json:"log"`
	}
	var list []item
	pm := models.ProjectModel{}
	projects, err := pm.Find()
	if err != nil {
		response(ctx, false, "获取项目列表失败:"+err.Error(), nil)
		return
	}
	for _, v := range projects {
		if v.PHId == 0 {
			list = append(list, item{v, "", 0, ""})
		} else {
			phm := models.ProjectHistoryModel{ProjectId: v.Id}
			project, err := phm.First()
			if err != nil {
				v.UpdatedAt = v.CreatedAt
				list = append(list, item{v, "", 0, ""})
			} else {
				list = append(list, item{v, project.Version, project.Download, project.Log})
			}
		}
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
	if version == "" || RHIds == "" {
		response(ctx, false, "请输入版本号,选择对应资源", nil)
		return
	}
	pm := &models.ProjectModel{Id: projectId}
	pm, err = pm.First()
	if err != nil {
		response(ctx, false, "获取当前项目详情失败:"+err.Error(), nil)
		return
	}
	phm := &models.ProjectHistoryModel{ProjectId: projectId}
	phm, err = phm.First()
	if err == nil && utils.VersionCompare(version, phm.Version) < 1 {
		response(ctx, false, "当前版本必须高于最新版本:"+phm.Version, nil)
		return
	}
	uploadDir := "uploads/" + time.Now().Format("2006/01/")
	if !utils.Mkdir(uploadDir) {
		response(ctx, false, "创建文件夹失败", nil)
		return
	}
	zipDir := uploadDir + utils.FileName(pm.Name, version) + utils.Zip
	fZip, err := os.Create(zipDir)
	if err != nil {
		response(ctx, false, "创建zip文件失败: "+err.Error(), nil)
		return
	}
	w := zip.NewWriter(fZip)
	resourceLog := ""
	defer func() { _ = w.Close() }()
	for _, id := range utils.StrToIntSlice(RHIds, ",") {
		rhm := &models.ResourceHistoryModel{Id: id}
		rhm, err := rhm.First()
		if err != nil {
			response(ctx, false, "获取资源版本失败: "+err.Error(), nil)
			return
		}
		fw, err := w.Create(path.Base(rhm.Path))
		if err != nil {
			response(ctx, false, "创建打包文件失败: "+err.Error(), nil)
			return
		}
		fileContent, err := ioutil.ReadFile(rhm.Path)
		if err != nil {
			response(ctx, false, "读取要打包的文件内容失败"+err.Error(), nil)
			return
		}
		_, err = fw.Write(fileContent)
		if err != nil {
			response(ctx, false, "将文件内容写入压缩包失败"+err.Error(), nil)
			return
		}
		resourceLog += fmt.Sprintf("\t%s:\t%s\n", rhm.File, rhm.Log)
	}
	comment := ""
	comment += "项目描述: " + pm.Desc + "\n\n"
	comment += "项目日志: " + logStr + "\n\n"
	comment += "资源日志:\n" + resourceLog + "\n\n"
	comment += "版权所有,侵权必究, Copyright ©无锡亚天光电科技有限公司"
	err = w.SetComment(comment)
	if err != nil {
		utils.Error("向压缩包写入注释失败")
	}
	phm = &models.ProjectHistoryModel{
		ProjectId: projectId,
		Version:   version,
		Log:       logStr,
		RHIds:     RHIds,
		Path:      zipDir,
		Hash:      utils.Md5Str(string(rune(projectId)) + version + logStr + RHIds),
	}
	if err := phm.Insert(); err != nil {
		response(ctx, false, "保存项目版本失败: "+err.Error(), nil)
		return
	}
	pm.PHId = phm.Id
	if err := pm.Update(); err != nil {
		response(ctx, false, "更新当前项目失败:"+err.Error(), nil)
		return
	}
	log(ctx, fmt.Sprintf("添加项目版本[ %s ]", comment))
	response(ctx, true, "保存项目版本成功", nil)
}

//添加项目
func ProjectAdd(ctx iris.Context) {
	name := ctx.FormValue("name")
	desc := ctx.FormValue("desc")
	if name == "" || desc == "" {
		response(ctx, false, "请输入项目名称和简介", nil)
		return
	}

	user := auth(ctx)
	if user == nil {
		return
	}
	pm := &models.ProjectModel{
		Name:   name,
		Desc:   desc,
		UserId: user.ID,
	}
	if err := pm.Insert(); err != nil {
		response(ctx, false, "保存项目失败:"+err.Error(), nil)
		return
	}
	log(ctx, fmt.Sprintf("添加项目[ %s ],描述[ %s ]", pm.Name, pm.Desc))
	response(ctx, true, "保存项目成功", nil)
}

//项目详情
func ProjectDetail(ctx iris.Context) {
	id, err := ctx.URLParamInt("id")
	if err != nil || id < 1 {
		response(ctx, false, "项目ID不合法", nil)
		return
	}
	pm := &models.ProjectModel{Id: id}
	model, err := pm.First()
	if err != nil || model == nil {
		response(ctx, false, "查找项目失败", nil)
		return
	}
	if model.PHId == 0 {
		response(ctx, true, "", nil)
		return
	}
	ph := &models.ProjectHistoryModel{Id: model.PHId}
	ph, err = ph.First()
	if err != nil || ph == nil {
		response(ctx, false, "查找项目版本失败", nil)
		return
	}
	var RHIds pie.Strings
	RHIds = strings.Split(ph.RHIds, ",")
	rh := models.ResourceHistoryModel{}
	rhs, err := rh.FindByIDs(RHIds.Ints())
	if err != nil || len(rhs) == 0 {
		response(ctx, false, "查找资源失败", nil)
		return
	}
	for k, v := range rhs {
		rm := &models.ResourceModel{Id: v.ResourceId}
		rm, _ = rm.First()
		if rm != nil {
			rhs[k].File = rm.Name
		}
	}
	response(ctx, true, "查找项目详情成功", rhs)
}

//项目更新
func ProjectUpdate(ctx iris.Context) {
	id, err := ctx.PostValueInt("id")
	if err != nil || id < 1 {
		response(ctx, false, "项目ID不合法", nil)
		return
	}
	name := ctx.FormValue("name")
	desc := ctx.FormValue("desc")
	if name == "" || desc == "" {
		response(ctx, false, "请输入项目名称和简介", nil)
		return
	}

	user := auth(ctx)
	if user == nil {
		return
	}
	pm := &models.ProjectModel{
		Id:     id,
		Name:   name,
		Desc:   desc,
		UserId: user.ID,
	}
	if err := pm.Update(); err != nil {
		response(ctx, false, "更新项目失败:"+err.Error(), nil)
		return
	}
	response(ctx, true, "更新项目成功", nil)
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
	if err = model.Update(); err != nil {
		utils.Error("更新项目下载量失败:", err)
	}
	log(ctx, fmt.Sprintf("下载项目[ %s ],版本[ %s ]", phm.Path, phm.Log))
	err = ctx.SendFile(model.Path, path.Base(model.Path))
	if err != nil {
		utils.Error("下载项目失败:", err)
	}
}
