
build:
	go build -o bin/nipo cmd/nipo/main.go 

windows_build:
	GOOS=windows go build -o bin/nipo.exe cmd/nipo/main.go 
