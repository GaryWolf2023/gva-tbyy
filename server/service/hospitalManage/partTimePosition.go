package hospitalManage

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	response2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/privateUtils"
	"time"
)

type PartTimePositionService struct{}

// =================================================================================
// 个人兼职管理
// =================================================================================

func (p *PartTimePositionService) GetPartTimePositionList(req request.GetItemList) (error, interface{}) {
	var partTimePosition []response.PartTimePosition
	err0 := global.GVA_DB.AutoMigrate(&partTimePosition)
	if err0 != nil {
		fmt.Printf("%v", err0)
	}
	global.GVA_DB.Table("hr_employee_person_job").Where("employee_id = ?", req.EmployeeId).Find(&partTimePosition)
	return nil, nil
}
func (p *PartTimePositionService) AddPartTimePosition(req request.AddItem) error {
	var partTimePosition hospitalManage.HrEmployeePersonJob
	err0 := global.GVA_DB.AutoMigrate(&partTimePosition)
	if err0 != nil {
		fmt.Printf("%v", err0)
	}
	result := global.GVA_DB.Table("hr_employee_person_job").Create(&req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (p *PartTimePositionService) DeletePartTimePosition(req request.DeleteItem) error {
	result := global.GVA_DB.Table("hr_employee_person_job").Unscoped().Delete(&req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// =================================================================================
// job管理
// =================================================================================

func (p *PartTimePositionService) GetJobList(req request2.PageInfo) (error, interface{}) {
	var list []hospitalManage.HrJob
	var res response2.PageRes
	var total int64
	limit := req.PageSize
	offset := (req.Page - 1) * req.PageSize
	if req.PageSize == 0 {
		limit = 999999
	}
	result := global.GVA_DB.Table("hr_job").Where("name ->>'zh_CN' LIKE ?", fmt.Sprintf("%%%s%%", req.Keyword)).Offset(offset).Limit(limit).Find(&list).Count(&total)
	if result.Error != nil {
		return result.Error, nil
	}
	res.Data = list
	res.Total = total
	return nil, res
}

// 创建
func (p *PartTimePositionService) CreateJob(req request.CreateJob) error {
	var modelInfo hospitalManage.HrJob
	fmt.Println(req)
	modelInfo.CreateDate = time.Now()
	modelInfo.WriteDate = time.Now()
	modelInfo.CreateUid = req.EmployeeId
	modelInfo.WriteUid = req.EmployeeId
	dataName := map[string]string{
		"en_US": req.Name,
		"zh_CN": req.Name,
	}
	jsonData, _ := json.Marshal(dataName)
	modelInfo.Name = string(jsonData)
	modelInfo.DepartmentId = req.DepartmentId
	modelInfo.NoOfEmployee = req.NoOfEmployee
	modelInfo.NoOfRecruitment = req.NoOfRecruitment
	modelInfo.UserId = req.UserId
	modelInfo.Active = req.Active
	modelInfo.NaturePost = req.NaturePost
	modelInfo.CategoryPost = req.CategoryPost
	modelInfo.JobCode = req.JobCode
	modelInfo.Show = req.Show
	result := global.GVA_DB.Table("hr_job").Create(&modelInfo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *PartTimePositionService) UpdateJob(req request.UpdateJob) error {
	var modelInfo = hospitalManage.HrJob{}
	err := privateUtils.MapFields(req.UpdateObject, modelInfo)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(modelInfo)
	modelInfo.WriteDate = time.Now()
	modelInfo.WriteUid = req.UserId
	fmt.Println(modelInfo)
	return nil
}

func (p *PartTimePositionService) DeleteJob(req request.DeleteJob) error {
	var info interface{}
	result := global.GVA_DB.Table("hr_job").Where("id = ?", req.ID).First(&info).Unscoped().Delete(&info)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
