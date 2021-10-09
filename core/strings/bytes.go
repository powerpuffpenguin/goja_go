package strings

import "unsafe"

func StringToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
