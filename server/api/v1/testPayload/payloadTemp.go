package testPayload

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/testPayloadOne/request"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
	"strings"
)

type PayloadTemplate struct{}

func (p *PayloadTemplate) CreatePayloadTemp(c *gin.Context) {
	var createTemp request.CreateTemp
	createTemp.CreatePerson = c.PostForm("createPerson")
	createTemp.TempName = c.PostForm("tempName")
	createTemp.TempType = c.PostForm("tempType")
	createTemp.TempContentType = c.PostForm("tempContentType")
	createTemp.Active, _ = strconv.ParseBool(c.PostForm("active"))
	fmt.Println(createTemp)

	header, err1 := c.FormFile("file")
	if err1 != nil {
		response.FailWithMessage("没有文件上传", c)
		return
	}
	fmt.Println(header.Filename)

	obj, err3 := PayloadTempService.CreatePayloadTemp(createTemp, header.Filename, header)
	if err3 != nil {
		response.FailWithMessage("上传模板失败", c)
		return
	}
	response.OkWithDetailed(obj, "上传模板成功", c)
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
	fmt.Println(tempList)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	data, err1 := PayloadTempService.GetPayloadTempList(tempList)
	if err1 != nil {
		response.FailWithMessage("查找失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

func (p *PayloadTemplate) GetPayloadTemp(c *gin.Context) {
	id := c.Query("id")
	fmt.Println(id)
	if id == "" {
		response.FailWithMessage("参数不能为空", c)
		return
	}
	data, err := PayloadTempService.GetPayloadTemplate(id)
	if err != nil {
		response.FailWithMessage("数据获取失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

func (p *PayloadTemplate) UpdatePayloadOfFile(c *gin.Context) {
	var updateOfFile request.UpdateOfFile
	updateOfFile.ID = c.PostForm("id")
	updateOfFile.OldContent = c.PostForm("oldContent")
	updateOfFile.ObjectName = c.PostForm("objectName")
	updateOfFile.BucketName = c.PostForm("bucketName")
	updateOfFile.PayloadTemplate = c.PostForm("payload_template")
	version, _ := strconv.Atoi(c.PostForm("version"))
	updateOfFile.Version = int8(version)
	file, err := c.FormFile("file")
	if err != nil || updateOfFile.ID == "" || updateOfFile.OldContent == "" || updateOfFile.ObjectName == "" || updateOfFile.BucketName == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	err1 := PayloadTempService.UpdatePayloadOfFile(updateOfFile, file)
	if err1 != nil {
		response.FailWithMessage("更新失败，请联系管理员", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (p *PayloadTemplate) GetFileOfTemp(c *gin.Context) {
	var GetTempFile request.GetTempFile
	err := c.ShouldBindQuery(&GetTempFile)
	fmt.Println(GetTempFile)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	fmt.Println(GetTempFile)
	data, fileName, err1 := PayloadTempService.GetFileOfTemp(GetTempFile)
	if err1 != nil {
		response.FailWithMessage("获取文件失败", c)
		return
	}
	localFile, err2 := os.Create(fileName)
	if err2 != nil {
		response.FailWithMessage("创建文件失败", c)
		return
	}
	defer localFile.Close()
	_, err3 := io.Copy(localFile, data)
	if err3 != nil {
		response.FailWithMessage("文件写入失败", c)
		return
	}
	// 设置响应头，包括 Content-Type 和 Content-Disposition
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	c.File(fileName)
}

func (p *PayloadTemplate) GetFile(c *gin.Context) {
	var getFile request.GetFile
	err := c.ShouldBindQuery(&getFile)
	fmt.Println(getFile)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	data, err1 := PayloadTempService.GetFile(getFile)
	if err1 != nil {
		response.FailWithMessage("获取文件失败", c)
		return
	}
	parts := strings.Split(getFile.ObjectName, "/")
	fileName := parts[len(parts)-1]
	localFile, err2 := os.Create(fileName)
	if err2 != nil {
		response.FailWithMessage("创建文件失败", c)
		return
	}
	defer localFile.Close()
	_, err3 := io.Copy(localFile, data)
	if err3 != nil {
		response.FailWithMessage("文件写入失败", c)
		return
	}
	// 设置响应头，包括 Content-Type 和 Content-Disposition
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	c.File(fileName)
}
