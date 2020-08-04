package utilsbit

import (
	"reflect"
	"rfgocore/utilsstring"
	"testing"
)

func TestIntToString(t *testing.T) {
	var data string = utilsstring.IntToString(8080)
	var desireResult string = "8080"

	if !reflect.DeepEqual(data, desireResult) {
		t.Errorf("data %s == %s", data, desireResult)
	}

}
