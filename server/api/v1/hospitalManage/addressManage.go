package hospitalManage

import (
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"github.com/gin-gonic/gin"
)

// 地址管理，员工工作地址管理（这哪是没理解是什么意思，这个工作地址）

type AddressManageApi struct{}

func (a *AddressManageApi) CreateAddress(c *gin.Context) {
	var createInfo request.CreateAddress
	err := c.ShouldBindJSON(&createInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err = AddressManageService.CreateAddress(createInfo)
	if err != nil {
		response.FailWithMessage("创建失败", c)
		return
	}
	response.FailWithMessage("创建成功", c)
}

func (a *AddressManageApi) UpdateAddress(c *gin.Context) {
	var updateInfo request.UpdateAddress
	err := c.ShouldBindJSON(&updateInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err = AddressManageService.UpdateAddress(updateInfo)
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	response.FailWithMessage("更新成功", c)
}

func (a *AddressManageApi) DeleteAddress(c *gin.Context) {
	var deleteInfo request.DeleteAddress
	err := c.ShouldBindQuery(&deleteInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err = AddressManageService.DeleteAddress(deleteInfo)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.FailWithMessage("删除成功", c)
}

func (a *AddressManageApi) GetAddressList(c *gin.Context) {
	var pageInfo request2.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1, data := AddressManageService.GetAddressList(pageInfo)
	if err1 != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.FailWithDetailed(data, "获取成功", c)
}

func (a *AddressManageApi) GetAddressInfo(c *gin.Context) {
	var id int
	err := c.ShouldBindQuery(&id)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1, data := AddressManageService.GetAddressInfo(id)
	if err1 != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.FailWithDetailed(data, "获取成功", c)
}
