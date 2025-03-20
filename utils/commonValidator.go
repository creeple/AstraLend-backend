/**
 * @Author: lixiang
 * @Date: 2025-03-20 14:48
 * @Description: commonValidator.go
 */

package utils

import (
	"AstraLend-backend/api/constant"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
)

func CommonValidator(ctx *gin.Context, req interface{}) (int, map[string]string) {
	err := ctx.ShouldBind(req)
	if err == io.EOF {
		return constant.ParameterEmptyErr, nil
	} else if err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		errorsMap := make(map[string]string)
		for _, e := range errs {
			errorsMap[e.Field()] = e.Translate(Translator)
		}
		return constant.ParameterEmptyErr, errorsMap
	}
	return constant.CommonSuccess, nil
}
