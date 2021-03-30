// THIS FILE IS AUTOGENERATED BY 'rclgo generate'
package rosidl_runtime_c

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo CFLAGS: -I/opt/ros/foxy/include

#include "rosidl_runtime_c/string.h"
#include "rosidl_runtime_c/primitives_sequence.h"

*/
import "C"
import (
	"unsafe"
)


// Bool
type CBool = C.bool
type CBool__Sequence = C.rosidl_runtime_c__boolean__Sequence

func Bool__Sequence_to_Go(goSlice *[]bool, cSlice CBool__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]bool, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.bool)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_bool * uintptr(i)),
		))
		(*goSlice)[i] = bool(*cIdx)
	}
}
func Bool__Sequence_to_C(cSlice *CBool__Sequence, goSlice []bool) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.bool)(C.malloc((C.size_t)(C.sizeof_bool * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.bool)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_bool * uintptr(i)),
		))
		*cIdx = (C.bool)(v)
	}
}
func Bool__Array_to_Go(goSlice []bool, cSlice []CBool) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i] = bool(cSlice[i])
	}
}
func Bool__Array_to_C(cSlice []CBool, goSlice []bool) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = C.bool(goSlice[i])
	}
}
// Byte
type CByte = C.uint8_t
type CByte__Sequence = C.rosidl_runtime_c__octet__Sequence

func Byte__Sequence_to_Go(goSlice *[]byte, cSlice CByte__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]byte, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.uint8_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_uint8_t * uintptr(i)),
		))
		(*goSlice)[i] = byte(*cIdx)
	}
}
func Byte__Sequence_to_C(cSlice *CByte__Sequence, goSlice []byte) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.uint8_t)(C.malloc((C.size_t)(C.sizeof_uint8_t * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.uint8_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_uint8_t * uintptr(i)),
		))
		*cIdx = (C.uint8_t)(v)
	}
}
func Byte__Array_to_Go(goSlice []byte, cSlice []CByte) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i] = byte(cSlice[i])
	}
}
func Byte__Array_to_C(cSlice []CByte, goSlice []byte) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = C.uint8_t(goSlice[i])
	}
}
// Float32
type CFloat32 = C.float
type CFloat32__Sequence = C.rosidl_runtime_c__float__Sequence

func Float32__Sequence_to_Go(goSlice *[]float32, cSlice CFloat32__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]float32, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.float)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_float * uintptr(i)),
		))
		(*goSlice)[i] = float32(*cIdx)
	}
}
func Float32__Sequence_to_C(cSlice *CFloat32__Sequence, goSlice []float32) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.float)(C.malloc((C.size_t)(C.sizeof_float * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.float)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_float * uintptr(i)),
		))
		*cIdx = (C.float)(v)
	}
}
func Float32__Array_to_Go(goSlice []float32, cSlice []CFloat32) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i] = float32(cSlice[i])
	}
}
func Float32__Array_to_C(cSlice []CFloat32, goSlice []float32) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = C.float(goSlice[i])
	}
}
// Float64
type CFloat64 = C.double
type CFloat64__Sequence = C.rosidl_runtime_c__double__Sequence

func Float64__Sequence_to_Go(goSlice *[]float64, cSlice CFloat64__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]float64, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.double)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_double * uintptr(i)),
		))
		(*goSlice)[i] = float64(*cIdx)
	}
}
func Float64__Sequence_to_C(cSlice *CFloat64__Sequence, goSlice []float64) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.double)(C.malloc((C.size_t)(C.sizeof_double * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.double)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_double * uintptr(i)),
		))
		*cIdx = (C.double)(v)
	}
}
func Float64__Array_to_Go(goSlice []float64, cSlice []CFloat64) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i] = float64(cSlice[i])
	}
}
func Float64__Array_to_C(cSlice []CFloat64, goSlice []float64) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = C.double(goSlice[i])
	}
}
// Int16
type CInt16 = C.int16_t
type CInt16__Sequence = C.rosidl_runtime_c__int16__Sequence

func Int16__Sequence_to_Go(goSlice *[]int16, cSlice CInt16__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]int16, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.int16_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_int16_t * uintptr(i)),
		))
		(*goSlice)[i] = int16(*cIdx)
	}
}
func Int16__Sequence_to_C(cSlice *CInt16__Sequence, goSlice []int16) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.int16_t)(C.malloc((C.size_t)(C.sizeof_int16_t * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.int16_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_int16_t * uintptr(i)),
		))
		*cIdx = (C.int16_t)(v)
	}
}
func Int16__Array_to_Go(goSlice []int16, cSlice []CInt16) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i] = int16(cSlice[i])
	}
}
func Int16__Array_to_C(cSlice []CInt16, goSlice []int16) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = C.int16_t(goSlice[i])
	}
}
// Int32
type CInt32 = C.int32_t
type CInt32__Sequence = C.rosidl_runtime_c__int32__Sequence

func Int32__Sequence_to_Go(goSlice *[]int32, cSlice CInt32__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]int32, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.int32_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_int32_t * uintptr(i)),
		))
		(*goSlice)[i] = int32(*cIdx)
	}
}
func Int32__Sequence_to_C(cSlice *CInt32__Sequence, goSlice []int32) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.int32_t)(C.malloc((C.size_t)(C.sizeof_int32_t * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.int32_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_int32_t * uintptr(i)),
		))
		*cIdx = (C.int32_t)(v)
	}
}
func Int32__Array_to_Go(goSlice []int32, cSlice []CInt32) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i] = int32(cSlice[i])
	}
}
func Int32__Array_to_C(cSlice []CInt32, goSlice []int32) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = C.int32_t(goSlice[i])
	}
}
// Int64
type CInt64 = C.int64_t
type CInt64__Sequence = C.rosidl_runtime_c__int64__Sequence

func Int64__Sequence_to_Go(goSlice *[]int64, cSlice CInt64__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]int64, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.int64_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_int64_t * uintptr(i)),
		))
		(*goSlice)[i] = int64(*cIdx)
	}
}
func Int64__Sequence_to_C(cSlice *CInt64__Sequence, goSlice []int64) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.int64_t)(C.malloc((C.size_t)(C.sizeof_int64_t * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.int64_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_int64_t * uintptr(i)),
		))
		*cIdx = (C.int64_t)(v)
	}
}
func Int64__Array_to_Go(goSlice []int64, cSlice []CInt64) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i] = int64(cSlice[i])
	}
}
func Int64__Array_to_C(cSlice []CInt64, goSlice []int64) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = C.int64_t(goSlice[i])
	}
}
// Int8
type CInt8 = C.int8_t
type CInt8__Sequence = C.rosidl_runtime_c__int8__Sequence

func Int8__Sequence_to_Go(goSlice *[]int8, cSlice CInt8__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]int8, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.int8_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_int8_t * uintptr(i)),
		))
		(*goSlice)[i] = int8(*cIdx)
	}
}
func Int8__Sequence_to_C(cSlice *CInt8__Sequence, goSlice []int8) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.int8_t)(C.malloc((C.size_t)(C.sizeof_int8_t * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.int8_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_int8_t * uintptr(i)),
		))
		*cIdx = (C.int8_t)(v)
	}
}
func Int8__Array_to_Go(goSlice []int8, cSlice []CInt8) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i] = int8(cSlice[i])
	}
}
func Int8__Array_to_C(cSlice []CInt8, goSlice []int8) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = C.int8_t(goSlice[i])
	}
}
// Uint16
type CUint16 = C.uint16_t
type CUint16__Sequence = C.rosidl_runtime_c__uint16__Sequence

func Uint16__Sequence_to_Go(goSlice *[]uint16, cSlice CUint16__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]uint16, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.uint16_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_uint16_t * uintptr(i)),
		))
		(*goSlice)[i] = uint16(*cIdx)
	}
}
func Uint16__Sequence_to_C(cSlice *CUint16__Sequence, goSlice []uint16) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.uint16_t)(C.malloc((C.size_t)(C.sizeof_uint16_t * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.uint16_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_uint16_t * uintptr(i)),
		))
		*cIdx = (C.uint16_t)(v)
	}
}
func Uint16__Array_to_Go(goSlice []uint16, cSlice []CUint16) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i] = uint16(cSlice[i])
	}
}
func Uint16__Array_to_C(cSlice []CUint16, goSlice []uint16) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = C.uint16_t(goSlice[i])
	}
}
// Uint32
type CUint32 = C.uint32_t
type CUint32__Sequence = C.rosidl_runtime_c__uint32__Sequence

func Uint32__Sequence_to_Go(goSlice *[]uint32, cSlice CUint32__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]uint32, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.uint32_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_uint32_t * uintptr(i)),
		))
		(*goSlice)[i] = uint32(*cIdx)
	}
}
func Uint32__Sequence_to_C(cSlice *CUint32__Sequence, goSlice []uint32) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.uint32_t)(C.malloc((C.size_t)(C.sizeof_uint32_t * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.uint32_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_uint32_t * uintptr(i)),
		))
		*cIdx = (C.uint32_t)(v)
	}
}
func Uint32__Array_to_Go(goSlice []uint32, cSlice []CUint32) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i] = uint32(cSlice[i])
	}
}
func Uint32__Array_to_C(cSlice []CUint32, goSlice []uint32) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = C.uint32_t(goSlice[i])
	}
}
// Uint64
type CUint64 = C.uint64_t
type CUint64__Sequence = C.rosidl_runtime_c__uint64__Sequence

func Uint64__Sequence_to_Go(goSlice *[]uint64, cSlice CUint64__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]uint64, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.uint64_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_uint64_t * uintptr(i)),
		))
		(*goSlice)[i] = uint64(*cIdx)
	}
}
func Uint64__Sequence_to_C(cSlice *CUint64__Sequence, goSlice []uint64) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.uint64_t)(C.malloc((C.size_t)(C.sizeof_uint64_t * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.uint64_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_uint64_t * uintptr(i)),
		))
		*cIdx = (C.uint64_t)(v)
	}
}
func Uint64__Array_to_Go(goSlice []uint64, cSlice []CUint64) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i] = uint64(cSlice[i])
	}
}
func Uint64__Array_to_C(cSlice []CUint64, goSlice []uint64) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = C.uint64_t(goSlice[i])
	}
}
// Uint8
type CUint8 = C.uint8_t
type CUint8__Sequence = C.rosidl_runtime_c__uint8__Sequence

func Uint8__Sequence_to_Go(goSlice *[]uint8, cSlice CUint8__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]uint8, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.uint8_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_uint8_t * uintptr(i)),
		))
		(*goSlice)[i] = uint8(*cIdx)
	}
}
func Uint8__Sequence_to_C(cSlice *CUint8__Sequence, goSlice []uint8) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.uint8_t)(C.malloc((C.size_t)(C.sizeof_uint8_t * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.uint8_t)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_uint8_t * uintptr(i)),
		))
		*cIdx = (C.uint8_t)(v)
	}
}
func Uint8__Array_to_Go(goSlice []uint8, cSlice []CUint8) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i] = uint8(cSlice[i])
	}
}
func Uint8__Array_to_C(cSlice []CUint8, goSlice []uint8) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = C.uint8_t(goSlice[i])
	}
}
