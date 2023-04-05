package utils

import (
	"reflect"
)

func Contains(slice interface{}, val interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == val {
			return true
		}
	}

	return false
}