package main

import (
	"os"

	"github.com/xarick/gin-casbin-project/models"
	"github.com/xarick/gin-casbin-project/routes"
)

func init() {
	os.Setenv("JWT_SECRET", "123456789")
}

func main() {
	db, _ := models.ConnectDB()
	routes.SetupRoutes(db)
}
