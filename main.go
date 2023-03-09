package main

import (
	"github.com/xarick/gin-casbin-project/models"
	"github.com/xarick/gin-casbin-project/routes"
)

func main() {

	db, _ := models.ConnectDB()
	routes.SetupRoutes(db)
}
