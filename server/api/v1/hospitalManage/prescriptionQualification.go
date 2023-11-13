package hospitalManage

import (
	"fmt"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"github.com/gin-gonic/gin"
)

type PrescriptionQualification struct{}

// 处方资质管理--------包含处方资质管理----------个人处方资质管理

func (p *PrescriptionQualification) CreatePQ(c *gin.Context) {
	var createInfo request.CreatePQ
	err := c.ShouldBindJSON(&createInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := PPQService.CreatePQ(createInfo)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithMessage("获取成功", c)
}

func (p *PrescriptionQualification) UpdatePQ(c *gin.Context) {
	var updateInfo request.UpdatePQ
	err := c.ShouldBindJSON(&updateInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := PPQService.UpdatePQ(updateInfo)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (p *PrescriptionQualification) DeletePQ(c *gin.Context) {
	var deleteInfo request.DeletePQ
	err := c.ShouldBindQuery(&deleteInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := PPQService.DeletePQ(deleteInfo)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (p *PrescriptionQualification) GetListPQ(c *gin.Context) {
	var list request2.PageInfo
	err := c.ShouldBindQuery(&list)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1, data := PPQService.GetListPQ(list)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

func (p *PrescriptionQualification) AddPersonPQItem(c *gin.Context) {
	var info request.AddPPQ
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := PPQService.AddPPQ(info)
	if err1 != nil {
		response.FailWithMessage("新增失败", c)
		return
	}
	response.OkWithMessage("获取失败", c)
}

func (p *PrescriptionQualification) GetPersonPQList(c *gin.Context) {
	var id request.GetListPPQ
	err := c.ShouldBindQuery(&id)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1, data := PPQService.GetListPPQ(id)
	if err1 != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

func (p *PrescriptionQualification) DeletePersonPQ(c *gin.Context) {
	var id request.DeletePPQ
	err := c.ShouldBindQuery(&id)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := PPQService.DeletePPQ(id)
	if err1 != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithMessage("获取成功", c)
}
