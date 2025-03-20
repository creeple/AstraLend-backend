/**
 * @Author: lixiang
 * @Date: 2025-03-20 10:43
 * @Description: poolDataInfoService.go
 */

package services

import (
	"AstraLend-backend/api/constant"
	"AstraLend-backend/api/models"
	"AstraLend-backend/db"
)

func GetPoolDataByChainId(chainId int, data *[]models.PoolData) int {
	err := db.Mysql.Table("pooldata").Where("chain_id = ?", chainId).Order("pool_id asc").Find(data).Debug().Error
	if err != nil {
		return constant.CommonErrServerErr
	}
	return constant.CommonSuccess
}
