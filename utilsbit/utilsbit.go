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

// convertToShiftType is a function to convert bool to ShiftType operator
func convertToShiftType(shiftLeftValue bool) ShiftType {
	if shiftLeftValue {
		return ShiftLeft
	}
	return ShiftRight
}

// shiftBits function to shift bits for value in indicate shift value positions.
func shiftBits(shiftType ShiftType, data uint64, shiftvalue uint64) uint64 {
	var result uint64 = data

	if shiftType == ShiftLeft {
		result = data << shiftvalue
	} else {
		result = data >> shiftvalue
	}

	return result
}

// uint64ToHex function to convert uint64 to hex string value
func uint64ToHex(value uint64) string {
	return fmt.Sprintf("%0x", value)
}

// function to convert uint 64 value to hex
func hexToUint64(value string) (uint64, error) {
	return strconv.ParseUint(value, 16, 64)
}
