package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	iconv "github.com/djimenez/iconv-go"
)

func fileCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileFlag := flag.String("file", "", "path/to/file.csv")
	flag.Parse()
	input := *fileFlag
	if input == "" {
		log.Fatal()
	}

	file, err := os.Open(input)
	fileCheck(err)
	defer file.Close()

	//reader := csv.NewReader(file)
	converter, _ := iconv.NewReader(file, "sjis", "utf-8")
	reader := csv.NewReader(converter)
	//record, err := reader.Read()
	//fileCheck(err)

	export := fmt.Sprintf("./result/%s.csv", "data")
	wFile, _ := newFile(export)
	writer := bufio.NewWriter(wFile)
	for {
		record, err := reader.Read() // 1行読み出す
		if err == io.EOF {
			break
		} else {
			fileCheck(err)
		}
		sp := strings.Split(record[0], "/")
		out := fmt.Sprintf("%s, %s, %s\n", sp[0], sp[1], record[1])
		fmt.Print(out)
		writer.WriteString(out)
		writer.Flush()
	}
	wFile.Close()

}
func newFile(fn string) (*os.File, bool) {
	_, exist := os.Stat(fn)
	fp, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return fp, os.IsNotExist(exist)
}
