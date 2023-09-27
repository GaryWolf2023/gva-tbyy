package payloadOne

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	testpayloadone "github.com/flipped-aurora/gin-vue-admin/server/model/testPayloadOne"
	"github.com/flipped-aurora/gin-vue-admin/server/model/testPayloadOne/request"
)

type PayloadTempService struct{}

func (p *PayloadTempService) CreatePayloadTemp(req request.CreateTemp) error {
	var createTemp testpayloadone.PayloadTemplate
	fmt.Println(req.TempFile)

	createTemp.TempVersion = "0.0.1"
	createTemp.TempName = req.TempName
	createTemp.TempType = req.TempType
	createTemp.TempContentType = req.TempContentType
	createTemp.Active = req.Active
	createTemp.CreatePerson = req.CreatePerson
	createTemp.TempContent = ""
	global.GVA_DB.Create(&createTemp)
	return nil
}

func (p *PayloadTempService) UpdatePayloadTemp(req request.UpdateTemp) error {
	return nil
}

//func (p *PayloadTempService) GetPayloadTempList(req *request.GetPayloadList) (data []response.GetTemp, err error) {
//	return [], nil
//}

func (p *PayloadTempService) DeletePayloadTemp(req string) error {
	return nil
}
