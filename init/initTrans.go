package init

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func InitTrans(locaal string) error {
	// 修改gin框架中得validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New()
		enT := en.New()
		// 第一个参数是备用得语言环境，后面是应该支持得语言环境
		un := ut.New(enT, zhT, enT)
		trans, ok := un.GetTranslator(locaal)
		if !ok {
			return errors.New("un.GetTranslator出错了")
		}
		switch locaal {
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			en_translations.RegisterDefaultTranslations(v, trans)
		}
		return nil
	}
	return errors.New("binding.Validator.Engine().(*validator.Validate) 出错了")
}
