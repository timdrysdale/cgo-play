//
// main.go
// Copyright (C) 2019 Tim Hughes
//
// Distributed under terms of the MIT license.
//
package main

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L. -lPocketVnaApi_x64
#include "pocketvna.h"
*/
import "C"
import "fmt"

func main() {

	handle := C.PVNA_DeviceHandler(nil)
	result := C.pocketvna_get_first_device_handle(&handle)
	fmt.Println(result)

}
