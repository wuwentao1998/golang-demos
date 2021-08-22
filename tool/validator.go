package tool

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var (
	ErrVarTypeNotStruct = errors.New("variable type to be verified is not structure")

	defaultValidator = validator.New()
)

func ValidateStruct(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("%w: obj is nil", ErrVarTypeNotStruct)
	}

	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}

	if valueType != reflect.Struct {
		return fmt.Errorf("%w: %v", ErrVarTypeNotStruct, valueType)

	}

	return defaultValidator.Struct(obj)
}
