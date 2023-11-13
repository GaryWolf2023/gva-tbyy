package hospitalManage

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	response2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/response"
	"time"
)

type JobTitleService struct{}

func (j *JobTitleService) CreateJobTitle(req request.CreateJobTitle) error {
	var info hospitalManage.HospitalPerformanceCalled
	var createinfo hospitalManage.CreateJT
	result := global.GVA_DB.Table("hospital_performance_called").Where("name = ?", req.Name).First(&info)
	if info.Name == req.Name && result.Error == nil {
		return errors.New(fmt.Sprintf("已经有名为%s的职称", req.Name))
	}
	createinfo.Name = req.Name
	createinfo.CreateUid = req.EmployeeId
	createinfo.WriteUid = req.EmployeeId
	createinfo.WriteDate = time.Now()
	createinfo.CreateDate = time.Now()
	createinfo.Grade = req.Grade
	createinfo.Note = req.Note
	createinfo.Floorvalue = req.Floorvalue
	createinfo.Ceilingvalue = req.Ceilingvalue
	createinfo.IndicatValue = req.IndicatValue
	createinfo.Category = req.Category
	result1 := global.GVA_DB.Table("hospital_performance_called").Create(&createinfo)
	if result1.Error != nil {
		return result1.Error
	}
	return nil
}
func (j *JobTitleService) UpdateJobTitle(req request.UpdateJobTitle) error {
	var info hospitalManage.HospitalPerformanceCalled
	result := global.GVA_DB.Table("hospital_performance_called").Where("id = ?", req.ID).First(&info)
	if result.Error != nil {
		return result.Error
	}
	info.WriteUid = req.EmployeeId
	info.Name = req.Name
	info.Category = req.Category
	info.Ceilingvalue = req.Ceilingvalue
	info.Floorvalue = req.Floorvalue
	info.Note = req.Note
	info.IndicatValue = req.IndicatValue
	info.Grade = req.Grade
	info.WriteDate = time.Time{}
	result1 := global.GVA_DB.Table("hospital_performance_called").Where("id = ?", req.ID).Updates(&info)
	if result1.Error != nil {
		return result1.Error
	}
	return nil
}
func (j *JobTitleService) DeleteJobTitle(req request.DeleteJobTitle) error {
	result := global.GVA_DB.Table("hospital_performance_called").Where("id = ?", req.ID).Unscoped().Delete(&req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (j *JobTitleService) GetJobTitleList(req request2.PageInfo) (error, interface{}) {
	var list []response.JobTitleList
	var total int64
	var pageRes response2.PageRes
	offset := (req.Page - 1) * req.PageSize
	limit := req.PageSize
	if limit == 0 {
		limit = 20
	}
	if req.Keyword != "" {
		global.GVA_DB.Table("hospital_performance_called").Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keyword)).Offset(offset).Limit(limit).Find(&list).Count(&total)
	} else {
		global.GVA_DB.Table("hospital_performance_called").Offset(offset).Limit(limit).Find(&list).Count(&total)
	}
	pageRes.Total = total
	pageRes.Data = list
	return nil, pageRes
}
