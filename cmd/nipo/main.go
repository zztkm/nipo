package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/zztkm/nipo"
)

const (
	name     = "nipo"
	version  = "0.1.0"
	revision = "HEAD"
)

var printVersion = flag.Bool("version", false, "print version")


func createConfigFileName() string {
	file := "token.json"

	if runtime.GOOS == "windows" {
		file = filepath.Join(os.Getenv("APPDATA"), "nipo", file)
	} else {
		file = filepath.Join(os.Getenv("HOME"), ".config", "nipo", file)
	}

	return file
}

func getConfig() nipo.Config {
	file := createConfigFileName()

	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Failed to read config file: ", err)
		os.Exit(1)
	}
	var c nipo.Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		fmt.Println("Failed to unmarshal file: ", err)
		os.Exit(1)
	}
	return c
}

func fatal(format string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, format, err)
	} else {
		fmt.Fprint(os.Stderr, format)
	}
	os.Exit(1)
}

func usage() {
	fmt.Println(`
nipo <command> [arguments]
	init 日報のテンプレートを出力します
		$ nipo init [mac or vm]
	converte Markdown書式のファイルをNIPO書式に変換して標準出力します
		$ nipo converte file.md
	`)
	os.Exit(1)
}

func main() {
	flag.Parse()

	if *printVersion {
		fmt.Printf("%s %s (rev: %s/%s)\n", name, version, revision, runtime.Version())
		return
	}

	if flag.NArg() <= 1 {
		usage()
	}

	switch flag.Arg(0) {
	case "i", "init":
		pattern := flag.Arg(1)
		nipo.PrintTemplate(pattern)
	case "c", "converte":
		filename := flag.Arg(1)
		nipo.MarkdownToNipo(filename)
	default:
		usage()
	}
}
