package hospitalManage

import (
	"fmt"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"github.com/gin-gonic/gin"
)

type JobTitleApi struct{}

func (j *JobTitleApi) GetJobTitleList(c *gin.Context) {
	var commonList request2.PageInfo
	err1 := c.ShouldBindQuery(&commonList)
	if err1 != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err2, list := JobTitleService.GetJobTitleList(commonList)
	if err2 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err2), c)
		return
	}
	response.OkWithDetailed(list, "获取成功", c)
}

func (j *JobTitleApi) CreateJobTitle(c *gin.Context) {
	var createInfo request.CreateJobTitle
	err := c.ShouldBindJSON(&createInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := JobTitleService.CreateJobTitle(createInfo)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithMessage("创建成功", c)
	return
}

func (j *JobTitleApi) UpdateJobTitle(c *gin.Context) {
	var updateInfo request.UpdateJobTitle
	err := c.ShouldBindJSON(&updateInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := JobTitleService.UpdateJobTitle(updateInfo)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithMessage("更新成功", c)
	return
}

func (j *JobTitleApi) DeleteJobTitle(c *gin.Context) {
	var deleteId request.DeleteJobTitle
	err := c.ShouldBindQuery(&deleteId)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := JobTitleService.DeleteJobTitle(deleteId)
	if err1 != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err1), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
