package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
)

type SysTemMinioApi struct{}

func (s *SysTemMinioApi) UpdateFileToMinio(c *gin.Context) {
	var updateFile system.UpdateFile
	updateFile.BucketName = c.PostForm("bucketName")
	updateFile.ObjectName = c.PostForm("objectName")

	fmt.Println(updateFile)
	header, err1 := c.FormFile("file")
	if err1 != nil {
		response.FailWithMessage("没有文件上传", c)
		return
	}
	fmt.Println(header.Filename)
	minioService.UpdateFile(updateFile, header, header.Filename)
}
func (s *SysTemMinioApi) GetFileToMinio(c *gin.Context) {
	var getFile system.GetFile
	err := c.ShouldBindQuery(&getFile)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	data, err1 := minioService.GetFile(getFile)
	if err1 != nil {
		response.FailWithMessage("下载失败", c)
		return
	}
	response.OkWithDetailed(data, "OK", c)
}
func (s *SysTemMinioApi) DeleteToMinio(c *gin.Context) {
	var deleteFile system.DeleteFile
	err := c.ShouldBindQuery(&deleteFile)
	if err != nil {
		response.FailWithMessage("参数错误", c)
	}
	err1 := minioService.DeleteFile(deleteFile)
	if err1 != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
