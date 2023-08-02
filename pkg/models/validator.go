package models

import (
	"regexp"

	"github.com/go-playground/validator"
)


func ValidateScript(fl validator.FieldLevel) bool {
	script := fl.Field().String()

   return !regexp.MustCompile(`(?i)<script[\s\S]*?>[\s\S]*?</script[\s\S]*?>`).MatchString(script)
}