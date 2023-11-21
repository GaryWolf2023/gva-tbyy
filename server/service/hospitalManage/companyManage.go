package hospitalManage

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	response2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/hospitalManage/request"
)

type CompanyManageService struct{}

func (cms *CompanyManageService) CreateCompany(req request.CreateCompany) error {
	return nil
}

func (cms *CompanyManageService) UpdateCompany(req request.UpdateCompany) error {
	return nil
}

// GetCompanyList 需不需要将检索的keyword设置为函数参数，
func (cms *CompanyManageService) GetCompanyList(req request2.PageInfo) (error, response2.PageRes) {
	var infoList []hospitalManage.ResCompany
	var resInfo response2.PageRes
	var total int64
	var offset int
	var limit int
	if req.PageSize == 0 {
		limit = 9999999
		offset = 0
	} else {
		limit = req.PageSize
		offset = (req.Page - 1) * req.PageSize
	}
	result := global.GVA_DB.Table("res_company").Where("name LIKE ?", fmt.Sprintf("%%%s%%", req.Keyword)).Offset(offset).Limit(limit).Find(&infoList).Count(&total)
	if result.Error != nil {
		return result.Error, response2.PageRes{}
	}
	resInfo.Total = total
	resInfo.Data = infoList
	return nil, resInfo
}

func (cms *CompanyManageService) DeleteCompany(req request.DeleteCompany) error {
	result := global.GVA_DB.Table("res_company").Unscoped().Delete(&req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
