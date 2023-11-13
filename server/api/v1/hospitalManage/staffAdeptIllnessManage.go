package hospitalManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"github.com/gin-gonic/gin"
)

type StaffAdeptIllnessApi struct{}

// GetAdeptIllnessList 根据员工ID获取员工擅长疾病列表
func (s *StaffAdeptIllnessApi) GetAdeptIllnessList(c *gin.Context) {
	var getIllnessList request.GetIllnessList
	err := c.ShouldBindQuery(&getIllnessList)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1, data := AdeptIllnessService.GetAdeptIllnessList(getIllnessList)
	if err1 != nil {
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

// AddAdeptIllnessItem 新增擅长疾病
func (s *StaffAdeptIllnessApi) AddAdeptIllnessItem(c *gin.Context) {
	var adeptIllness request.AddIllness
	err := c.ShouldBindJSON(&adeptIllness)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := AdeptIllnessService.AddAdeptIllness(adeptIllness)
	if err1 != nil {
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// DeleteAdeptIllnessItem 删除相关擅长疾病
func (s *StaffAdeptIllnessApi) DeleteAdeptIllnessItem(c *gin.Context) {
	var deleteId request.DeleteIllness
	err := c.ShouldBindQuery(&deleteId)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := AdeptIllnessService.DeleteAdeptIllnessItem(deleteId)
	if err1 != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
