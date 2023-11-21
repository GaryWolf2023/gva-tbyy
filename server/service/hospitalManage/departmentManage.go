package hospitalManage

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	response2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/response"
	"time"
)

type DepartmentManageService struct{}

func (d *DepartmentManageService) GetList(req request2.PageInfo) (error, response2.PageRes) {
	var infoList []response.Info
	var resInfo response2.PageRes
	var total int64
	var offset int
	var limit int
	if req.PageSize == 0 {
		limit = 9999999
		offset = 0
	} else {
		limit = req.PageSize
		offset = (req.Page - 1) * req.PageSize
	}
	result := global.GVA_DB.Table("hr_department").Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keyword)).Offset(offset).Limit(limit).Find(&infoList).Count(&total)
	if result.Error != nil {
		return result.Error, response2.PageRes{}
	}
	resInfo.Total = total
	resInfo.Data = infoList
	return nil, resInfo
}
func (d *DepartmentManageService) CreateDep(req request.CreateInfo) error {
	var createInfo hospitalManage.HrDepartment
	var parentInfo hospitalManage.HrDepartment
	result := global.GVA_DB.Table("hr_department").Where("id = ?", req.ParentId).First(&parentInfo)
	if result.Error != nil {
		return result.Error
	}
	createInfo.CreateDate = time.Now()
	createInfo.WriteDate = time.Now()
	createInfo.CreateUid = req.EmployeeId
	createInfo.WriteUid = req.EmployeeId
	createInfo.CompanyId = req.CompanyId
	createInfo.ParentId = req.ParentId
	createInfo.Name = req.Name
	createInfo.Note = req.Note

	createInfo.CompleteName = parentInfo.CompleteName + "/" + req.Name
	createInfo.ParentPath = parentInfo.ParentPath + "/" + fmt.Sprintf("%v", parentInfo.ID) + "/"
	result1 := global.GVA_DB.Table("hr_department").Create(&createInfo)
	if result1.Error != nil {
		return result1.Error
	}
	return nil
}

// UpdateDep 更新科室部门
func (d *DepartmentManageService) UpdateDep(req request.UpdateInfo) error {
	var updateInfo hospitalManage.HrDepartment
	var parentInfo hospitalManage.HrDepartment
	result := global.GVA_DB.Table("hr_department").Where("id = ?", req.Id).First(&updateInfo)
	if result.Error != nil {
		return result.Error
	}
	updateInfo.CreateDate = time.Now()
	updateInfo.WriteDate = time.Now()
	updateInfo.CreateUid = req.EmployeeId
	updateInfo.WriteUid = req.EmployeeId
	updateInfo.CompanyId = req.CompanyId
	updateInfo.ParentId = req.ParentId
	updateInfo.Name = req.Name
	updateInfo.Note = req.Note

	updateInfo.CompleteName = parentInfo.CompleteName + "/" + req.Name
	updateInfo.ParentPath = parentInfo.ParentPath + "/" + fmt.Sprintf("%v", parentInfo.ID) + "/"
	result1 := global.GVA_DB.Table("hr_department").Create(&updateInfo)
	if result1.Error != nil {
		return result1.Error
	}
	return nil
}

func (d *DepartmentManageService) DeleteDep(req request.DeleteInfo) error {
	var deleteInfo hospitalManage.HrDepartment
	result1 := global.GVA_DB.Table("hr_department").Where("id = ?", req.Id).Find(&deleteInfo).Unscoped().Delete(&deleteInfo)
	if result1.Error != nil {
		return result1.Error
	}
	return nil
}

// BatchesOfDelete 批量删除
func (d *DepartmentManageService) BatchesOfDelete(req request.BatchesOfDelete) error {
	var deleteInfo hospitalManage.HrDepartment
	result1 := global.GVA_DB.Table("hr_department").Unscoped().Delete(&deleteInfo, req)
	if result1.Error != nil {
		return result1.Error
	}
	return nil
}
