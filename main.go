/**
 * @Author: lixiang
 * @Date: 2025-03-18 18:09
 * @Description: main.go.go
 */

package main

import (
	"AstraLend-backend/api/routes"
	"AstraLend-backend/config"
	"AstraLend-backend/db"
	"AstraLend-backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	address := config.Config.Mysql.Address
	fmt.Println(address)
	utils.InitTranslator()
	db.InitMysql()
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	routes.InitRoute(app)
	_ = app.Run(":8080")
}
