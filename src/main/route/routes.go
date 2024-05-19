package route

import (
	"HunCoding/src/main/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	//-------- User --------//
	r.GET("/get-user-by-id/:userId", controller.FindUserById)
	r.GET("/get-user-by-email/:email", controller.FindUserByEmail)
	r.POST("/create-user", controller.CreateUser)
	r.PUT("/update-ser/:userId", controller.UpdateUser)
	r.DELETE("/delete-user/:userId", controller.DeleteUser)
}
