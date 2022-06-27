package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	anim "github.com/jjcapellan/farch-cli/pkg/animation"
	"github.com/jjcapellan/farch-cli/pkg/pass"
	. "github.com/jjcapellan/jj-archiver"
)

var command, input, output, password string
var wg sync.WaitGroup

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
			anim.StopAnim("An error has occurred")
			wg.Wait()
			showHelp()
			fmt.Println(err.Error())
			os.Exit(1)
		}
		anim.StopAnim("Backup tasks completed")
		wg.Wait()
	}

	if command == "restore" {
		err = restore(input, output, password)
		if err != nil {
			anim.StopAnim("An error has occurred")
			wg.Wait()
			showHelp()
			fmt.Println(err.Error())
			os.Exit(1)
		}
		anim.StopAnim("Restore tasks completed")
		wg.Wait()
	}
}

func backup(inputPath string, outputPath string, password string) error {

	wg.Add(1)
	go func() {
		defer wg.Done()
		anim.AnimatedInfo("Packaging...[1/4]")
	}()

	packedData, err := PackFolder(inputPath)
	if err != nil {
		return err
	}

	anim.SetInfo("Compressing...[2/4]")
	gzipData, err := Compress(packedData, "file.gzip")
	if err != nil {
		return err
	}

	anim.SetInfo("Encrypting...[3/4]")
	encryptData, err := Encrypt(gzipData, password)
	if err != nil {
		return err
	}

	anim.SetInfo("Writing file...[4/4]")
	err = WriteFile(outputPath, encryptData, 0666)

	return err
}

func restore(inputPath string, outputPath string, password string) error {

	wg.Add(1)
	go func() {
		defer wg.Done()
		anim.AnimatedInfo("Reading file...[1/4]")
	}()

	data, err := ReadFile(inputPath)
	if err != nil {
		return err
	}

	anim.SetInfo("Decrypting...[2/4]")
	decryptData, err := Decrypt(data, password)
	if err != nil {
		return err
	}

	anim.SetInfo("Decompressing...[3/4]")
	gzipData, _, err := Decompress(decryptData)
	if err != nil {
		return err
	}

	anim.SetInfo("Unpacking...[4/4]")
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
	fmt.Println("--- farch CLI v0.1.0 ---" +
		"\nfarch is a command line utility to pack, compress and encrypt a folder and save it into a file.\n ")
}
