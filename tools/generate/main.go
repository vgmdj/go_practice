package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "f", "", "filename")
}

func main() {
	flag.Parse()
	log.Println("input filename: ", fileName)

	dirName := fmt.Sprintf("leetcode/%s", dealFileName(fileName))
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}

	touchFile(dirName)

}

func dealFileName(str string) string {
	dArr := strings.Split(str, ".")
	if len(dArr) != 2 {
		panic("to many dot")
	}

	dArr[0] = "0000" + dArr[0]
	dArr[0] = dArr[0][len(dArr[0])-4:]

	dArr[1] = strings.Join(strings.Split(dArr[1], " "), "_")
	dArr[1] = strings.TrimLeft(dArr[1], "_")

	return strings.Join(dArr, ".")
}

func touchFile(dirName string) {
	fArr := strings.Split(dirName, ".")
	fn := fArr[len(fArr)-1]

	content := fmt.Sprintf("package %s", fn)
	f, err := os.Create(fmt.Sprintf("%s/%s.go", dirName, fn))
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	_, err = f.Write([]byte(content))
	if err != nil {
		panic(err.Error())
	}
}
