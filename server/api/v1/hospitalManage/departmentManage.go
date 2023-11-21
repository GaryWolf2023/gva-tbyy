package hospitalManage

import (
	"fmt"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"github.com/gin-gonic/gin"
)

type DepartmentManageApi struct{}

func (d *DepartmentManageApi) GetDepList(c *gin.Context) {
	var pageRes request2.PageInfo
	err := c.ShouldBindQuery(&pageRes)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1, data := DepartmentManageService.GetList(pageRes)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}
func (d *DepartmentManageApi) CreateDep(c *gin.Context) {
	var createInfo request.CreateInfo
	err := c.ShouldBindJSON(&createInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := DepartmentManageService.CreateDep(createInfo)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
func (d *DepartmentManageApi) UpdateDep(c *gin.Context) {
	var updateInfo request.UpdateInfo
	err := c.ShouldBindJSON(&updateInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := DepartmentManageService.UpdateDep(updateInfo)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
func (d *DepartmentManageApi) DeleteDep(c *gin.Context) {
	var deleteInfo request.DeleteInfo
	err := c.ShouldBindQuery(&deleteInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := DepartmentManageService.DeleteDep(deleteInfo)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
func (d *DepartmentManageApi) BatchesOfDelete(c *gin.Context) {
	var batchesOfDelete request.BatchesOfDelete
	err := c.ShouldBindJSON(&batchesOfDelete)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := DepartmentManageService.BatchesOfDelete(batchesOfDelete)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
