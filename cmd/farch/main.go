package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	anim "github.com/jjcapellan/farch-cli/pkg/animation"
	"github.com/jjcapellan/farch-cli/pkg/pass"
	. "github.com/jjcapellan/jj-archiver"
)

var command, input, output, password string
var tFlag, vFlag, hFlag bool
var wg sync.WaitGroup

func main() {

	showVersion()

	flag.BoolVar(&tFlag, "t", false, "Shows execution time")
	flag.BoolVar(&vFlag, "version", false, "Shows program version")
	flag.BoolVar(&hFlag, "h", false, "Shows help")
	flag.BoolVar(&hFlag, "help", false, "Shows help")

	flag.Usage = showHelp

	flag.Parse()

	if hFlag {
		showHelp()
		os.Exit(0)
	}

	if vFlag {
		os.Exit(0)
	}

	err := validateArgs(flag.Args())
	if err != nil {
		showHelp()
		fmt.Println(err.Error())
		os.Exit(1)
	}

	password = pass.GetPassword()

	startTime := time.Now()

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

	if tFlag {
		elapsed := time.Since(startTime)
		fmt.Printf("Elapsed time: %s\n", elapsed)
	}

	os.Exit(0)
}

func backup(inputPath string, outputPath string, password string) error {

	wg.Add(1)
	go func() {
		defer wg.Done()
		anim.AnimatedInfo("Packaging...[1/4]")
	}()

	var packedData []byte
	var err error
	if isDir(inputPath) {
		packedData, err = PackFolder(inputPath)
	} else {
		var pathsArray = []string{inputPath}
		packedData, err = PackArray(pathsArray)
	}
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

func isDir(path string) bool {
	fileInfo, _ := os.Stat(path)
	return fileInfo.Mode().IsDir()
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
	fmt.Println(HELP)
}

func showVersion() {
	fmt.Println(HEADER)
}
