/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/TIERS/rclgo/pkg/gogen"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func validateGenerateArgs(cmd *cobra.Command, args []string) error {
	if getString(cmd, "root-path") == "" {
		if os.Getenv("AMENT_PREFIX_PATH") == "" {
			return fmt.Errorf("You haven't sourced your ROS2 environment! Cannot autodetect --root-path. Source your ROS2 or pass --root-path")
		}
		return fmt.Errorf("root-path is required")
	}
	_, err := os.Stat(getString(cmd, "root-path"))
	if err != nil {
		return fmt.Errorf("root-path error: %v", err)
	}
	if getString(cmd, "dest-path") == "" {
		return fmt.Errorf("dest-path is required")
	}
	_, err = os.Stat(getString(cmd, "dest-path"))
	if err != nil {
		return fmt.Errorf("dest-path error: %v", err)
	}
	return nil
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Go bindings for ROS2 interface definitions under <root-path>",
	RunE: func(cmd *cobra.Command, args []string) error {
		rootPath := getString(cmd, "root-path")
		destPath := getString(cmd, "dest-path")
		config := getGogenConfig(cmd)
		if err := gogen.GenerateGolangMessageTypes(config, rootPath, destPath); err != nil {
			return fmt.Errorf("failed to generate interface bindings: %w", err)
		}
		if err := gogen.GenerateROS2AllMessagesImporter(config, destPath); err != nil {
			return fmt.Errorf("failed to generate all importer: %w", err)
		}
		return nil
	},
	Args: validateGenerateArgs,
}

var generateRclgoCmd = &cobra.Command{
	Use:   "generate-rclgo",
	Short: "Generate Go code that forms a part of rclgo",
	RunE: func(cmd *cobra.Command, args []string) error {
		rootPath := getString(cmd, "root-path")
		destPath := getString(cmd, "dest-path")
		config := getGogenConfig(cmd)
		if err := gogen.GeneratePrimitives(config, destPath); err != nil {
			return fmt.Errorf("failed to generate primitive types: %w", err)
		}
		if err := gogen.GenerateROS2ErrorTypes(rootPath, destPath); err != nil {
			return fmt.Errorf("failed to generate error types: %w", err)
		}
		return nil
	},
	Args: validateGenerateArgs,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	configureFlags(generateCmd, ".")

	rootCmd.AddCommand(generateRclgoCmd)
	configureFlags(generateRclgoCmd, gogen.RclgoRepoRootPath())
}

func configureFlags(cmd *cobra.Command, destPathDefault string) {
	cmd.PersistentFlags().StringP("root-path", "r", os.Getenv("AMENT_PREFIX_PATH"), "Root lookup path for ROS2 .msg files. If ROS2 environment is sourced, is autodetected.")
	cmd.PersistentFlags().StringP("dest-path", "d", destPathDefault, "Destination directory for the Golang typed converted ROS2 messages. ROS2 Message structure is preserved as <ros2-package>/msg/<msg-name>")
	cmd.PersistentFlags().String("rclgo-import-path", gogen.DefaultConfig.RclgoImportPath, "Import path of rclgo library")
	cmd.PersistentFlags().String("message-module-prefix", gogen.DefaultConfig.MessageModulePrefix, "Import path prefix for generated message binding modules")
	bindPFlags(cmd)
}

func getPrefix(cmd *cobra.Command) string {
	parts := []string{}
	for c := cmd; c != c.Root(); c = c.Parent() {
		parts = append(parts, c.Name())
	}
	for i := 0; i < len(parts)/2; i++ {
		parts[i], parts[len(parts)-i-1] = parts[len(parts)-i-1], parts[i]
	}
	prefix := strings.Join(parts, ".")
	if prefix != "" {
		prefix += "."
	}
	return prefix
}

func getString(cmd *cobra.Command, key string) string {
	return viper.GetString(getPrefix(cmd) + key)
}

func bindPFlags(cmd *cobra.Command) {
	prefix := getPrefix(cmd)
	cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		viper.BindPFlag(prefix+f.Name, f)
	})
	cmd.LocalFlags().VisitAll(func(f *pflag.Flag) {
		viper.BindPFlag(prefix+f.Name, f)
	})
}

func getGogenConfig(cmd *cobra.Command) *gogen.Config {
	return &gogen.Config{
		RclgoImportPath:     getString(cmd, "rclgo-import-path"),
		MessageModulePrefix: getString(cmd, "message-module-prefix"),
	}
}
