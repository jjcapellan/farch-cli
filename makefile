DST_LINUX=./build/linux/
DST_WIN=./build/win64/

build_linux:
	GOARCH=amd64 GOOS=linux go build -o ${DST_LINUX}farch ./cmd/farch/main.go

build_win:
	GOARCH=amd64 GOOS=windows go build -o ${DST_WIN}farch.exe ./cmd/farch/main.go

all: build_linux build_win