package participant

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type GetParticipantApi struct{}

// GetParticipantsByType
// @Tags      GetParticipantApi
// @Summary   获取参与人列表（可检索）
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      participant.GetParticipantApi     true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "创建payload成功"
// @Router    /payload/createPayload [post]
func (g *GetParticipantApi) GetParticipantsByType(c *gin.Context) {
	err := GetPatService.GetParticipantsByType()
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
}

// GetParticipantsBySearch
// @Tags      GetParticipantApi
// @Summary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      participant.GetParticipantApi     true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "创建payload成功"
// @Router    /payload/createPayload [post]
func (g *GetParticipantApi) GetParticipantsBySearch(c *gin.Context) {
	err := GetPatService.GetParticipantsByType()
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
}
