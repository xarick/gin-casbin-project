package routes

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/xarick/gin-casbin-project/controller"
	"github.com/xarick/gin-casbin-project/middleware"
	"github.com/xarick/gin-casbin-project/repository"
	"gorm.io/gorm"
)

// SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "OK",
	// 	})
	// })

	// Initialize  casbin adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}

	// Load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	//add policy
	if hasPolicy := enforcer.HasPolicy("doctor", "report", "read"); !hasPolicy {
		enforcer.AddPolicy("doctor", "report", "read")
	}
	if hasPolicy := enforcer.HasPolicy("doctor", "report", "write"); !hasPolicy {
		enforcer.AddPolicy("doctor", "report", "write")
	}
	if hasPolicy := enforcer.HasPolicy("patient", "report", "read"); !hasPolicy {
		enforcer.AddPolicy("patient", "report", "read")
	}

	userRepository := repository.NewUserRepository(db)

	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}

	userController := controller.NewUserController(userRepository)

	apiRoutes := r.Group("/api")

	{
		apiRoutes.POST("/register", userController.AddUser(enforcer))
		apiRoutes.POST("/signin", userController.SignInUser)
	}

	userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
	{
		userProtectedRoutes.GET("/", middleware.Authorize("report", "read", enforcer), userController.GetAllUser)
		userProtectedRoutes.GET("/:user", middleware.Authorize("report", "read", enforcer), userController.GetUser)
		userProtectedRoutes.PUT("/:user", middleware.Authorize("report", "write", enforcer), userController.UpdateUser)
		userProtectedRoutes.DELETE("/:user", middleware.Authorize("report", "write", enforcer), userController.DeleteUser)
	}

	r.Run()

}
