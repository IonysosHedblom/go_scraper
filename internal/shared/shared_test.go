package shared

import "testing"

func TestConvertStringToInt64(t *testing.T) {
	var num int64 = 5
	var str string = "5"
	expected := &num
	actual, _ := ConvertStringToInt64(str)

	if *actual != *expected {
		t.Errorf("ConvertstringtoInt(%s) actual %v, wanted %v", str, actual, expected)
	}
}

func TestConvertStringToInt64Error(t *testing.T) {
	var str string = "asd"
	_, err := ConvertStringToInt64(str)

	if err == nil {
		t.Errorf("ConvertstringtoInt(%s) expected error got %v", str, err)
	}
}
