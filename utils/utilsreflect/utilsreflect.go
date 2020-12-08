package utilsreflect

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// SetField : method to set field value
func SetField(instance interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(instance).Elem()
	structFieldValue := structValue.FieldByName(strings.Title(name))

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)

	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

// FillMapIntoStruct : method for fill map into instance
func FillMapIntoStruct(instance interface{}, mapValues map[string]interface{}) error {
	for k, v := range mapValues {
		err := SetField(instance, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
