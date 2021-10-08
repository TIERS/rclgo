/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package std_srvs_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_srvs__rosidl_typesupport_c -lstd_srvs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <std_srvs/srv/empty.h>
*/
import "C"

import (
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	"github.com/TIERS/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("std_srvs/Empty", EmptyTypeSupport)
}

type _EmptyTypeSupport struct {}

func (s _EmptyTypeSupport) Request() types.MessageTypeSupport {
	return Empty_RequestTypeSupport
}

func (s _EmptyTypeSupport) Response() types.MessageTypeSupport {
	return Empty_ResponseTypeSupport
}

func (s _EmptyTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__std_srvs__srv__Empty())
}

// Modifying this variable is undefined behavior.
var EmptyTypeSupport types.ServiceTypeSupport = _EmptyTypeSupport{}
