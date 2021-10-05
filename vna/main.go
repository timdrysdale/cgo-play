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

	result := C.pocketvna_force_unlock_devices()
	fmt.Println(result)
	handle := C.PVNA_DeviceHandler(nil)
	result = C.pocketvna_get_first_device_handle(&handle)
	fmt.Println(result)
	result = C.pocketvna_release_handle(&handle)
	fmt.Println(result)
}
