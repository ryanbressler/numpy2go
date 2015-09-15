package main

import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

//export Test
func Test(data *C.double, length C.int) {

	// See on:
	// https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices

	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(data)),
		Len:  int(length),
		Cap:  int(length),
	}
	slice := *(*[]float64)(unsafe.Pointer(&header))
	fmt.Printf("Go says %v\n", slice)

}

func main() {}
