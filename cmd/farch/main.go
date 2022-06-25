package main

import (
	"fmt"

	"github.com/jjcapellan/farch-cli/pkg/pass"
)

var input, output, password string

func main() {
	showTitle()
	setArgs()
	checkArgs()
	password = pass.GetPassword()
}

func backup() error {
	return nil
}

func restore() error {
	return nil
}

func checkArgs() {

}

func setArgs() {

}

func showHelp() {

}

func showTitle() {
	fmt.Println("************************\n***** farch v0.1.0 *****\n************************" +
		"\nfarch is a command line utility to pack, compress and encrypt a folder and save it into a file.\n" +
		"To get help use:\nfarch { -h | --help }")
}
