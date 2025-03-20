/**
 * @Author: lixiang
 * @Date: 2025-03-20 9:18
 * @Description: astra.go
 */

package controller

import (
	"AstraLend-backend/api/constant"
	"AstraLend-backend/api/models"
	"AstraLend-backend/api/models/request"
	"AstraLend-backend/api/models/response"
	"AstraLend-backend/api/services"
	"AstraLend-backend/utils"
	"github.com/gin-gonic/gin"
)

type Astra struct {
}

func (astra *Astra) GetPoolBaseInfo(ctx *gin.Context) {
	req := request.PoolRequest{}
	var result []models.PoolBaseInfo
	//校验参数
	code, data := utils.CommonValidator(ctx, &req)
	if code != constant.CommonSuccess {
		response.Response(ctx, code, data)
		return
	}
	//获取数据
	chainId := req.ChainId
	code, err := services.GetPoolBaseInfoByChainId(chainId, &result)
	if err != nil {
		response.Response(ctx, code, err)
	}
	response.Response(ctx, constant.CommonSuccess, result)
}

func (astra *Astra) GetPoolBaseList(ctx *gin.Context) {
	req := request.PoolRequest{}
	var result []models.PoolData
	//参数校验
	code, data := utils.CommonValidator(ctx, &req)
	if code != constant.CommonSuccess {
		response.Response(ctx, code, data)
		return
	}
	chainId := req.ChainId
	code = services.GetPoolDataByChainId(chainId, &result)
	if code != constant.CommonSuccess {
		response.Response(ctx, code, nil)
		return
	}
	response.Response(ctx, constant.CommonSuccess, result)
}
