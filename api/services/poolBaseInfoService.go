/**
 * @Author: lixiang
 * @Date: 2025-03-20 10:43
 * @Description: poolBaseInfoService.go
 */

package services

import (
	"AstraLend-backend/api/constant"
	"AstraLend-backend/api/models"
	"AstraLend-backend/db"
	"encoding/json"
	"fmt"
)

func GetPoolBaseInfoByChainId(chainId int, poolBaseInfos *[]models.PoolBaseInfo) (int, error) {
	var poolBases []models.PoolBases
	fmt.Print(db.Mysql)
	err := db.Mysql.Table("poolbases").Where("chain_id = ?", chainId).Order("pool_id asc").Find(&poolBases).Debug().Error

	if err != nil {
		return constant.CommonErrServerErr, err
	}
	for _, poolBase := range poolBases {
		borrowTokenInfo := models.BorrowTokenInfo{}
		_ = json.Unmarshal([]byte(poolBase.BorrowTokenInfo), &borrowTokenInfo)
		lendTokenInfo := models.LendTokenInfo{}
		_ = json.Unmarshal([]byte(poolBase.LendTokenInfo), &lendTokenInfo)
		*poolBaseInfos = append(*poolBaseInfos, models.PoolBaseInfo{
			PoolID:                 poolBase.PoolID,
			AutoLiquidateThreshold: poolBase.AutoLiquidateThreshold,
			BorrowSupply:           poolBase.BorrowSupply,
			BorrowToken:            poolBase.BorrowToken,
			BorrowTokenInfo:        borrowTokenInfo,
			EndTime:                poolBase.EndTime,
			InterestRate:           poolBase.InterestRate,
			JpCoin:                 poolBase.JpCoin,
			LendSupply:             poolBase.LendSupply,
			LendToken:              poolBase.LendToken,
			LendTokenInfo:          lendTokenInfo,
			MartgageRate:           poolBase.MartgageRate,
			MaxSupply:              poolBase.MaxSupply,
			SettleTime:             poolBase.SettleTime,
			SpCoin:                 poolBase.SpCoin,
			State:                  poolBase.State,
		})
	}
	return constant.CommonSuccess, nil
}
func SavePoolBaseInfo(poolBase *models.PoolBases) {
}
