package MsCustomerRoute

import (
	MsCustomerController "github.com/My-Dios/go-legends/modules/v1/controllers/ms-customer"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	router.GET("api/v1/customer", MsCustomerController.Index)
	router.GET("api/v1/customer/show/:id", MsCustomerController.Show)
	router.POST("api/v1/customer/create", MsCustomerController.Create)
	router.PUT("api/v1/customer/update/:id", MsCustomerController.Update)
	router.DELETE("api/v1/customer/delete/:id", MsCustomerController.Delete)
}
