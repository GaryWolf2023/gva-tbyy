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

type InsuranceManageService struct{}

func (i *InsuranceManageService) CreateInsurance(req request.CreateInsurance) error {
	req.CreateDate = time.Now()
	result := global.GVA_DB.Table("hr_employee_insurance").Create(&req)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("%v", result.Error))
	}
	return nil
}
func (i *InsuranceManageService) UpdateInsurance(req request.UpdateInsurance) error {
	req.UpdateDate = time.Now()
	result := global.GVA_DB.Table("hr_employee_insurance").Updates(&req)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("%v", result.Error))
	}
	return nil
}
func (i *InsuranceManageService) DeleteInsurance(req request.DeleteInsurance) error {
	var data hospitalManage.HrEmployeeInsurance
	fmt.Println(req)
	result := global.GVA_DB.Table("hr_employee_insurance").Where("id = ?", req.ID).First(&data)
	fmt.Println(data)
	global.GVA_DB.Table("hr_employee_insurance").Delete(&data)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("%v", result.Error))
	}
	return nil
}

func (i *InsuranceManageService) GetInsuranceList(req request2.PageInfo) (error, interface{}) {
	var res response2.PageRes
	var list []response.GetInsuranceList
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var total int64
	result := global.GVA_DB.Table("hr_employee_insurance").Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keyword)).Offset(offset).Limit(limit).Find(&list).Count(&total)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("%v", result.Error)), list
	}
	res.Data = list
	res.Total = total
	return nil, res
}

func (i *InsuranceManageService) GetInsuranceInfo(req request.GetInsuranceInfo) (error, response.GetInsuranceList) {
	var info response.GetInsuranceList
	result := global.GVA_DB.Table("hr_employee_insurance").Where("id = ?", req.ID).Limit(1).Find(&info)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("%v", result.Error)), info
	}
	return nil, info
}

func (i *InsuranceManageService) GetPersonInsuranceList(req request.GetPersonInsuranceList) (error, []response.GetInsuranceList) {
	var insuranceList []response.GetPersonInsurance
	//var list []response.GetInsuranceList
	fmt.Println(req)
	result := global.GVA_DB.Table("hr_employee_person_insurances").Where("employee_id = ?", req.ID).Find(&insuranceList)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("%v", result.Error)), []response.GetInsuranceList{}
	}

	var idList []response.GetInsuranceList
	var insurance response.GetInsuranceList
	fmt.Println("---------------------++++++++++++++++++++", len(insuranceList))
	for index, v := range insuranceList {
		fmt.Println(insuranceList[index])
		global.GVA_DB.Table("hr_employee_insurance").Where("id = ?", v.InsuranceId).Find(&insurance)
		idList = append(idList, insurance)
	}
	//fmt.Println(idList)
	//if len(idList) == 0 {
	//	return errors.New("未找到数据"), []response.GetInsuranceList{}
	//}
	//
	//result1 := global.GVA_DB.Table("hr_employee_insurance").Where(idList).Find(&list)
	//if result1.Error != nil {
	//	return errors.New(fmt.Sprintf("%v", result.Error)), []response.GetInsuranceList{}
	//}
	return nil, idList
}

func (i *InsuranceManageService) AddPersonInsurance(req request.AddPersonInsurance) error {
	var record request.AddPersonInsurance
	result1 := global.GVA_DB.Table("hr_employee_person_insurances").Where("insurance_id = ? AND employee_id = ?", req.InsuranceId, req.EmployeeId).First(&record)
	if result1.Error == nil {
		return errors.New(fmt.Sprintf("%v", result1.Error))
	}
	result2 := global.GVA_DB.Table("hr_employee_person_insurances").Create(&req)
	if result2.Error != nil {
		return errors.New(fmt.Sprintf("%v", result2.Error))
	}
	return nil
}

func (i *InsuranceManageService) DeletePersonInsurance(req request.DeletePersonInsurance) error {
	//result1 := global.GVA_DB.Table("hr_employee_person_insurances").Where("insurance_id = ? AND employee_id = ?", req.InsuranceId, req.EmployeeId).Unscoped().Delete(&req)
	//var insurance hospitalManage.HrEmployeePersonInsurances
	//insurance.InsuranceId = req.InsuranceId
	//insurance.EmployeeId = req.EmployeeId
	fmt.Println("871263671273651263614236541254351235---------------------", req)
	result1 := global.GVA_DB.Table("hr_employee_person_insurances").Where("insurance_id = ? AND employee_id = ?", req.InsuranceId, req.EmployeeId).Unscoped().Delete(&req)
	fmt.Println(result1)
	if result1.Error != nil {
		return errors.New(fmt.Sprintf("%v", result1.Error))
	}
	return nil
}
