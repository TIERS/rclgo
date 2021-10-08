/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/TIERS/rclgo/pkg/rclgo/types"
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/TIERS/rclgo/internal/msgs/std_msgs/msg"
	primitives "github.com/TIERS/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/joy.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/Joy", JoyTypeSupport)
}

// Do not create instances of this type directly. Always use NewJoy
// function instead.
type Joy struct {
	Header std_msgs_msg.Header `yaml:"header"`// The timestamp is the time at which data is received from the joystick.
	Axes []float32 `yaml:"axes"`// The axes measurements from a joystick.
	Buttons []int32 `yaml:"buttons"`// The buttons measurements from a joystick.
}

// NewJoy creates a new Joy with default values.
func NewJoy() *Joy {
	self := Joy{}
	self.SetDefaults()
	return &self
}

func (t *Joy) Clone() *Joy {
	c := &Joy{}
	c.Header = *t.Header.Clone()
	if t.Axes != nil {
		c.Axes = make([]float32, len(t.Axes))
		copy(c.Axes, t.Axes)
	}
	if t.Buttons != nil {
		c.Buttons = make([]int32, len(t.Buttons))
		copy(c.Buttons, t.Buttons)
	}
	return c
}

func (t *Joy) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Joy) SetDefaults() {
	t.Header.SetDefaults()
	t.Axes = nil
	t.Buttons = nil
}

// CloneJoySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneJoySlice(dst, src []Joy) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var JoyTypeSupport types.MessageTypeSupport = _JoyTypeSupport{}

type _JoyTypeSupport struct{}

func (t _JoyTypeSupport) New() types.Message {
	return NewJoy()
}

func (t _JoyTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__Joy
	return (unsafe.Pointer)(C.sensor_msgs__msg__Joy__create())
}

func (t _JoyTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__Joy__destroy((*C.sensor_msgs__msg__Joy)(pointer_to_free))
}

func (t _JoyTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Joy)
	mem := (*C.sensor_msgs__msg__Joy)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.Float32__Sequence_to_C((*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.axes)), m.Axes)
	primitives.Int32__Sequence_to_C((*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.buttons)), m.Buttons)
}

func (t _JoyTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Joy)
	mem := (*C.sensor_msgs__msg__Joy)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.Float32__Sequence_to_Go(&m.Axes, *(*primitives.CFloat32__Sequence)(unsafe.Pointer(&mem.axes)))
	primitives.Int32__Sequence_to_Go(&m.Buttons, *(*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.buttons)))
}

func (t _JoyTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__Joy())
}

type CJoy = C.sensor_msgs__msg__Joy
type CJoy__Sequence = C.sensor_msgs__msg__Joy__Sequence

func Joy__Sequence_to_Go(goSlice *[]Joy, cSlice CJoy__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Joy, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__Joy__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__Joy * uintptr(i)),
		))
		JoyTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Joy__Sequence_to_C(cSlice *CJoy__Sequence, goSlice []Joy) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__Joy)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__Joy * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__Joy)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__Joy * uintptr(i)),
		))
		JoyTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Joy__Array_to_Go(goSlice []Joy, cSlice []CJoy) {
	for i := 0; i < len(cSlice); i++ {
		JoyTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Joy__Array_to_C(cSlice []CJoy, goSlice []Joy) {
	for i := 0; i < len(goSlice); i++ {
		JoyTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
