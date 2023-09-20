package participant

import "github.com/gin-gonic/gin"

type ParticipantApi struct{}

// CreateParticipant
// @Tags      ParticipantApi
// @Summary   创建参与人
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      participant.ParticipantApi     true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "创建payload成功"
// @Router    /payload/createPayload [post]
func (p *ParticipantApi) CreateParticipant(c *gin.Context) {
	//	 创建这个参与人需要
	//c.ShouldBindJSON()
}

// UpdateParticipant
// @Tags      ParticipantApi
// @Summary   更新参与人信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      participant.ParticipantApi     true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "创建payload成功"
// @Router    /payload/createPayload [post]
func (p *ParticipantApi) UpdateParticipant(c *gin.Context) {

}

// DeleteParticipant
// @Tags      ParticipantApi
// @Summary   删除参与人
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      participant.ParticipantApi     true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "创建payload成功"
// @Router    /payload/createPayload [post]
func (p *ParticipantApi) DeleteParticipant(c *gin.Context) {

}

// FindParticipant
// @Tags      ParticipantApi
// @Summary   查找参与人
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      participant.ParticipantApi     true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "创建payload成功"
// @Router    /payload/createPayload [post]
func (p *ParticipantApi) FindParticipant(c *gin.Context) {

}
