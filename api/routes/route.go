/**
 * @Author: lixiang
 * @Date: 2025-03-19 20:59
 * @Description: route.go
 */

package routes

import (
	"AstraLend-backend/api/controller"
	"github.com/gin-gonic/gin"
)

func InitRoute(e *gin.Engine) {
	group := e.Group("/api")
	user := controller.User{}
	group.POST("/login", user.Login)
}
