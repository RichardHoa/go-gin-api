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

func WriteJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

func WriteErrorResponse(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{
		"error": err.Error(),
	})
}


func PrintStructFields(s interface{}) {
	t := reflect.TypeOf(s)
	val := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := val.Field(i)

		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
}