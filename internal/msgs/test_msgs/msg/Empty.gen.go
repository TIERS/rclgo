/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package test_msgs_msg
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltest_msgs__rosidl_typesupport_c -ltest_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/msg/empty.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/Empty", EmptyTypeSupport)
}

// Do not create instances of this type directly. Always use NewEmpty
// function instead.
type Empty struct {
}

// NewEmpty creates a new Empty with default values.
func NewEmpty() *Empty {
	self := Empty{}
	self.SetDefaults()
	return &self
}

func (t *Empty) Clone() *Empty {
	c := &Empty{}
	return c
}

func (t *Empty) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Empty) SetDefaults() {
}

// CloneEmptySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneEmptySlice(dst, src []Empty) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var EmptyTypeSupport types.MessageTypeSupport = _EmptyTypeSupport{}

type _EmptyTypeSupport struct{}

func (t _EmptyTypeSupport) New() types.Message {
	return NewEmpty()
}

func (t _EmptyTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__msg__Empty
	return (unsafe.Pointer)(C.test_msgs__msg__Empty__create())
}

func (t _EmptyTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__msg__Empty__destroy((*C.test_msgs__msg__Empty)(pointer_to_free))
}

func (t _EmptyTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _EmptyTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _EmptyTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__msg__Empty())
}

type CEmpty = C.test_msgs__msg__Empty
type CEmpty__Sequence = C.test_msgs__msg__Empty__Sequence

func Empty__Sequence_to_Go(goSlice *[]Empty, cSlice CEmpty__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Empty, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.test_msgs__msg__Empty__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__Empty * uintptr(i)),
		))
		EmptyTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Empty__Sequence_to_C(cSlice *CEmpty__Sequence, goSlice []Empty) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__msg__Empty)(C.malloc((C.size_t)(C.sizeof_struct_test_msgs__msg__Empty * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.test_msgs__msg__Empty)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_test_msgs__msg__Empty * uintptr(i)),
		))
		EmptyTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Empty__Array_to_Go(goSlice []Empty, cSlice []CEmpty) {
	for i := 0; i < len(cSlice); i++ {
		EmptyTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Empty__Array_to_C(cSlice []CEmpty, goSlice []Empty) {
	for i := 0; i < len(goSlice); i++ {
		EmptyTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
