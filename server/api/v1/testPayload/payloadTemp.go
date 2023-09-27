package testPayload

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/testPayloadOne/request"
	"github.com/gin-gonic/gin"
)

type PayloadTemplate struct{}

func (p *PayloadTemplate) CreatePayloadTemp(c *gin.Context) {
	var createTemp request.CreateTemp
	//file就是我们上传的文件
	err := c.ShouldBind(&createTemp)

	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	err1 := PayloadTempService.CreatePayloadTemp(createTemp)
	if err1 != nil {
		response.FailWithMessage("上传模板失败", c)
		return
	}
	response.OkWithMessage("上传模板成功", c)
}
func (p *PayloadTemplate) UpdatePayloadTemp(c *gin.Context) {
	var updateTemp request.UpdateTemp
	err := c.ShouldBind(&updateTemp)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := PayloadTempService.UpdatePayloadTemp(updateTemp)
	if err1 != nil {
		response.FailWithMessage("更新模板失败", c)
		return
	}
	response.OkWithMessage("更新模板成功", c)
}
func (p *PayloadTemplate) DeletePayloadTemp(c *gin.Context) {
	tempId := c.Query("id")
	fmt.Println(tempId)
	if tempId == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	err := PayloadTempService.DeletePayloadTemp(tempId)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
func (p *PayloadTemplate) GetPayloadTempList(c *gin.Context) {
	var tempList request.GetTempList
	err := c.ShouldBind(&tempList)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
}
func (p *PayloadTemplate) GetPayloadTemp(c *gin.Context) {

}
