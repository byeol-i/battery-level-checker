package device

import (
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

func SpecValidator(spec *Spec) error {
	var validate *validator.Validate

	err := validate.Struct(spec)
	if err != nil {
		logger.Error("Device's Spec is not valid", zap.Any("spec", spec))
		return err
	}

	return nil
}


func BatteryLevelValidator(level *BatteryLevel) error {
	var validate *validator.Validate

	err := validate.Struct(level)
	if err != nil {
		logger.Error("Device's BatteryLevel is not valid", zap.Any("BatteryLevel", level))
		return err
	}

	return nil
}
