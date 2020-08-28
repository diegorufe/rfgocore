package utilsstring

import (
	"fmt"
	"strings"
)

// TrimValue constants value to use trim
const TrimValue string = "\t \n"

// IntToString Method for conver int value to string
func IntToString(value int) string {
	return fmt.Sprint(value)
}

// IsEmpty method to check string is empty
func IsEmpty(value string) bool {
	var result string = strings.TrimLeft(value, TrimValue)
	result = strings.TrimRight(value, TrimValue)
	return result != ""
}

// IsNotEmpty method to check string is not empty
func IsNotEmpty(value string) bool {
	return !IsEmpty(value)
}
