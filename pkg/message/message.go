package message

import "github.com/go-playground/validator/v10"

func MsgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is below the minimum allowed"
	case "max":
		return "Value exceeds the maximum allowed"
	default:
		return "Invalid value"
	}
}
