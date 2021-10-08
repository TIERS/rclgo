/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces_msg
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/TIERS/rclgo/internal/msgs/std_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/multi_array_layout.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/MultiArrayLayout", MultiArrayLayoutTypeSupport)
}

// Do not create instances of this type directly. Always use NewMultiArrayLayout
// function instead.
type MultiArrayLayout struct {
	Dim []std_msgs_msg.MultiArrayDimension `yaml:"dim"`// Array of dimension properties
	DataOffset uint32 `yaml:"data_offset"`// padding bytes at front of data
}

// NewMultiArrayLayout creates a new MultiArrayLayout with default values.
func NewMultiArrayLayout() *MultiArrayLayout {
	self := MultiArrayLayout{}
	self.SetDefaults()
	return &self
}

func (t *MultiArrayLayout) Clone() *MultiArrayLayout {
	c := &MultiArrayLayout{}
	if t.Dim != nil {
		c.Dim = make([]std_msgs_msg.MultiArrayDimension, len(t.Dim))
		std_msgs_msg.CloneMultiArrayDimensionSlice(c.Dim, t.Dim)
	}
	c.DataOffset = t.DataOffset
	return c
}

func (t *MultiArrayLayout) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MultiArrayLayout) SetDefaults() {
	t.Dim = nil
	t.DataOffset = 0
}

// CloneMultiArrayLayoutSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMultiArrayLayoutSlice(dst, src []MultiArrayLayout) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MultiArrayLayoutTypeSupport types.MessageTypeSupport = _MultiArrayLayoutTypeSupport{}

type _MultiArrayLayoutTypeSupport struct{}

func (t _MultiArrayLayoutTypeSupport) New() types.Message {
	return NewMultiArrayLayout()
}

func (t _MultiArrayLayoutTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__MultiArrayLayout
	return (unsafe.Pointer)(C.example_interfaces__msg__MultiArrayLayout__create())
}

func (t _MultiArrayLayoutTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__MultiArrayLayout__destroy((*C.example_interfaces__msg__MultiArrayLayout)(pointer_to_free))
}

func (t _MultiArrayLayoutTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MultiArrayLayout)
	mem := (*C.example_interfaces__msg__MultiArrayLayout)(dst)
	std_msgs_msg.MultiArrayDimension__Sequence_to_C((*std_msgs_msg.CMultiArrayDimension__Sequence)(unsafe.Pointer(&mem.dim)), m.Dim)
	mem.data_offset = C.uint32_t(m.DataOffset)
}

func (t _MultiArrayLayoutTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MultiArrayLayout)
	mem := (*C.example_interfaces__msg__MultiArrayLayout)(ros2_message_buffer)
	std_msgs_msg.MultiArrayDimension__Sequence_to_Go(&m.Dim, *(*std_msgs_msg.CMultiArrayDimension__Sequence)(unsafe.Pointer(&mem.dim)))
	m.DataOffset = uint32(mem.data_offset)
}

func (t _MultiArrayLayoutTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__MultiArrayLayout())
}

type CMultiArrayLayout = C.example_interfaces__msg__MultiArrayLayout
type CMultiArrayLayout__Sequence = C.example_interfaces__msg__MultiArrayLayout__Sequence

func MultiArrayLayout__Sequence_to_Go(goSlice *[]MultiArrayLayout, cSlice CMultiArrayLayout__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MultiArrayLayout, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__MultiArrayLayout__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__MultiArrayLayout * uintptr(i)),
		))
		MultiArrayLayoutTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func MultiArrayLayout__Sequence_to_C(cSlice *CMultiArrayLayout__Sequence, goSlice []MultiArrayLayout) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__MultiArrayLayout)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__MultiArrayLayout * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__MultiArrayLayout)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__MultiArrayLayout * uintptr(i)),
		))
		MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func MultiArrayLayout__Array_to_Go(goSlice []MultiArrayLayout, cSlice []CMultiArrayLayout) {
	for i := 0; i < len(cSlice); i++ {
		MultiArrayLayoutTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MultiArrayLayout__Array_to_C(cSlice []CMultiArrayLayout, goSlice []MultiArrayLayout) {
	for i := 0; i < len(goSlice); i++ {
		MultiArrayLayoutTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
