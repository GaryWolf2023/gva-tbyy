package participant

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/participant/request"
	"github.com/gin-gonic/gin"
)

type TempManageParticipant struct{}

func (t *TempManageParticipant) CreateTempOfParticipant(c *gin.Context) {
	var createTemp request.CreateTempOfPat
	err := c.ShouldBindJSON(&createTemp)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := PatTempService.CreatePatTemp(createTemp)
	if err1 != nil {
		response.FailWithMessage("服务器错误，创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
func (t *TempManageParticipant) DeleteTempOfParticipant(c *gin.Context) {
	var deleteTemp request.DeleteTempOfPat
	err := c.ShouldBind(&deleteTemp)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := PatTempService.DeletePatTemp(deleteTemp)
	if err1 != nil {
		response.FailWithMessage("服务器错误，删除失败，请联系管理员", c)
		return
	}
}
func (t *TempManageParticipant) UpdateTempOfParticipant(c *gin.Context) {
	var updateTemp request.UpdateTempOfPat
	err := c.ShouldBind(&updateTemp)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := PatTempService.UpdatePatTemp(updateTemp)
	if err1 != nil {
		response.FailWithMessage("服务器错误，更新失败，请联系管理员", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
func (t *TempManageParticipant) GetTempOfParticipant(c *gin.Context) {
	var getTemp request.GetTempOfPat
	err := c.ShouldBind(&getTemp)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1, data := PatTempService.GetPatTemp(getTemp)
	if err1 != nil {
		response.FailWithMessage("服务器错误，获取失败，请联系管理员", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}
