package validation

import (
	"fmt"

	"github.com/xinliangnote/go-gin-api/configs"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	zhTranslation "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func init() {
	lang := configs.Get().Language.Local

	if lang == configs.ZhCN {
		trans, _ = ut.New(zh.New()).GetTranslator("zh")
		if err := zhTranslation.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans); err != nil {
			fmt.Println("validator zh translation error", err)
		}
	}

	if lang == configs.EnUS {
		trans, _ = ut.New(en.New()).GetTranslator("en")
		if err := enTranslation.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans); err != nil {
			fmt.Println("validator en translation error", err)
		}
	}
}

func Error(err error) (message string) {
	if validationErrors, ok := err.(validator.ValidationErrors); !ok {
		return err.Error()
	} else {
		for _, e := range validationErrors {
			message += e.Translate(trans) + ";"
		}
	}
	return message
}
