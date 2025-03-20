/**
 * @Author: lixiang
 * @Date: 2025-03-19 11:02
 * @Description: login.go
 */

package request

import (
	"AstraLend-backend/api/constant"
	"AstraLend-backend/api/models/response"
	"AstraLend-backend/log"
	"AstraLend-backend/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
)

type LoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *LoginRequest) ValidateLogin(c *gin.Context, req *LoginRequest) int {
	err := c.ShouldBind(req)
	//判断请求体是否为空
	if err == io.EOF {
		return constant.ParameterEmptyErr
	}
	if err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		for _, e := range errs {
			if e.Field() == "Name" && e.Tag() == "required" {
				return constant.PNameEmpty
			}
			if e.Field() == "Password" && e.Tag() == "required" {
				return constant.PNameEmpty
			}
			return constant.CommonErrServerErr
		}
	}
	return constant.CommonSuccess
}

func (r *LoginRequest) Login(req *LoginRequest, token *response.Token) int {
	log.Logger.Sugar().Info("start login : ", req)
	if req.Name == "admin" && req.Password == "password" {
		tokenString, err := utils.CreateToken(req.Name)
		if err != nil {
			log.Logger.Sugar().Error("create token failed", err)
			return constant.CommonErrServerErr
		}
		token.TokenId = tokenString
		//存储到redis
		//_ = db.RedisSet(req.Name, "login_success", config.Config.Jwt.ExpireTime)
		return constant.CommonSuccess
	} else {
		return constant.NameOrPasswordErr
	}
}
