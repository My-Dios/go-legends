package main

import (
	migration "github.com/My-Dios/go-legends/migrations"
	route "github.com/My-Dios/go-legends/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	migration.Migrate()
	route.Route(r)
	r.Run(":8080")
}
