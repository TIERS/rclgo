/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces_srv

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <example_interfaces/srv/trigger.h>
*/
import "C"

import (
	"github.com/TIERS/rclgo/pkg/rclgo/typemap"
	"github.com/TIERS/rclgo/pkg/rclgo/types"

	"unsafe"
)

func init() {
	typemap.RegisterService("example_interfaces/Trigger", TriggerTypeSupport)
}

type _TriggerTypeSupport struct {}

func (s _TriggerTypeSupport) Request() types.MessageTypeSupport {
	return Trigger_RequestTypeSupport
}

func (s _TriggerTypeSupport) Response() types.MessageTypeSupport {
	return Trigger_ResponseTypeSupport
}

func (s _TriggerTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__example_interfaces__srv__Trigger())
}

// Modifying this variable is undefined behavior.
var TriggerTypeSupport types.ServiceTypeSupport = _TriggerTypeSupport{}
