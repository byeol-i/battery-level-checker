package device

import (
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/byeol-i/battery-level-checker/pkg/models"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

func SpecValidator(spec *DeviceSpec) error {
	validate := validator.New()
	validate.RegisterValidation("script", models.ValidateScript)

	err := validate.Struct(spec)
	if err != nil {
		logger.Error("Device's Spec is not valid", zap.Any("spec", spec))
		return err
	}

	return nil
}


func BatteryLevelValidator(level *BatteryLevel) error {
	validate := validator.New()
	validate.RegisterValidation("script", models.ValidateScript)

	err := validate.Struct(level)
	if err != nil {
		logger.Error("Device's BatteryLevel is not valid", zap.Any("BatteryLevel", level))
		return err
	}

	return nil
}
