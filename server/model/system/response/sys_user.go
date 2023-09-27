package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type PersonInfo struct {
	Name        string `json:"name"`
	Sex         string `json:"sex"`
	Age         string `json:"age"`
	IDCard      string `json:"id_card"`
	PhoneNumber string `json:"phone_number"`
}

type StaffInfo struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	JobTitle         string `json:"job_title"`
	PerformanceClass string `json:"performanceclass"`
}

type OneStaffInfo struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	JobTitle         string `json:"job_title"`
	PerformanceClass string `json:"performanceclass"`
	WorkPhone        string `json:"phone"`
	WorkEmail        string `json:"email"`
	Introduction     string `json:"introduction"`
	DepartmentId     string `json:"departmentId"`
}
