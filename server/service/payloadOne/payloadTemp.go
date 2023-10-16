package payloadOne

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	testpayloadone "github.com/flipped-aurora/gin-vue-admin/server/model/testPayloadOne"
	"github.com/flipped-aurora/gin-vue-admin/server/model/testPayloadOne/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/testPayloadOne/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"strconv"
	"strings"
)

type PayloadTempService struct{}

func (p *PayloadTempService) CreatePayloadTemp(req request.CreateTemp, filename string, file *multipart.FileHeader) (object any, err error) {
	var createTemp testpayloadone.PayloadTemplate
	var utilsMinio utils.MinioOperate
	//fmt.Println(req.TempFile)
	createTemp.TempVersion = 1
	createTemp.TempName = req.TempName
	createTemp.TempType = req.TempType
	createTemp.TempContentType = req.TempContentType
	createTemp.Active = req.Active
	createTemp.CreatePerson = req.CreatePerson

	f, err0 := file.Open()
	if err0 != nil {
		return nil, err0
	}
	minioInfo, err1 := utilsMinio.SaveFileToMinio("EMR", fmt.Sprintf("电子病历/%s/%s", req.TempType, filename), filename, f, file.Size)

	if err1 != nil {
		fmt.Println("向minio存储文件失败")
		global.GVA_LOG.Error("向minio存储文件失败")
		return nil, err1
	}
	minioInfoJson, _ := json.Marshal(minioInfo)
	createTemp.TempContent = string(minioInfoJson)
	err2 := global.GVA_DB.AutoMigrate(&createTemp)
	if err2 != nil {
		return nil, err2
	}
	global.GVA_DB.Create(&createTemp)
	return nil, nil
}

func (p *PayloadTempService) UpdatePayloadTemp(req request.UpdateTemp) error {
	var Temp testpayloadone.PayloadTemplate
	fmt.Println(req)
	global.GVA_DB.Table("payload_templates").Where("id = ?", req.TempID).Find(&Temp)
	global.GVA_DB.Model(&Temp).Updates(map[string]interface{}{
		"temp_name":         req.TempName,
		"temp_content_type": req.TempContentType,
		"temp_type":         req.TempType,
		"active":            req.Active,
		"update_person":     req.UpdatePerson,
	})
	return nil
}

func (p *PayloadTempService) DeletePayloadTemp(id string) error {
	global.GVA_DB.Table("payload_templates").Where("id = ?", id).Update("active", false)
	return nil
}

func (p *PayloadTempService) GetPayloadTempList(req request.GetTempList) ([]response.GetTemp, error) {
	//var timeFormat = timer.InitTimer{}
	page := req.Page
	pageSize := req.PageSize
	if pageSize == 0 {
		pageSize = 20
	}
	if page == 0 {
		page = 1
	}
	limit := pageSize
	offset := pageSize * (page - 1)
	var resData []response.GetTemp
	fmt.Println("---------------------------------------------------------------")
	if req.SearchData != "" {
		global.GVA_DB.Table("payload_templates").Offset(offset).Limit(limit).Where("temp_name LIKE ? AND active = ? AND temp_type = ?", fmt.Sprintf("%%%s%%", req.SearchData), true, req.Type).Find(&resData)
	} else {
		global.GVA_DB.Table("payload_templates").Offset(offset).Limit(limit).Where("active = ? AND temp_type = ?", true, req.Type).Find(&resData)
	}
	fmt.Println(resData)
	return resData, nil
}

func (p *PayloadTempService) GetPayloadTemplate(id string) (res response.GetTemp, err error) {
	var resData response.GetTemp
	global.GVA_DB.Table("payload_templates").Where("id = ? AND active = ?", id, true).Find(&resData)
	return resData, nil
}

func (p *PayloadTempService) UpdatePayloadOfFile(req request.UpdateOfFile, file *multipart.FileHeader) error {
	var utilsMinio utils.MinioOperate
	var hisTemp testpayloadone.HistoryTemplate
	size := file.Size
	f, err := file.Open()
	if err != nil {
		return err
	}
	minioInfo, err1 := utilsMinio.SaveFileToMinio(req.BucketName, req.ObjectName, file.Filename, f, size)
	if err1 != nil {
		return err
	}
	minioInfoJson, _ := json.Marshal(minioInfo)
	id, _ := strconv.Atoi(req.ID)
	hisTemp.TempID = int8(id)
	hisTemp.TempContent = req.OldContent
	hisTemp.TempVersion = req.Version
	hisTemp.PayloadTemplate = req.PayloadTemplate
	fmt.Println(hisTemp)
	fmt.Println(string(minioInfoJson))
	global.GVA_DB.Table("payload_templates").Where("id = ?", id).Updates(map[string]interface{}{"temp_content": string(minioInfoJson), "temp_version": req.Version + 1})
	_ = global.GVA_DB.AutoMigrate(&hisTemp)
	global.GVA_DB.Table("history_templates").Create(&hisTemp)
	fmt.Println(minioInfo)
	return nil
}

func (p *PayloadTempService) GetFileOfTemp(req request.GetTempFile) (*minio.Object, string, error) {
	var utilsMinio utils.MinioOperate
	var temp testpayloadone.PayloadTemplate
	fmt.Println(req)
	global.GVA_DB.Table("payload_templates").Where("temp_name = ? AND temp_type = ?", req.TempName, req.TempType).First(&temp)
	var getFile request.TempContent
	fmt.Println(temp)
	err := json.Unmarshal([]byte(temp.TempContent), &getFile)
	if err != nil {
		return nil, "", err
	}
	fmt.Println(getFile)
	parts := strings.Split(getFile.Key, "/")
	fileName := parts[len(parts)-1]
	fileData, err1 := utilsMinio.GetFileFromMinio(getFile.Bucket, getFile.Key)
	if err1 != nil {
		return nil, fileName, err1
	}
	return fileData, fileName, nil
}

func (p *PayloadTempService) GetFile(req request.GetFile) (*minio.Object, error) {
	var utilsMinio utils.MinioOperate
	fileData, err1 := utilsMinio.GetFileFromMinio(req.BucketName, req.ObjectName)
	if err1 != nil {
		return nil, err1
	}
	return fileData, nil
}
