package main

const VERSION string = "v0.1.0"

const HEADER string = "--- farch CLI " + VERSION + " ---" +
	"\nCopyright © 2022 Juan Jose Capellan" +
	"\nLicense MIT: <https://raw.githubusercontent.com/jjcapellan/farch-cli/master/LICENSE>" +
	"\nfarch is a command line utility to pack, compress and encrypt files and folders and save them into a file.\n "

const HELP string = "\nUsage:\n" +
	"    farch [options] command input [output]\n" +
	"\nAvailable commands:\n" +
	"    * backup   : archives a file or folder\n" +
	"    * restore  : restores backup file\n" +
	"\nAvailable options:\n" +
	"    -h, --help : Shows help\n" +
	"    -t         : Shows execution time\n" +
	"    --version  : Shows program version\n" +
	"\nExamples:\n" +
	"    farch backup projectsfolder backups/projects.crp\n" +
	"    farch -t backup projectsfolder\n" +
	"    farch -t restore backups/projects.crp destFolder\n" +
	"    farch restore backups/projects.crp\n" +
	"\nDefaults:\n" +
	"output_file_path = bk_+ base path of input_folder + .crp (Ex: root/fold1/fold2 -> bk_fold2.crp)\n" +
	"output_folder = current directory\n "
