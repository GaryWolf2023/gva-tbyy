package hospitalManage

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	response2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/response"
	"strings"
	"time"
)

type StaffManageService struct{}

func (s *StaffManageService) GetStaffList(req request.PageInfoOfGet) response2.PageRes {
	var total int64
	page := req.Page
	pageSize := req.PageSize
	if pageSize == 0 {
		pageSize = 10
	}
	if page == 0 {
		page = 1
	}
	limit := pageSize
	offset := pageSize * (page - 1)
	var StaffList []response.BasicStaffInfo
	var resData response2.PageRes
	global.GVA_DB.Table("hr_employee").Count(&total).Offset(offset).Limit(limit).Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keyword)).Find(&StaffList)
	resData.Data = StaffList
	resData.Total = total
	return resData
}

func (s *StaffManageService) GetStaffInfo(id string) (response.BasicStaffInfo, error) {
	var staffInfo response.BasicStaffInfo
	result := global.GVA_DB.Table("hr_employee").Where("id = ?", id).First(&staffInfo)
	if result.Error != nil {
		return staffInfo, errors.New(fmt.Sprintf("获取失败%v", result.Error))
	}
	return staffInfo, nil
}

// 更新系统配置信息
func (s *StaffManageService) UpdateStaffInfo(req request2.StaffInfo) error {
	fmt.Println(req)
	result := global.GVA_DB.Table("hr_employee").Where("id = ?", req.ID).Updates(req)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("更新失败%v", result.Error))
	}
	return nil
}

// 更新系统配置信息
func (s *StaffManageService) UpdateSysConfig(req request2.SystemConfig) error {
	fmt.Println(req)
	result := global.GVA_DB.Table("hr_employee").Where("id = ?", req.ID).Updates(req)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("更新失败%v", result.Error))
	}
	return nil
}

// 更新基础数据
func (s *StaffManageService) UpdateBasicStaffInfo(req request2.BasicInfo) error {
	fmt.Println(req)
	result := global.GVA_DB.Table("hr_employee").Where("id = ?", req.ID).Updates(req)
	if result.Error != nil {
		return errors.New(fmt.Sprintf("更新失败%v", result.Error))
	}
	return nil
}

func (s *StaffManageService) DeleteStaff(id string) error {
	global.GVA_DB.Table("hr_employee").Where("id = ?", id).Update("active", fmt.Sprintf("false"))
	//	var staffInfo response.StaffInfo
	//	global.GVA_DB.Table("hr_employee").Where("id = ?", id).First(&staffInfo)
	//	if bool() {
	//		return errors.New("删除失败")
	//	} else {
	//		return nil
	//	}
	return nil
}
func (s *StaffManageService) BatchDeleteStaff(req []string) error {
	return nil
}

func (s *StaffManageService) CreateStaffInfo(req request2.CreateStaff) error {
	var staffInfo hospitalManage.HrEmployee2
	staffInfo.CreateDate = time.Now()
	staffInfo.WriteDate = time.Now()
	staffInfo.CreateUid = req.EmployeeId
	staffInfo.WriteUid = req.EmployeeId
	for k, v := range req.UserInfo {
		firstChar := []rune(k)[:1]
		upperFirstChar := strings.ToUpper(string(firstChar))
		result := upperFirstChar + string([]rune(k)[1:])
		staffInfo.result = v
	}
}
