package utils

import (
	"reflect"
	"unsafe"
)

func UpdateField[T any](value any, fieldName string, data T) {
	val := reflect.Indirect(reflect.ValueOf(value))
	member := val.FieldByName(fieldName)
	pointer := unsafe.Pointer(member.UnsafeAddr())
	realPointer := (*T)(pointer)
	*realPointer = data
}

func GetField[T any](value any) T {
	val := reflect.Indirect(reflect.ValueOf(value))
	member := val.FieldByName("claims")
	pointer := unsafe.Pointer(member.UnsafeAddr())
	realPointer := (*T)(pointer)
	if realPointer == nil {
		var v T
		return v
	}
	return *realPointer
}
