package hospitalManage

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AddressManageRouter struct{}

func (a *AddressManageRouter) InitAddressManageRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	addressManageRouter := Router.Group("address").Use(middleware.OperationRecord())
	addressManageApi := v1.ApiGroupApp.HospitalManageApiGroup.AddressManageApi
	{
		addressManageRouter.POST("create", addressManageApi.CreateAddress)
		addressManageRouter.POST("update", addressManageApi.UpdateAddress)
		addressManageRouter.DELETE("delete", addressManageApi.DeleteAddress)
		addressManageRouter.GET("getList", addressManageApi.GetAddressList)
		addressManageRouter.GET("getInfo", addressManageApi.GetAddressInfo)
	}
}
