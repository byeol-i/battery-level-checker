package user

import (
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/byeol-i/battery-level-checker/pkg/models"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)


func UserValidator(spec *UserImpl) error {
	validate := validator.New()
	validate.RegisterValidation("script", models.ValidateScript)

	err := validate.Struct(spec)
	if err != nil {
		logger.Error("User's Spec is not valid", zap.Any("spec", spec))
		return err
	}

	return nil
}

func TokenValidator(spec *Token) error {
	validate := validator.New()
	validate.RegisterValidation("script", models.ValidateScript)

	err := validate.Struct(spec)
	if err != nil {
		logger.Error("User's Spec is not valid", zap.Any("spec", spec))
		return err
	}

	return nil
}

