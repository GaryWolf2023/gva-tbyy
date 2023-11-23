package hospitalManage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	response2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
	"strconv"
	"time"
)

type AddressManageService struct{}

func (a *AddressManageService) CreateAddress(req request.CreateAddress) error {
	var info hospitalManage.ResPartner
	result := global.GVA_DB.Where("").First(&info)
	if result.Error != nil {

	}
	info.CreateUid = req.EmployeeId
	info.WriteUid = req.EmployeeId
	info.CreateDate = time.Now()
	info.WriteDate = time.Now()
	info.DisplayName = req.DisplayName
	info.Phone = req.Phone
	info.Email = req.Email
	info.UserId = req.UserId
	info.City = strconv.Itoa(req.City)
	info.CountryId = req.CountryId
	info.CompanyId = req.CompanyId
	info.StateId = req.StateId
	info.Vat = strconv.Itoa(req.Vat)

	return nil
}

func (a *AddressManageService) UpdateAddress(req request.UpdateAddress) error {
	return nil
}

func (a *AddressManageService) DeleteAddress(req request.DeleteAddress) error {
	return nil
}

func (a *AddressManageService) GetAddressList(req request2.PageInfo) (error, response2.PageRes) {

	return nil, response2.PageRes{}
}

func (a *AddressManageService) GetAddressInfo(id int) (error, interface{}) {
	return nil, nil
}
