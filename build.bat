:: windows linux darwin freebsd
:: GOARCH 386 amd64 arm

set GOOS=linux
set GOARCH=amd64
go build -o ./bin/imageserver_linux main.go

set GOOS=windows
set GOARCH=amd64
go build -o ./bin/imageserver.exe main.go