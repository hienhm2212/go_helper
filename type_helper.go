package helper

import (
	"bytes"
	"reflect"
)

var numericZeros = []interface{}{
	int(0),
	int8(0),
	int16(0),
	int32(0),
	int64(0),
	uint(0),
	uint8(0),
	uint16(0),
	uint32(0),
	uint64(0),
	float32(0),
	float64(0),
}

// numToFloat64 converts any numeric value to float64
func numToFloat64(v interface{}) (float64, bool) {
	var vf float64
	isOk := true

	switch vn := v.(type) {
	case uint8:
		vf = float64(vn)
	case uint16:
		vf = float64(vn)
	case uint32:
		vf = float64(vn)
	case uint64:
		vf = float64(vn)
	case int:
		vf = float64(vn)
	case int8:
		vf = float64(vn)
	case int16:
		vf = float64(vn)
	case int32:
		vf = float64(vn)
	case int64:
		vf = float64(vn)
	case float32:
		vf = float64(vn)
	case float64:
		vf = float64(vn)
	default:
		isOk = false
	}

	return vf, isOk
}

func isEqual(expected interface{}, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	if exp, ok := expected.([]byte); ok {
		act, ok := actual.([]byte)
		if !ok {
			return false
		}

		if exp == nil || act == nil {
			return true
		}

		return bytes.Equal(exp, act)
	}

	return reflect.DeepEqual(expected, actual)

}
