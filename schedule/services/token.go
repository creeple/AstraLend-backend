/**
 * @Author: lixiang
 * @Date: 2025-03-22 14:51
 * @Description: token.go
 */

package services

import (
	"AstraLend-backend/api/models"
	"AstraLend-backend/db"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

func GetTokenInfo(token string, chainId string) (error, models.TokenInfo) {
	tokenInfo := models.TokenInfo{}
	redisKey := "token:" + chainId + ":" + token
	redisTokenInfo, _ := db.RedisGet(redisKey)
	// If the token information is not in the cache, query the database
	if len(redisTokenInfo) <= 0 {
		err := db.Mysql.Table("token_info").Where("token = ? and chainId = ?", token, chainId).First(&tokenInfo).Debug().Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, tokenInfo
			} else {
				return errors.New("record select err " + err.Error()), tokenInfo
			}
		}
		_ = db.RedisSet(redisKey, models.RedisTokenInfo{
			Token:   token,
			ChainId: chainId,
			Price:   tokenInfo.Price,
			Logo:    tokenInfo.Logo,
			Symbol:  tokenInfo.Symbol,
		}, 0)
		return nil, tokenInfo
	} else {
		redisToken := models.RedisTokenInfo{}
		_ = json.Unmarshal(redisTokenInfo, &redisToken)
		return nil, models.TokenInfo{
			Logo:    redisToken.Logo,
			Token:   token,
			Symbol:  redisToken.Symbol,
			ChainId: chainId,
			Price:   redisToken.Price,
		}
	}
}
