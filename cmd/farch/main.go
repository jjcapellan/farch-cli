package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jjcapellan/farch-cli/pkg/pass"
)

var command, input, output, password string

func main() {

	showTitle()

	err := validateArgs(flag.Args())
	if err != nil {
		showHelp()
		fmt.Println(err.Error())
		os.Exit(1)
	}

	password = pass.GetPassword()
}

func backup() error {
	return nil
}

func restore() error {
	return nil
}

func showHelp() {
	fmt.Println("Usage:\n" +
		"    farch command input [output]" +
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
		"output_folder = current directory\n\n ")
}

func showTitle() {
	fmt.Println("************************\n***** farch v0.1.0 *****\n************************" +
		"\nfarch is a command line utility to pack, compress and encrypt a folder and save it into a file.\n" +
		"To get help use:\nfarch { -h | --help }")
}
