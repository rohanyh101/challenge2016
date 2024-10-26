package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/roh4nyh/qube_challenge_2016/controllers"
)

func DistributorRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/distributors", controllers.GetDistributors)
	incomingRoutes.GET("/distributors/:distributor_id", controllers.GetDistributor)
	incomingRoutes.GET("/distributors/:distributor_id/check", controllers.CheckDistributorPermission)

	incomingRoutes.POST("/distributors/add", controllers.AddDistributor)
	incomingRoutes.PUT("/distributors/:distributor_id", controllers.UpdateDistributor)
}
