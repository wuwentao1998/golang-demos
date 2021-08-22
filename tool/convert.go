package tool

import (
	"reflect"
	"unsafe"
)

func String2Bytes(str string) []byte {
	l := len(str)
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*(*reflect.StringHeader)(unsafe.Pointer(&str))).Data,
		Len:  l,
		Cap:  l,
	}))
}

func Bytes2String(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
