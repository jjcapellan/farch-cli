# FARCH
Farch is a simple command line utility to pack, compress and encrypt files and folders into a single file.  
  
---

## INSTALATION
Farch is compiled into a single file with all dependencies included. So, simply download the file and assign execution permissions if necessary.  
### Linux
1. Download executable here:  https://github.com/jjcapellan/farch-cli/releases/download/v0.1.0/farch
2. Copy it to any location within the system PATH (Ex: <code>cp ./farch /usr/local/bin</code>)
3. If you are not allowed to execute the file, then you have to change file permissions (Ex: <code>chmod 755 ./farch</code>)
### Windows
1. Download executable here: https://github.com/jjcapellan/farch-cli/releases/download/v0.1.0/farch.exe
2. Copy it to any location within the system PATH or add new location to the system PATH:  
Control Panel > Advanced system settings > Environment Variables > System Variables > Edit System Variable (PATH)  
---

## USAGE
```
farch [options] command input [output]
```

### Available commands
* *backup*  : archives a file or folder
* *restore* : restores a backup file

### Available options:
* *-h, --help*  : Shows help (disables any command execution)
* *-t*          : shows execution time
* *--version*   : Shows program version (disables any command execution)

### Examples:
```
$ farch backup projectsfolder backups/projects.crp
$ farch -t backup projectsfolder
$ farch -t restore backups/projects.crp destFolder
$ farch restore backups/projects.crp
$ farch --version
$ farch --help
```
### Defaults:
* output_file_path = bk_+ *base path of input_folder* + .crp (Ex: root/fold1/fold2 -> bk_fold2.crp)
* output_folder = *current directory*


## License
**FARCH** is licensed under the terms of the MIT open source license.