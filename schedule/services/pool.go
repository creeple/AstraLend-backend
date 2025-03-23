/**
 * @Author: lixiang
 * @Date: 2025-03-22 14:12
 * @Description: pool.go
 */

package services

import (
	"AstraLend-backend/api/models"
	"AstraLend-backend/config"
	"AstraLend-backend/db"
	"AstraLend-backend/log"
	astralend "AstraLend-backend/schedule/contracts"
	"AstraLend-backend/utils"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

func UpdatePoolInfo() {
	testNetConfig := config.Config.TestNet
	chainId := testNetConfig.ChainId
	//连接节点
	conn, err := ethclient.Dial(testNetConfig.NetUrl)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	astraLendPool, err := astralend.NewAstralend(common.HexToAddress(testNetConfig.AstraAddress), conn)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	borrowFee, err := astraLendPool.AstralendCaller.BorrowFee(nil)
	lendFee, err := astraLendPool.AstralendCaller.LendFee(nil)
	poolLength, err := astraLendPool.AstralendCaller.PoolLength(nil)

	for i := 0; i < int(poolLength.Int64()); i++ {
		poolId := i + 1
		poolBaseInfo, err := astraLendPool.AstralendCaller.PoolBaseInfos(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Logger.Error(err.Error())
			continue
		}
		_, borrowTokenInfo := GetTokenInfo(poolBaseInfo.BorrowToken.String(), chainId)
		_, lendTokenInfo := GetTokenInfo(poolBaseInfo.LendToken.String(), chainId)
		lendTokenJson, _ := json.Marshal(models.LendToken{
			LendFee:    lendFee.String(),
			TokenLogo:  lendTokenInfo.Logo,
			TokenName:  lendTokenInfo.Symbol,
			TokenPrice: lendTokenInfo.Price,
		})
		borrowTokenJson, _ := json.Marshal(models.BorrowToken{
			BorrowFee:  borrowFee.String(),
			TokenLogo:  borrowTokenInfo.Logo,
			TokenName:  borrowTokenInfo.Symbol,
			TokenPrice: borrowTokenInfo.Price,
		})

		poolBase := models.PoolBases{

			SettleTime:             poolBaseInfo.SettleTime.String(),
			PoolId:                 poolId,
			ChainId:                chainId,
			EndTime:                poolBaseInfo.EndTime.String(),
			InterestRate:           poolBaseInfo.InterestRate.String(),
			MaxSupply:              poolBaseInfo.MaxSupply.String(),
			LendSupply:             poolBaseInfo.LendSupply.String(),
			BorrowSupply:           poolBaseInfo.BorrowSupply.String(),
			MartgageRate:           poolBaseInfo.MortgageRate.String(),
			LendToken:              poolBaseInfo.LendToken.String(),
			LendTokenSymbol:        lendTokenInfo.Symbol,
			LendTokenInfo:          string(lendTokenJson),
			BorrowToken:            poolBaseInfo.BorrowToken.String(),
			BorrowTokenSymbol:      borrowTokenInfo.Symbol,
			BorrowTokenInfo:        string(borrowTokenJson),
			State:                  utils.IntToString(int(poolBaseInfo.State)),
			SpCoin:                 poolBaseInfo.SpCoin.String(),
			JpCoin:                 poolBaseInfo.JpCoin.String(),
			AutoLiquidateThreshold: poolBaseInfo.AutoLiquidateThreshold.String(),
		}
		key := "base_info:pool_" + chainId + "_" + utils.IntToString(poolId)
		hasValue, value, baseInfoMd5Str := getPoolMd5(&poolBase, key)
		if !hasValue || value != baseInfoMd5Str {

		}
	}

}
func getPoolMd5(baseInfo *models.PoolBases, key string) (bool, string, string) {
	baseInfoJson, _ := json.Marshal(baseInfo)
	baseInfoMd5 := utils.Md5(string(baseInfoJson))
	value, _ := db.RedisGet(key)
	if len(value) > 0 {
		return true, strings.Trim(string(value), ""), baseInfoMd5
	} else {
		return false, "", baseInfoMd5
	}
}
