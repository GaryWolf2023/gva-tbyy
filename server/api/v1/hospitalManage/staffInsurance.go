package hospitalManage

import (
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"github.com/gin-gonic/gin"
)

type StaffInsuranceApi struct{}

func (s *StaffInsuranceApi) CreateInsurance(c *gin.Context) {
	var create request.CreateInsurance
	err1 := c.ShouldBindJSON(&create)
	if err1 != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err2 := InsuranceService.CreateInsurance(create)
	if err2 != nil {
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (s *StaffInsuranceApi) UpdateInsurance(c *gin.Context) {
	var update request.UpdateInsurance
	err1 := c.ShouldBindJSON(&update)
	if err1 != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err2 := InsuranceService.UpdateInsurance(update)
	if err2 != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (s *StaffInsuranceApi) DeleteInsurance(c *gin.Context) {
	var delete1 request.DeleteInsurance
	err1 := c.ShouldBindQuery(&delete1)
	if err1 != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err2 := InsuranceService.DeleteInsurance(delete1)
	if err2 != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (s *StaffInsuranceApi) GetInsuranceList(c *gin.Context) {
	var list request2.PageInfo
	err1 := c.ShouldBindQuery(&list)
	if err1 != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err2, data := InsuranceService.GetInsuranceList(list)
	if err2 != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

func (s *StaffInsuranceApi) GetInsuranceInfo(c *gin.Context) {
	var info request.GetInsuranceInfo
	err1 := c.ShouldBindQuery(&info)
	if err1 != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err2, data := InsuranceService.GetInsuranceInfo(info)
	if err2 != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

func (s *StaffInsuranceApi) GetPersonInsuranceList(c *gin.Context) {
	var insuranceList request.GetPersonInsuranceList
	err1 := c.ShouldBindQuery(&insuranceList)
	if err1 != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err2, data := InsuranceService.GetPersonInsuranceList(insuranceList)
	if err2 != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}
func (s *StaffInsuranceApi) AddPersonInsurance(c *gin.Context) {
	var addInfo request.AddPersonInsurance
	err1 := c.ShouldBindJSON(&addInfo)
	if err1 != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err2 := InsuranceService.AddPersonInsurance(addInfo)
	if err2 != nil {
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}
func (s *StaffInsuranceApi) DeletePersonInsurance(c *gin.Context) {
	var deleteInfo request.DeletePersonInsurance
	err1 := c.ShouldBind(&deleteInfo)
	if err1 != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err2 := InsuranceService.DeletePersonInsurance(deleteInfo)
	if err2 != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
