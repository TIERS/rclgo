/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package rosidl_runtime_c

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo CFLAGS: -I/opt/ros/foxy/include

#include "rosidl_runtime_c/string.h"

*/
import "C"
import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
)

/*
Due to the way the rcl string representation differs from Go representation, for serdes purposes treat the String as ros2types.ROS2Msg
so no special string-specific exceptions need to be made to the already complex ROS2 Msg serdes templating.
*/
type String string

func NewString() *String {
	self := String("")
	self.SetDefaults(nil)
	return &self
}

func (t *String) SetDefaults(d interface{}) ros2types.ROS2Msg {
	switch d.(type) {
	case string:
		*t = String(d.(string))
	case String:
		*t = d.(String)
	case nil:
		// *t is already ""
	default:
		panic(fmt.Sprintf("interface conversion: interface {} is %#v, not rosidl_runtime_c.String\n", d))
	}
	return t
}

func (t *String) TypeSupport() unsafe.Pointer {
	fmt.Printf("rosidl_runtime_c.TypeSupport() called. This is never meant to be directly addressed as a stand-alone data object in the ROS2 messaging bus.")
	return unsafe.Pointer(t)
}
func (t *String) PrepareMemory() unsafe.Pointer {
	fmt.Printf("rosidl_runtime_c.PrepareMemory() called. This is never meant to be directly addressed as a stand-alone data object in the ROS2 messaging bus.")
	return unsafe.Pointer(t)
}
func (t *String) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	fmt.Printf("rosidl_runtime_c.ReleaseMemory() called. This is never meant to be directly addressed as a stand-alone data object in the ROS2 messaging bus.")
}
func (t *String) AsCStruct() unsafe.Pointer { // rosidl_runtime_c__String__assignn() does something like this, but to call it we still need to make a C string and free it.
	mem := (*C.rosidl_runtime_c__String)(C.malloc(C.sizeof_struct_rosidl_runtime_c__String)) //TODO add this to template generator

	mem.data = (*C.char)(C.malloc((C.size_t)(C.sizeof_char * uintptr(len(*t)+1))))

	for i := 0; i < len(*t); i++ {
		t.setDataCArrayIndex(mem, i, (*t)[i])
	}
	t.setDataCArrayIndex(mem, len(*t), '\x00')
	mem.size = C.size_t(len(*t))
	mem.capacity = C.size_t(len(*t) + 1)
	return unsafe.Pointer(mem)
}
func (t *String) setDataCArrayIndex(mem *C.rosidl_runtime_c__String, i int, v byte) {
	cIdx := (*C.uint8_t)(unsafe.Pointer(
		uintptr(unsafe.Pointer(mem.data)) + (C.sizeof_uint8_t * uintptr(i)),
	))
	*cIdx = (C.uint8_t)(v)
}
func (t *String) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.rosidl_runtime_c__String)(ros2_message_buffer)
	sb := strings.Builder{}
	sb.Grow(int(mem.size))
	for i := 0; i < int(mem.size); i++ {
		cIdx := (*C.uint8_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(mem.data)) + (C.sizeof_uint8_t * uintptr(i)),
		))
		sb.WriteByte(byte(*cIdx))
	}
	*t = String(sb.String())
}
func (t *String) Clone() ros2types.ROS2Msg {
	c := *t
	return &c
}

type CString = C.rosidl_runtime_c__String
type CString__Sequence = C.rosidl_runtime_c__String__Sequence

func String__Sequence_to_Go(goSlice *[]String, cSlice CString__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]String, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rosidl_runtime_c__String)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rosidl_runtime_c__String * uintptr(i)),
		))
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func String__Sequence_to_C(cSlice *CString__Sequence, goSlice []String) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rosidl_runtime_c__String)(C.malloc((C.size_t)(C.sizeof_struct_rosidl_runtime_c__String * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rosidl_runtime_c__String)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rosidl_runtime_c__String * uintptr(i)),
		))
		*cIdx = *(*C.rosidl_runtime_c__String)(v.AsCStruct())
	}
}
func String__Array_to_Go(goSlice []String, cSlice []CString) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func String__Array_to_C(cSlice []CString, goSlice []String) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.rosidl_runtime_c__String)(goSlice[i].AsCStruct())
	}
}
