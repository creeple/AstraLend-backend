/**
 * @Author: lixiang
 * @Date: 2025-03-18 18:09
 * @Description: main.go.go
 */

package main

import (
	"AstraLend-backend/api/routes"
	"AstraLend-backend/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	address := config.Config.Mysql.Address
	fmt.Println(address)
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	routes.InitRoute(app)
	_ = app.Run(":8080")
}
