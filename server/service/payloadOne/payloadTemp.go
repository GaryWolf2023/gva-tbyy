package payloadOne

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	testpayloadone "github.com/flipped-aurora/gin-vue-admin/server/model/testPayloadOne"
	"github.com/flipped-aurora/gin-vue-admin/server/model/testPayloadOne/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type PayloadTempService struct{}

func (p *PayloadTempService) CreatePayloadTemp(req request.CreateTemp) (object any, err error) {
	var createTemp testpayloadone.PayloadTemplate
	var utilsMinio utils.MinioOperate
	//fmt.Println(req.TempFile)

	createTemp.TempVersion = "0.0.1"
	createTemp.TempName = req.TempName
	createTemp.TempType = req.TempType
	createTemp.TempContentType = req.TempContentType
	createTemp.Active = req.Active
	createTemp.CreatePerson = req.CreatePerson
	createTemp.TempContent = ""
	uploadInfo, err := utilsMinio.SaveFileToMinio("EMR", fmt.Sprintf("EMR/电子病历/%s/", req.TempType), fmt.Sprintf("%s.html", req.TempName))
	fmt.Println(uploadInfo)
	if err != nil {
		return nil, err
	}
	global.GVA_DB.Create(&createTemp)
	return uploadInfo, nil
}

func (p *PayloadTempService) UpdatePayloadTemp(req request.UpdateTemp) error {
	//var Temp testpayloadone.PayloadTemplate
	return nil
}

//func (p *PayloadTempService) GetPayloadTempList(req *request.GetPayloadList) (data []response.GetTemp, err error) {
//	return [], nil
//}

func (p *PayloadTempService) DeletePayloadTemp(req string) error {
	return nil
}
