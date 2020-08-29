package utilsbit

import (
	"reflect"
	"rfgocore/utils/utilsstring"
	"testing"
)

func TestIntToString(t *testing.T) {
	var data string = utilsstring.IntToString(8080)
	var desireResult string = "8080"

	if !reflect.DeepEqual(data, desireResult) {
		t.Errorf("data %s == %s", data, desireResult)
	}

}

func TestIsEmpty(t *testing.T) {
	var data string = ""

	if utilsstring.IsEmpty(data) {
		t.Errorf("data %s", data)
	}

	data = " eee"

	if !utilsstring.IsEmpty(data) {
		t.Errorf("data %s", data)
	}

}

func TestIsNotEmpty(t *testing.T) {
	var data string = " asd"

	if utilsstring.IsNotEmpty(data) {
		t.Errorf("data %s", data)
	}

	data = ""

	if !utilsstring.IsNotEmpty(data) {
		t.Errorf("data %s", data)
	}

}
