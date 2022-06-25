package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func checkPath(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func validateArgs(args []string) error {

	if len(args) < 2 {
		showHelp()
		return errors.New("Error: Wrong command syntax")
	}

	// Command
	command = args[0]
	if command != "backup" && command != "restore" {
		showHelp()
		return errors.New("Error: Wrong command syntax")
	}

	if command == "backup" {
		return validateBackup(args)
	}

	return validateRestore(args)
}

func validateBackup(args []string) error {
	input = args[1]
	if input == "" {
		return errors.New("Error: Wrong command syntax")
	}

	if !checkPath(input) {
		return errors.New("Error: input folder not found")
	}

	if len(args) < 2 {
		currentDir, _ := os.Getwd()
		basePath := filepath.Base(currentDir)
		basePath = strings.ReplaceAll(basePath, ":", "")
		basePath = strings.ReplaceAll(basePath, ".", "")
		basePath = strings.ReplaceAll(basePath, "\\", "")
		output = "bk_" + basePath + ".crp"
		return nil
	}

	output = args[2]

	return nil
}

func validateRestore(args []string) error {
	input = args[1]
	if input == "" {
		return errors.New("Error: Wrong command syntax")
	}

	if !checkPath(input) {
		return errors.New("Error: input file not found")
	}

	if len(args) < 2 {
		output = "."
		return nil
	}

	output = args[2]

	return nil
}
