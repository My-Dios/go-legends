package route

import (
	MsCustomerRoute "github.com/My-Dios/go-legends/routes/ms-customer"
	MsItemRoute "github.com/My-Dios/go-legends/routes/ms-item"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	MsItemRoute.Route(router)
	MsCustomerRoute.Route(router)
}
