package extend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

var Translator ut.Translator

func RegisterTranslations(r *gin.Engine) error {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		zh := zh.New()
		uni := ut.New(en, en, zh)
		Translator, ok = uni.GetTranslator("zh")
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", "zh")
		}
		zh2.RegisterDefaultTranslations(v, Translator)

	}
	return nil
}
func getErrorInfo(errs validator.ValidationErrors) string {
	errInfos := make([]string, 0, len(errs))
	for _, e := range errs {
		errInfo := e.Translate(Translator)
		errInfos = append(errInfos, fmt.Sprintf("%v:%v", e.Field(), errInfo))
	}
	return strings.Join(errInfos, ",")
}
