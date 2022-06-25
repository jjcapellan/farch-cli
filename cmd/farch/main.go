package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jjcapellan/farch-cli/pkg/pass"
	. "github.com/jjcapellan/jj-archiver"
)

var command, input, output, password string

func main() {

	showTitle()

	flag.Parse()

	err := validateArgs(flag.Args())
	if err != nil {
		showHelp()
		fmt.Println(err.Error())
		os.Exit(1)
	}

	password = pass.GetPassword()

	if command == "backup" {
		err = backup(input, output, password)
		if err != nil {
			showHelp()
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	if command == "restore" {
		err = restore(input, output, password)
		if err != nil {
			showHelp()
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}

func backup(inputPath string, outputPath string, password string) error {
	packedData, err := PackFolder(inputPath)
	if err != nil {
		return err
	}

	gzipData, err := Compress(packedData, "file.gzip")
	if err != nil {
		return err
	}

	encryptData, err := Encrypt(gzipData, password)
	if err != nil {
		return err
	}

	err = WriteFile(outputPath, encryptData, 0666)

	return err
}

func restore(inputPath string, outputPath string, password string) error {
	data, err := ReadFile(inputPath)
	if err != nil {
		return err
	}

	decryptData, err := Decrypt(data, password)
	if err != nil {
		return err
	}

	gzipData, _, err := Decompress(decryptData)
	if err != nil {
		return err
	}

	err = Unpack(gzipData, outputPath)

	return err
}

func showHelp() {
	fmt.Println("\nUsage:\n" +
		"    farch command input [output]\n" +
		"** Backup:\n" +
		"    farch backup input_folder [output_file_path]\n" +
		"** Restore:\n" +
		"    farch restore input_file [output_folder]\n" +
		"\nExamples:\n" +
		"    farch backup projectsfolder backups/projects.crp\n" +
		"    farch backup projectsfolder\n" +
		"    farch restore backups/projects.crp destFolder\n" +
		"    farch restore backups/projects.crp\n" +
		"\nDefaults:\n" +
		"output_file_path = bk_+ base path of input_folder + .crp (Ex: root/fold1/fold2 -> bk_fold2.crp)\n" +
		"output_folder = current directory\n ")
}

func showTitle() {
	fmt.Println("************************\n***** farch v0.1.0 *****\n************************" +
		"\nfarch is a command line utility to pack, compress and encrypt a folder and save it into a file.\n" +
		"To get help use:\nfarch { -h | --help }")
}
