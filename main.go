/**
 * @Author: lixiang
 * @Date: 2025-03-18 18:09
 * @Description: main.go.go
 */

package main

import (
	"AstraLend-backend/config"
	"fmt"
)

func main() {
	address := config.Config.Mysql.Address
	fmt.Println(address)
}
