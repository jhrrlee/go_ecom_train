package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jhrrlee/go_ecom_train/controllers"
)

func UserRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/user/signup", controllers.Signup())
	incomingRoutes.POST("/user/login", controllers.Login())
	incomingRoutes.POST("/user/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/user/productview", controllers.SearchProduct())
	incomingRoutes.GET("/user/search", controllers.SearchProductByQuery())
}
