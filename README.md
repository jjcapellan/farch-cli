# FARCH
Farch is a simple command line utility to pack, compress and encrypt files and folders into a file.

# USAGE
```
farch [options] command input [output]
```
## Bakup file/folder:
```
farch [options] backup input_folder [output_file_path]
farch [options] backup input_file [output_file_path]
```
## Restore file/folder:
```
farch [options] restore input_file [output_folder]
```
## Options:
```
-t : shows execution time
```
## Examples:
```
$ farch backup projectsfolder backups/projects.crp
$ farch -t backup projectsfolder
$ farch -t restore backups/projects.crp destFolder
$ farch restore backups/projects.crp
```
## Defaults:
* output_file_path = bk_+ *base path of input_folder* + .crp (Ex: root/fold1/fold2 -> bk_fold2.crp)
* output_folder = *current directory*