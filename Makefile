
build:
	go build -o bin/nipo main.go 

windows_build:
	GOOS=windows go build -o bin/nipo.exe main.go 
