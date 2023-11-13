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
)

type PrescriptionQualificationService struct{}

func (p *PrescriptionQualificationService) CreatePQ(req request.CreatePQ) error {
	var info hospitalManage.HrEmployeeChufang
	info.CreateUid = req.EmployeeId
	info.WriteUid = req.EmployeeId
	info.Name = req.Name
	info.Active = req.Active
	result := global.GVA_DB.Table("hr_employee_chufang").Where("name = ?", req.Name).First(&info)
	if info.Name == req.Name && result.Error == nil {
		return errors.New(fmt.Sprintf("已经有名为%s的处方权", req.Name))
	}
	result1 := global.GVA_DB.Table("hr_employee_chufang").Create(&info)
	if result1.Error != nil {
		return result1.Error
	}
	return nil
}
func (p *PrescriptionQualificationService) UpdatePQ(req request.UpdatePQ) error {
	var info hospitalManage.HrEmployeeChufang
	info.WriteUid = req.EmployeeId
	info.Name = req.Name
	info.Active = req.Active
	info.ID = req.ID
	fmt.Println(info)
	result := global.GVA_DB.Table("hr_employee_chufang").Where("id = ?", req.ID).Updates(&info)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (p *PrescriptionQualificationService) DeletePQ(req request.DeletePQ) error {
	result := global.GVA_DB.Table("hr_employee_chufang").Where("id = ?", req.ID).Unscoped().Delete(&req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (p *PrescriptionQualificationService) GetListPQ(req request2.PageInfo) (error, interface{}) {
	var list []response.GetList
	var total int64
	var pageRes response2.PageRes
	offset := (req.Page - 1) * req.PageSize
	limit := req.PageSize
	if limit == 0 {
		limit = 20
	}
	if req.Keyword != "" {
		global.GVA_DB.Table("hr_employee_chufang").Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keyword)).Offset(offset).Limit(limit).Find(&list).Count(&total)
	} else {
		global.GVA_DB.Table("hr_employee_chufang").Offset(offset).Limit(limit).Find(&list).Count(&total)
	}
	pageRes.Data = list
	pageRes.Total = total
	return nil, pageRes
}
func (p *PrescriptionQualificationService) AddPPQ(req request.AddPPQ) error {
	var info request.AddPPQ
	var dbinfo hospitalManage.HrEmployeePersonPrescription
	global.GVA_DB.AutoMigrate(&dbinfo)
	result := global.GVA_DB.Table("hr_employee_person_prescription").Where("employee_id = ? AND prescription_id = ?", req.EmployeeId, req.PrescriptionId).First(&info)
	if result.Error == nil {
		return errors.New(fmt.Sprintf("已经存在"))
	}
	result1 := global.GVA_DB.Table("hr_employee_person_prescription").Create(&req)
	if result1.Error != nil {
		return result1.Error
	}
	return nil
}
func (p *PrescriptionQualificationService) DeletePPQ(req request.DeletePPQ) error {
	result1 := global.GVA_DB.Table("hr_employee_person_prescription").Where("employee_id = ? AND prescription_id = ?", req.EmployeeId, req.PrescriptionId).Unscoped().Delete(&req)
	if result1.Error != nil {
		return result1.Error
	}
	return nil
}
func (p *PrescriptionQualificationService) GetListPPQ(req request.GetListPPQ) (error, []hospitalManage.HrEmployeeChufang) {
	var dbinfo []hospitalManage.HrEmployeePersonPrescription
	var idList []int
	var info []hospitalManage.HrEmployeeChufang
	result1 := global.GVA_DB.Table("hr_employee_person_prescription").Where("employee_id = ?", req.EmployeeId).Find(&dbinfo)
	if result1.Error != nil {
		return result1.Error, []hospitalManage.HrEmployeeChufang{}
	}
	if len(dbinfo) == 0 {
		return nil, nil
	}
	for _, v := range dbinfo {
		idList = append(idList, v.PrescriptionId)
	}
	result2 := global.GVA_DB.Table("hr_employee_chufang").Find(&info, idList)
	if result2.Error != nil {
		return result2.Error, []hospitalManage.HrEmployeeChufang{}
	}
	return nil, info
}
