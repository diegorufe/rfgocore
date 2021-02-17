package utilsstring

import (
	"fmt"
	"strings"
)

// Empty constant
const Empty string = ""

// Space constant
const Space string = ""

// BreakLine constant
const BreakLine string = "\n"

// Tabulation constant
const Tabulation string = "\t"

// TrimValue constants value to use trim
const TrimValue string = Tabulation + " " + BreakLine

// IntToString Method for conver int value to string
func IntToString(value int) string {
	return fmt.Sprint(value)
}

// IsEmpty method to check string is empty
func IsEmpty(value string) bool {
	var result string = Trim(value)
	return result == Empty
}

// IsNotEmpty method to check string is not empty
func IsNotEmpty(value string) bool {
	return !IsEmpty(value)
}

// TrimAllSpace method for trim space
func TrimAllSpace(value string) string {
	return strings.TrimSpace(value)
}

// Trim method for trim left and right by \t \n values
func Trim(value string) string {
	return strings.TrimLeft(strings.TrimRight(value, TrimValue), TrimValue)
}
