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

}

func showTitle() {
	fmt.Println("************************\n***** farch v0.1.0 *****\n************************" +
		"\nfarch is a command line utility to pack, compress and encrypt a folder and save it into a file.\n" +
		"To get help use:\nfarch { -h | --help }")
}
