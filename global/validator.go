package global

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/constant/post"
)

var (

	Validate *validator.Validate
	once     sync.Once

)

func InitValidator() {

	once.Do(func() {
		Validate = validator.New()
		RegisterValidation()
	})

}

func RegisterValidation() {

	Validate.RegisterValidation("validRole", ValidateRole)
	Validate.RegisterValidation("validPostType", ValidatePostType)
	Validate.RegisterValidation("validPostCategory", ValidatePostCategory)
	Validate.RegisterValidation("validPostStatus", ValidatePostStatus)

}

func GetValidator() *validator.Validate {
	
	InitValidator()

	return Validate

}

func ValidateRole(fl validator.FieldLevel) bool {

	input := fl.Field().String()
	for _, allowed := range constant.AllowedRole {
		if string(allowed) == input {
			return true
		}
	}

	return false

}

func ValidatePostType(fl validator.FieldLevel) bool {
	
	input := fl.Field().String()
	for _, allowed := range post.AllowedType {
		if string(allowed) == input {
			return true
		}
	}

	return false

}

func ValidatePostCategory(fl validator.FieldLevel) bool {
	
	input := fl.Field().String()
	for _, allowed := range post.AllowedCategories {
		if string(allowed) == input {
			return true
		}
	}

	return false

}

func ValidatePostStatus(fl validator.FieldLevel) bool {
	
	input := fl.Field().String()
	for _, allowed := range post.AllowedStatus {
		if string(allowed) == input {
			return true
		}
	}

	return false

}
