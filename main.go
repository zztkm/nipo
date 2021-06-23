package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/Jeffail/gabs/v2"
	"github.com/pkg/browser"
)

const (
	name     = "nipo"
	version  = "0.1.0"
	revision = "HEAD"
)

var configFile = "nipo.json"

var printVersion = flag.Bool("version", false, "print version")

const layout = "2006-01-02"

// dir file 問わず存在確認に使える
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
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
	generate 実行時の日付でファイルを作成します。昨日のファイルが存在する場合は内容をコピーします。
		$ nipo generate
	sink Markdown書式のファイルをNIPO書式に変換して標準出力します
		$ nipo sink タスクトラッカー sink をブラウザで開きます
	`)
	os.Exit(1)
}

func converte(filename string) {

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		fatal("Failed to read config file: %s\n", err)
	}

	jsonParsed, err := gabs.ParseJSON(b)
	if err != nil {
		fatal("Failed to unmarshal file: %s\n", err)
	}

	file, err := os.Open(filename)
	if err != nil {
		fatal("Failed to open file: %s\n", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fatal("Failed to readLine: %s\n", err)
		}

		strLine := string(line)

		for key, child := range jsonParsed.Search("header").ChildrenMap() {
			// fmt.Printf("key: %v, value: %v\n", key, child.String())
			if strings.HasPrefix(strLine, key) {
				strLine = strings.Replace(strLine, key, child.Data().(string), -1)
			}
		}

		fmt.Print(strLine)
		if !isPrefix {
			fmt.Println()
		}
	}
}

func copyFile(src, dst string) {
	w, err := os.Create(dst)
	if err != nil {
		fatal("Failed to create file: %s\n", err)
	}
	defer w.Close()

	if fileExists(src) {
		r, err := os.Open(src)
		if err != nil {
			fatal("Failed to open file: %s\n", err)
		}
		defer r.Close()

		_, err = io.Copy(w, r)
		if err != nil {
			fatal("Failed to copy: %s\n", err)
		}
	}
}

// init nipo.json が存在しない場合に作成. 既に存在する場合は処理しない.
func nipoInit() {

	if !fileExists(configFile) {
		f, err := os.Create(configFile)
		if err != nil {
		}
		defer f.Close()
	}
}

func main() {
	flag.Parse()

	if *printVersion {
		fmt.Printf("%s %s (rev: %s/%s)\n", name, version, revision, runtime.Version())
		return
	}

	if flag.NArg() == 0 {
		usage()
	}

	switch flag.Arg(0) {
	case "i", "init":
		nipoInit()
	case "c", "converte":
		filename := flag.Arg(1)
		converte(filename)
	case "sink":
		url := "https://veltiosoft.dev/sink/"
		browser.OpenURL(url)
	case "generate":
		today := time.Now()
		// today := time.Now().Format(layout)
		yesterday := today.AddDate(0, 0, -1)
		dst := today.Format(layout) + ".md"
		src := yesterday.Format(layout) + ".md"
		copyFile(src, dst)
	default:
		usage()
	}
}
