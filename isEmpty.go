package isEmpty

import (
	"reflect"
)

// Values same as Value but for multiple values
func Values(values ...interface{}) bool {
	for _, value := range values {
		if Value(value) {
			return true
		}

	}
	return false
}

// Value will return true if variable has zero value of its type.
// Also it will return true slice length is 0.
func Value(value interface{}) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Array, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Map, reflect.Slice:
		if v.IsNil() {
			return true
		}
		return v.Len() == 0
	default:
		return value == nil
	}
}
