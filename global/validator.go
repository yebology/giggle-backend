package global

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/constant/post"
)

var (

	// Validate holds a single instance of the validator.Validate struct
	Validate *validator.Validate

	// once ensures the InitValidator function is executed only once, even in a concurrent environment
	once sync.Once

)

// InitValidator initializes the validator instance and registers custom validation functions.
// This function uses the `sync.Once` mechanism to ensure it only runs once.
func InitValidator() {
	once.Do(func() {
		Validate = validator.New() // Create a new validator instance
		RegisterValidation()       // Register custom validation rules
	})
}

// RegisterValidation registers all custom validation functions with the validator instance.
func RegisterValidation() {
	Validate.RegisterValidation("validRole", ValidateRole)
	Validate.RegisterValidation("validPostType", ValidatePostType)
	Validate.RegisterValidation("validPostCategory", ValidatePostCategory)
	Validate.RegisterValidation("validPostStatus", ValidatePostStatus)
	Validate.RegisterValidation("isFalse", ValidateProposalStatus)
}

// GetValidator returns the singleton validator instance.
// It ensures the validator is initialized before returning.
func GetValidator() *validator.Validate {
	InitValidator() // Ensure validator is initialized
	return Validate
}

// ValidateRole checks if the input string is a valid role as defined in the `constant.AllowedRole` slice.
func ValidateRole(fl validator.FieldLevel) bool {
	input := fl.Field().String()
	for _, allowed := range constant.AllowedRole {
		if string(allowed) == input {
			return true
		}
	}
	return false
}

// ValidatePostType checks if the input string is a valid post type as defined in the `post.AllowedType` slice.
func ValidatePostType(fl validator.FieldLevel) bool {
	input := fl.Field().String()
	for _, allowed := range post.AllowedType {
		if string(allowed) == input {
			return true
		}
	}
	return false
}

// ValidatePostCategory checks if the input string is a valid post category as defined in the `post.AllowedCategories` slice.
func ValidatePostCategory(fl validator.FieldLevel) bool {
	input := fl.Field().String()
	for _, allowed := range post.AllowedCategories {
		if string(allowed) == input {
			return true
		}
	}
	return false
}

// ValidatePostStatus checks if the input string is a valid post status as defined in the `post.AllowedStatus` slice.
func ValidatePostStatus(fl validator.FieldLevel) bool {
	input := fl.Field().String()
	for _, allowed := range post.AllowedStatus {
		if string(allowed) == input {
			return true
		}
	}
	return false
}

func ValidateProposalStatus(fl validator.FieldLevel) bool {
	return !fl.Field().Bool()
}
