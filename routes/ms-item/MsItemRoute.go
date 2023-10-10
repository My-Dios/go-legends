package MsItemRoute

import (
	MsItemController "github.com/My-Dios/go-legends/modules/v1/controllers/ms-item"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	router.GET("api/v1/item", MsItemController.Index)
	router.GET("api/v1/item/show/:id", MsItemController.Show)
	router.POST("api/v1/item/create", MsItemController.Create)
	router.PUT("api/v1/item/update/:id", MsItemController.Update)
	router.DELETE("api/v1/item/delete/:id", MsItemController.Delete)
}
