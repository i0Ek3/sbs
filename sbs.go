package sbs

import (
	"reflect"
	"unsafe"
)

// S2B converts string to []byte
func S2B(s string) []byte {
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	byt := reflect.SliceHeader{
		Data: str.Data,
		Len:  str.Len,
		Cap:  str.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&byt))
}

// B2S converts []byte to string
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
