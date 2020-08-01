package utilsbit

import (
	"fmt"
	"strconv"
)

// ShiftType is a type for operator bits
type ShiftType bool

const (
	// ShiftLeft indicate left shift operator
	ShiftLeft ShiftType = true
	// ShiftRight indicate right left shift operator
	ShiftRight ShiftType = false
)

// ConvertToShiftType is a function to convert bool to ShiftType operator
func ConvertToShiftType(shiftLeftValue bool) ShiftType {
	if shiftLeftValue {
		return ShiftLeft
	}
	return ShiftRight
}

// ShiftBits function to shift bits for value in indicate shift value positions.
func ShiftBits(shiftType ShiftType, data uint64, shiftvalue uint64) uint64 {
	var result uint64 = data

	if shiftType == ShiftLeft {
		result = data << shiftvalue
	} else {
		result = data >> shiftvalue
	}

	return result
}

// Uint64ToHex function to convert uint64 to hex string value
func Uint64ToHex(value uint64) string {
	return fmt.Sprintf("%0x", value)
}

// HexToUint64 function to convert uint 64 value to hex
func HexToUint64(value string) (uint64, error) {
	return strconv.ParseUint(value, 16, 64)
}
