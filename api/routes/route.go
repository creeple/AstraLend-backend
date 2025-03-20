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
	groupUser := e.Group("/api/user")
	user := controller.User{}
	groupUser.POST("/login", user.Login)
	groupPool := e.Group("/api/pool")
	pool := controller.Astra{}
	groupPool.POST("/poolBaseInfo", pool.GetPoolBaseInfo)
}
