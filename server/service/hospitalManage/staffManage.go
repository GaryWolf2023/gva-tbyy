package hospitalManage

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	response2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/response"
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

func (s *StaffManageService) GetStaffInfo(id string) response.BasicStaffInfo {
	var staffInfo response.BasicStaffInfo
	global.GVA_DB.Table("hr_employee").Where("id = ?", id).First(&staffInfo)
	return staffInfo
}

//func (s *StaffManageService) UpdateStaff(req request2.BasicStaffInfo) error {
//	fmt.Println(req)
//	result := global.GVA_DB.Table("hr_employee").Where("id = ?", req.ID).Updates(req)
//	if result.Error != nil {
//		return errors.New(fmt.Sprintf("更新失败%v", result.Error))
//	}
//	return nil
//}

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
