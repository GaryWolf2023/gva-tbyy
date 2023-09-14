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
	return nil
}

func (p *PayloadService) UpdatePayloadOne(payload request.UpdatePayload) error {
	var updatePayload testpayloadone.PayloadDoc
	//先查找数据
	//再将数据中的字段更换然后在update
	tx := global.GVA_DB.Where("payload_id = ?", payload.PayloadId).First(&updatePayload)
	fmt.Println(tx)
	return nil
}
func (p *PayloadService) DeletePayloadOne(reqBody request.DeletePayload) error {
	var resData response.PayloadDoc
	err := global.GVA_DB.Where("payload_id = ?", reqBody.PayloadId).Find(&resData).Limit(1)
	if err != nil {
		return err.Error
	} else {
		return nil
	}

}

// 根据payloadID获取payload
func (p *PayloadService) GetPayloadById(reqBody request.GetPayloadById) (res string) {
	var resData response.PayloadDoc
	global.GVA_DB.Where("payload_id = ?", reqBody.PayloadId).First(&resData)
	return ""
}

// 根据JobIdpayload
func (p *PayloadService) GetPayloadList(reqBody request.GetPayloadList) (res string) {
	var resData []response.PayloadDoc
	if reqBody.JobId != "" {
		global.GVA_DB.Table("payload_doc").Where("JobId = ?", reqBody.JobId).Find(&resData)
		fmt.Println(resData)
		return ""
	} else {
		global.GVA_DB.Table("payload_doc").Where("Unique = ?", reqBody.Unique).Find(&resData)
		return ""
	}
}
