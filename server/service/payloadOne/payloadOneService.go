package payloadOne

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/testPayloadOne/response"

	testpayloadone "github.com/flipped-aurora/gin-vue-admin/server/model/testPayloadOne"
	"github.com/flipped-aurora/gin-vue-admin/server/model/testPayloadOne/request"
)

type PayloadService struct {
}

// CreatePayloadOne 创建payload
func (p *PayloadService) CreatePayloadOne(payload request.CreatePayload) error {
	currentTime := time.Now()
	timestamp := currentTime.UnixNano() / int64(time.Millisecond)

	var payloadDoc = testpayloadone.PayloadDoc{}
	var payloadHeader = testpayloadone.Header{}
	var payloadMeta = testpayloadone.Meta{}
	var payloadContent = testpayloadone.PayloadContent{}

	var jobIdPayload = testpayloadone.JobIdPayload{}

	payloadDoc.Active = true
	payloadDoc.PayloadId = fmt.Sprintf("%v&%s", timestamp, payload.JobId)
	payloadHeader.Business = payload.Business
	payloadHeader.JobId = payload.JobId
	payloadHeader.Unique = payload.Unique
	payloadHeader.PayloadId = fmt.Sprintf("%v&%s", timestamp, payload.JobId)
	payloadMeta.DocName = payload.DocName
	payloadMeta.DocType = payload.DocType
	payloadMeta.Version = "0.0.1"
	if payload.Order != "" {
		payloadMeta.NeedPay = true
		payloadMeta.IsPay = false
	} else {
		payloadMeta.NeedPay = false
		payloadMeta.IsPay = true
	}
	payloadContent.Content = payload.PayloadContent
	payloadContent.Order = payload.Order
	jsonHeader, _ := json.Marshal(payloadHeader)
	payloadDoc.Header = string(jsonHeader)
	jsonMeta, _ := json.Marshal(payloadMeta)
	payloadDoc.Meta = string(jsonMeta)
	jsonPayload, _ := json.Marshal(payloadContent)
	payloadDoc.Payload = string(jsonPayload)
	fmt.Println(payloadDoc)
	err := global.GVA_DB.AutoMigrate(&payloadDoc)
	if err != nil {
		return err
	}
	global.GVA_DB.Create(&payloadDoc)

	jobIdPayload.JobID = payload.JobId
	jobIdPayload.Unique = payload.Unique
	jobIdPayload.PayloadId = fmt.Sprintf("%v&%s", timestamp, payload.JobId)
	err1 := global.GVA_DB.AutoMigrate(&jobIdPayload)
	if err1 != nil {
		return err1
	}
	global.GVA_DB.Create(&jobIdPayload)
	return nil
}

func (p *PayloadService) UpdatePayloadOne(payload request.UpdatePayload) error {
	var updatePayload testpayloadone.PayloadDoc
	var payloadContent testpayloadone.PayloadContent //先查找数据
	//再将数据中的字段更换然后在update
	global.GVA_DB.Table("payload_doc").Where("payload_id = ?", payload.PayloadId).First(&updatePayload)
	fmt.Println(updatePayload)
	err1 := json.Unmarshal([]byte(updatePayload.Payload), &payloadContent)
	if err1 != nil {
		fmt.Println("结构体转换失败")
		return err1
	}
	fmt.Println(payloadContent)
	payloadContent.Content = payload.PayloadContent
	jsonPaylaod, _ := json.Marshal(payloadContent)
	updatePayload.Payload = string(jsonPaylaod)
	global.GVA_DB.Table("payload_doc").Where("active = ? AND payload_id = ?", true, payload.PayloadId).Update("payload", string(jsonPaylaod))
	return nil
}

func (p *PayloadService) DeletePayloadOne(reqBody request.DeletePayload) error {
	//var resData response.PayloadDoc
	fmt.Println(reqBody)
	err := global.GVA_DB.Table("payload_doc").Where("payload_id = ?", reqBody.PayloadId).Update("active", false)
	if err != nil {
		return err.Error
	} else {
		return nil
	}
}

// 根据payloadID获取payload
func (p *PayloadService) GetPayloadById(reqBody request.GetPayloadById) (res response.PayloadDoc) {
	var resData response.PayloadDoc
	fmt.Println("1111111111111111")
	global.GVA_DB.Table("payload_doc").First(&resData, "payload_id = ?", reqBody.PayloadId)
	return resData
}

// 根据JobIdpayload
func (p *PayloadService) GetPayloadList(reqBody request.GetPayloadList) (res []response.PayloadList) {
	page := reqBody.Page
	pageSize := reqBody.PageSize
	if pageSize == 0 {
		pageSize = 10
	}
	if page == 0 {
		page = 1
	}

	limit := pageSize
	offset := pageSize * (page - 1)

	var resData []response.PayloadList
	if reqBody.JobId != "" {
		global.GVA_DB.Table("job_id_payload").Offset(offset).Limit(limit).Where("job_id = ?", reqBody.JobId).Find(&resData)
		fmt.Println(resData)
		return resData
	} else {
		global.GVA_DB.Table("job_id_payload").Offset(offset).Limit(limit).Where("unique = ?", reqBody.Unique).Find(&resData)
		return resData
	}
}
