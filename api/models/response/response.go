/**
 * @Author: lixiang
 * @Date: 2025-03-19 14:52
 * @Description: response.go
 */

package response

import (
	"AstraLend-backend/api/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Page struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

// ResponsePage 分页响应格式
func ResponsePage(ctx *gin.Context, code int, totalCount int, data interface{}) {
	//默认简体中文
	lang := constant.LangZh
	langValue, exists := ctx.Get("lang")
	if exists {
		lang = langValue.(int)
	}
	resp := Page{
		Code:  code,
		Total: totalCount,
		Data:  data,
		Msg:   constant.GetMsg(code, lang),
	}
	ctx.JSON(http.StatusOK, resp)
}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func Response(ctx *gin.Context, code int, data interface{}, httpStatus ...int) {
	lang := constant.LangZh
	langValue, exists := ctx.Get("lang")
	if exists {
		lang = langValue.(int)
	}
	resp := response{
		Code: code,
		Msg:  constant.GetMsg(code, lang),
		Data: data,
	}
	HttpStatus := 200
	if len(httpStatus) > 0 {
		HttpStatus = httpStatus[0]
	}
	ctx.JSON(HttpStatus, resp)
}

type Token struct {
	TokenId string `json:"token_id"`
}
