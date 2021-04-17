package nipo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// MarkdownToNipo マークダウン書式のファイルをNIPO書式に変換します
func MarkdownToNipo(filename string) error {
	err := readLine(filename)
	if err != nil {
		return err
	}
	return nil
}

func readLine(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		strLine := string(line)

		if strings.HasPrefix(strLine, "## ") {
			strLine = strings.Replace(strLine, "## ", "●", -1)
		}

		fmt.Print(strLine)
		if !isPrefix {
			fmt.Println()
		}
	}
	return nil
}
