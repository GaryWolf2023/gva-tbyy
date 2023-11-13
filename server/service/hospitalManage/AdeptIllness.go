package hospitalManage

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/response"
)

type AdeptIllnessService struct{}

func (a *AdeptIllnessService) GetAdeptIllnessList(req request.GetIllnessList) (error, []response.EmployeeAdeptIllness) {
	var illnessList []response.EmployeeAdeptIllness
	result := global.GVA_DB.Table("hr_employee_adept_illnesses").Where("employee_id = ?", req.EmployeeId).Find(&illnessList)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("%v", result.Error)), nil
	}
	return nil, illnessList
}

func (a *AdeptIllnessService) AddAdeptIllness(req request.AddIllness) error {
	fmt.Println(req)
	var illness response.EmployeeAdeptIllness
	var adeptIllness hospitalManage.HrEmployeeAdeptIllness
	err := global.GVA_DB.AutoMigrate(&adeptIllness)
	if err != nil {
		return err
	}
	result := global.GVA_DB.Table("hr_employee_adept_illnesses").Where("employee_id = ? AND illness_id = ?", req.EmployeeId, req.IllnessId).First(&illness)
	fmt.Println("++++++++++++++++++++++++++", result.Error)
	if result.Error == nil {
		return errors.New(fmt.Sprintf("%v", "不能重复添加"))
	}
	global.GVA_DB.Table("hr_employee_adept_illnesses").Create(&req)
	return nil
}

func (a *AdeptIllnessService) DeleteAdeptIllnessItem(req request.DeleteIllness) error {
	fmt.Println(req)
	result := global.GVA_DB.Table("hr_employee_adept_illnesses").Where("employee_id = ? AND illness_id = ?", req.EmployeeId, req.IllnessId).Unscoped().Delete(&req)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("%v", result.Error))
	}
	return nil
}
