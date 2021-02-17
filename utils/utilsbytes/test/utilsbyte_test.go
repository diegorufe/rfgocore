package utilsbytes

import (
	"reflect"
	utilsbytes "rfgocore/utils/utilsbytes"
	"testing"
)

func TestShiftLeft(t *testing.T) {
	var data uint64 = utilsbytes.ShiftBits(utilsbit.ShiftLeft, 80, 2)
	var desireResult uint64 = 320

	if !reflect.DeepEqual(data, desireResult) {
		t.Errorf("shiftBits %d == %d", data, desireResult)
	}

}

func TestShiftRight(t *testing.T) {
	var data uint64 = utilsbytes.ShiftBits(utilsbit.ShiftRight, 80, 2)
	var desireResult uint64 = 20

	if !reflect.DeepEqual(data, desireResult) {
		t.Errorf("shiftBits %d == %d", data, desireResult)
	}

}

func TestUint64ToHex(t *testing.T) {
	var data uint64 = 11525
	var hex string = utilsbytes.Uint64ToHex(data)

	if !reflect.DeepEqual(hex, "2d05") {
		t.Errorf("uintToHex %s == %s", hex, "2d05")
	}

}

func TestHexToUint64(t *testing.T) {
	var data string = "2d05"
	value, _ := utilsbytes.HexToUint64(data)
	var desire uint64 = 11525

	if !reflect.DeepEqual(value, desire) {
		t.Errorf("uintToHex %d == %d", value, desire)
	}

}
