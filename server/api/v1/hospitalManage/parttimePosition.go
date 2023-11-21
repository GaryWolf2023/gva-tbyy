package hospitalManage

import (
	"fmt"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"github.com/gin-gonic/gin"
)

type ManagePartTimePositionApi struct{}

func (m *ManagePartTimePositionApi) GetPartTimePositionList(c *gin.Context) {
	var list request.GetItemList
	err := c.ShouldBindQuery(&list)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1, data := PartTimePositionService.GetPartTimePositionList(list)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

func (m *ManagePartTimePositionApi) AddPartTimePosition(c *gin.Context) {
	var info request.AddItem
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := PartTimePositionService.AddPartTimePosition(info)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithMessage("获取成功", c)
}

func (m *ManagePartTimePositionApi) DeletePartTimePosition(c *gin.Context) {
	var info request.DeleteItem
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := PartTimePositionService.DeletePartTimePosition(info)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

//==================================================================================
// 下面是job管理
//==================================================================================

func (m *ManagePartTimePositionApi) GetJobList(c *gin.Context) {
	var pageInfo request2.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1, data := PartTimePositionService.GetJobList(pageInfo)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

func (m *ManagePartTimePositionApi) CreateJob(c *gin.Context) {
	var createInfo request.CreateJob
	err := c.ShouldBindJSON(&createInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err = PartTimePositionService.CreateJob(createInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (m *ManagePartTimePositionApi) UpdateJob(c *gin.Context) {
	var upd request.UpdateJob
	err := c.ShouldBindJSON(&upd)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err = PartTimePositionService.UpdateJob(upd)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (m *ManagePartTimePositionApi) DeleteJob(c *gin.Context) {
	var del request.DeleteJob
	err := c.ShouldBindQuery(&del)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err = PartTimePositionService.DeleteJob(del)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
