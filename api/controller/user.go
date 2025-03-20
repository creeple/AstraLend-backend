/**
 * @Author: lixiang
 * @Date: 2025-03-19 10:53
 * @Description: user.go
 */

package controller

import (
	"AstraLend-backend/api/constant"
	"AstraLend-backend/api/models/request"
	"AstraLend-backend/api/models/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct{}

func (t *User) Login(ctx *gin.Context) {
	fmt.Println(ctx.Request.URL.String())
	req := request.LoginRequest{}
	res := response.Token{}
	//请求参数封装+参数校验
	validateCode := req.ValidateLogin(ctx, &req)
	if validateCode != constant.CommonSuccess {
		response.Response(ctx, validateCode, nil)
		return
	}
	//请求登录接口
	loginCode := req.Login(&req, &res)
	if loginCode != constant.CommonSuccess {
		response.Response(ctx, loginCode, nil)
		return
	}
	response.Response(ctx, constant.CommonSuccess, nil)

}
