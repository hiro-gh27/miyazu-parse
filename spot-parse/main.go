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
)

func fileCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileFlag := flag.String("file", "", "path/to/file.csv")
	flag.Parse()
	inputFile := *fileFlag
	if inputFile == "" {
		log.Fatal()
	}

	path := inputFile[:12]
	//date := csvPath[:4] + "/" + csvPath[4:6] + "/" + csvPath[6:8]
	date := path[:4] + "/" + path[4:6] + "/" + path[6:8]
	//fmt.Println(date)

	file, err := os.Open(inputFile)
	fileCheck(err)
	defer file.Close()

	reader := csv.NewReader(file)

	// 属性作り いる
	record, err := reader.Read()
	if err == io.EOF {
	} else {
		fileCheck(err)
	}

	column := fmt.Sprintf("%s, %s, %s", "日付", record[0], "人数")
	column += "\n"
	fmt.Printf(column)
	date = record[1]

	//天気csvを読み込んでから，場所.csvに書き込むイメージ
	for {
		record, err := reader.Read() // 1行読み出す
		if err == io.EOF {
			break
		} else {
			fileCheck(err)
		}
		name := strings.Replace(record[0], "/", "_", 1)
		name = strings.Replace(name, " ", "", -1)
		export := fmt.Sprintf("./result/%s.csv", name)
		fp, isFirst := newFile(export)
		writer := bufio.NewWriter(fp)
		if isFirst {
			_, err := writer.WriteString(column)
			if err != nil {
				log.Fatal(err)
			}
		}
		output := date
		for index := 0; index < len(record); index++ {
			output += ", " + record[index]
		}
		output += "\n"
		writer.WriteString(output)
		writer.Flush()
		fp.Close()
	}
}
func newFile(fn string) (*os.File, bool) {
	_, exist := os.Stat(fn)
	fp, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return fp, os.IsNotExist(exist)
}
