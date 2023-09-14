package testPayload

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/testPayloadOne/request"
	"github.com/gin-gonic/gin"
)

type TestPayloadOne struct {
}

// CreatePayload
// @Tags      TestPayloadOne
// @Summary   创建payload
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      testPayload.TestPayloadOne                  true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "创建payload成功"
// @Router    /payload/createPayload [post]
func (t *TestPayloadOne) CreatePayload(c *gin.Context) {
	payloadOne := request.CreatePayload{}
	err1 := c.ShouldBindJSON(&payloadOne)
	if err1 != nil {
		response.FailWithMessage("创建失败,参数错误", c)
	}
	err2 := PayloadService.CreatePayloadOne(payloadOne)
	if err2 != nil {
		response.FailWithMessage("创建失败", c)
	}
	response.Ok(c)
}

// UpdatePayload
// @Tags      TestPayloadOne
// @Summary   更新payload
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      testPayload.TestPayloadOne                  true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "操作成功"
// @Router    /payload/updatePayload [post]
func (t *TestPayloadOne) UpdatePayload(c *gin.Context) {
	payload := request.UpdatePayload{}
	err1 := c.ShouldBindJSON(&payload)
	if err1 != nil {
		response.FailWithMessage("参数错误", c)
	}
	err2 := PayloadService.UpdatePayloadOne(payload)
	if err2 != nil {
		response.FailWithMessage("更新失败", c)
	}
	response.Ok(c)
}

// DeletePayload
// @Tags      TestPayloadOne
// @Summary   删除payload（封存）
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      testPayload.TestPayloadOne                  true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "操作成功"
// @Router    /payload/deletePayload [post]
func (t *TestPayloadOne) DeletePayload(c *gin.Context) {
	payload := request.DeletePayload{}
	err1 := c.ShouldBindJSON(&payload)
	if err1 != nil {
		response.FailWithMessage("参数错误", c)
	}
	err2 := PayloadService.DeletePayloadOne(payload)
	if err2 != nil {
		response.FailWithMessage("更新失败", c)
	} else {
		response.Ok(c)
	}
}

// GetPayloadById
// @Tags      TestPayloadOne
// @Summary   获取某条payload
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      testPayload.TestPayloadOne                  true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "操作成功"
// @Router    /payload/getPayloadById [post]
func (t *TestPayloadOne) GetPayloadById(c *gin.Context) {
	var payload request.GetPayloadById
	err1 := c.ShouldBindJSON(&payload)
	if err1 != nil {
		response.FailWithMessage("参数错误", c)
	}
	data := PayloadService.GetPayloadById(payload)
	response.OkWithData(data, c)
}

// GetPayloadList
// @Tags      TestPayloadOne
// @Summary   获取同一个job_id或者unique的payload
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      testPayload.TestPayloadOne                  true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "操作成功"
// @Router    /payload/getPayloadList [post]
func (t *TestPayloadOne) GetPayloadList(c *gin.Context) {
	var payload request.GetPayloadList
	err1 := c.ShouldBindJSON(&payload)
	if err1 != nil {
		response.FailWithMessage("参数错误", c)
	}
	data := PayloadService.GetPayloadList(payload)
	response.OkWithData(data, c)
}
