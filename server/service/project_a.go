package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateProjectA
// @description   create a ProjectA
// @param     projectA               model.ProjectA
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateProjectA(projectA model.ProjectA) (err error) {
	err = global.GVA_DB.Create(&projectA).Error
	return err
}

// @title    DeleteProjectA
// @description   delete a ProjectA
// @auth                     （2020/04/05  20:22）
// @param     projectA               model.ProjectA
// @return                    error

func DeleteProjectA(projectA model.ProjectA) (err error) {
	err = global.GVA_DB.Delete(projectA).Error
	return err
}

// @title    DeleteProjectAByIds
// @description   delete ProjectAs
// @auth                     （2020/04/05  20:22）
// @param     projectA               model.ProjectA
// @return                    error

func DeleteProjectAByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.ProjectA{},"id in ?",ids.Ids).Error
	return err
}

// @title    UpdateProjectA
// @description   update a ProjectA
// @param     projectA          *model.ProjectA
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateProjectA(projectA *model.ProjectA) (err error) {
	err = global.GVA_DB.Save(projectA).Error
	return err
}

// @title    GetProjectA
// @description   get the info of a ProjectA
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    ProjectA        ProjectA

func GetProjectA(id uint) (err error, projectA model.ProjectA) {
	err = global.GVA_DB.Where("id = ?", id).First(&projectA).Error
	return
}

// @title    GetProjectAInfoList
// @description   get ProjectA list by pagination, 分页获取ProjectA
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetProjectAInfoList(info request.ProjectASearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.ProjectA{})
    var projectAs []model.ProjectA
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Male != nil {
        db = db.Where("`male` = ?",info.Male)
    }
    if info.Age != 0 {
        db = db.Where("`age` = ?",info.Age)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&projectAs).Error
	return err, projectAs, total
}