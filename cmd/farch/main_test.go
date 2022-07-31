package main

import (
	"bytes"
	"os"
	"testing"

	. "github.com/jjcapellan/jj-archiver"
)

func isSameSize(file1 string, file2 string) bool {
	b1, _ := ReadFile(file1)
	b2, _ := ReadFile(file2)

	return len(b1) == len(b2)
}

func isSameFile(file1 string, file2 string) bool {

	b1, _ := ReadFile(file1)
	b2, _ := ReadFile(file2)

	return bytes.Equal(b1, b2)
}

func TestBackup(t *testing.T) {
	input := "test_assets/testfolder"
	output := "test_assets/temp_file.crp"
	pass := "abcde"

	err := backup(input, output, pass)
	defer os.Remove(output)
	if err != nil {
		t.Fatalf("Error on backup: %s", err.Error())
	}

	// Only file size can be compared. In each encription generated keys are different.
	if !isSameSize("test_assets/testfolder.crp", output) {
		t.Fatalf("Error on backup: generated file not valid")
	}
}

func TestBackupFile(t *testing.T) {
	input := "test_assets/testfolder/file.txt"
	output := "test_assets/temp_filetxt.crp"
	pass := "abcde"

	err := backup(input, output, pass)
	defer os.Remove(output)
	if err != nil {
		t.Fatalf("Error on backup: %s", err.Error())
	}

	// Only file size can be compared. In each encription generated keys are different.
	if !isSameSize("test_assets/filetxt.crp", output) {
		t.Fatalf("Error on backup: generated file not valid")
	}
}

func TestRestore(t *testing.T) {
	input := "test_assets/testfolder.crp"
	output := "."
	pass := "abcde"

	err := restore(input, output, pass)
	defer os.RemoveAll("./testfolder")
	if err != nil {
		t.Fatalf("Error on restore: %s", err.Error())
	}

	if !isSameFile("test_assets/testfolder/samples2/file3.txt", "testfolder/samples2/file3.txt") {
		t.Fatalf("Error on restore: restored files not valid")
	}
}
