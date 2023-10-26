package hospitalManage

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"github.com/gin-gonic/gin"
	"reflect"
)

type StaffManage struct{}

func (s *StaffManage) GetStaffInfoList(c *gin.Context) {
	var PageInfo request.PageInfoOfGet
	if c.ShouldBindQuery(&PageInfo) != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	data := StaffService.GetStaffList(PageInfo)
	response.OkWithDetailed(data, "获取成功", c)
}

func (s *StaffManage) GetStaffInfo(c *gin.Context) {
	employeeId := c.Query("id")
	fmt.Println(employeeId)
	if employeeId == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	data := StaffService.GetStaffInfo(employeeId)
	response.OkWithDetailed(data, "获取成功", c)
}

// UpdateBasicStaffInfo 更新员工基础信息
func (s *StaffManage) UpdateBasicStaffInfo(c *gin.Context) {
	var staffInfo request2.BasicInfo
	err := c.ShouldBindJSON(&staffInfo)
	fmt.Println(staffInfo)
	v := reflect.ValueOf(staffInfo)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		key := t.Field(i).Name
		value := field.Interface()

		fmt.Printf("Key: %s, ----------------Value: %v\n", key, value)
	}
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	fmt.Println(staffInfo)
	err1 := StaffService.UpdateBasicStaffInfo(staffInfo)
	if err1 != nil {
		response.FailWithMessage("更新失败", c)
	}
	response.OkWithMessage("更新成功", c)
}
func (s *StaffManage) DeleteStaff(c *gin.Context) {
	employeeId := c.Params.ByName("id")
	if employeeId == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := StaffService.DeleteStaff(employeeId)
	if err1 != nil {
		response.FailWithMessage("删除失败", c)
	}
	response.OkWithMessage("删除成功", c)
}
func (s *StaffManage) BatchDeleteStaff(c *gin.Context) {
}
func (s *StaffManage) CreateStaff(c *gin.Context) {
}
func (s *StaffManage) BatchCreateStaff(c *gin.Context) {
}
