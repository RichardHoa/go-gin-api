package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {

	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSONResponse(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

func WriteErrorResponse(w http.ResponseWriter, status int, err error) {
	WriteJSONResponse(w, status, map[string]string{
		"error": err.Error(),
	})
}

func DebuggingPrinting(s interface{}) {
	t := reflect.TypeOf(s)
	val := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := val.Field(i)

		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
}

func CreateFriendlyErrorMSG(err error) map[string]string {
	friendlyErrors := make(map[string]string)
	if errors, ok := err.(validator.ValidationErrors); ok {
		// Map validation errors into a more friendly format
		for _, validationError := range errors {
			field := validationError.Field()
			tag := validationError.Tag()
			// Create user-friendly messages based on the validation tag
			switch tag {
			case "required":
				friendlyErrors[field] = fmt.Sprintf("%s is required.", field)
			case "max":
				friendlyErrors[field] = fmt.Sprintf("%s cannot be longer than %s characters.", field, validationError.Param())
			case "min":
				friendlyErrors[field] = fmt.Sprintf("%s must be at least %s characters.", field, validationError.Param())
			case "email":
				friendlyErrors[field] = fmt.Sprintf("%s must be a valid email address.", field)
			default:
				friendlyErrors[field] = fmt.Sprintf("%s is invalid.", field)
			}
		}
	}

	return friendlyErrors

}
