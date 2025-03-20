/**
 * @Author: lixiang
 * @Date: 2025-03-20 15:04
 * @Description: translator.go
 */

package utils

import (
	"AstraLend-backend/log"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zt "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
)

var Translator ut.Translator

func InitTranslator() {
	validate := validator.New()
	zhT := zh.New()
	uni := ut.New(zhT, zhT)
	Translator, _ := uni.GetTranslator("zh")
	if err := zt.RegisterDefaultTranslations(validate, Translator); err != nil {
		log.Logger.Error("register default translation error", zap.Error(err))
	}
}
