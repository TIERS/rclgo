/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package gogen

import (
	"container/list"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

/*
Looks up where the rclgo-module is installed and returns a path where to write Golang bindings for ROS2 messages.
*/
func GetGoConvertedROS2MsgPackagesDir() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(file, "../..", "ros2/msgs")
}

func Generate(rootPath string, destPath string) {
	mds := list.New()
	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		matched, err := regexp.MatchString(`/msg/.+?\.msg$`, path)
		if err != nil {
			fmt.Printf("Error when matching path='%s' against regex='%s'", path, `\.msg$`)
		}
		if matched {
			fmt.Printf("Generating: %s\n", path)
			md, err := GenerateGolangTypeFromROS2MessagePath(path, destPath)
			if err != nil {
				fmt.Printf("Error converting ROS2 Message '%s' to '%s', error: %v\n", path, destPath, err)
			}
			mds.PushBack(md)
		}
		return nil
	})

	Generate_rosidl_runtime_c_sequence_handlers(destPath)
}
