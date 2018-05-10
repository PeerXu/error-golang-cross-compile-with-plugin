all:
	GOOS=linux GOARCH=arm GOARM=5 go build -o plugin.so -buildmode=plugin plugin/plugin.go
	GOOS=linux GOARCH=arm GOARM=5 go build -o hello main.go
