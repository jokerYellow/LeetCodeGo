package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	readMe := "# LeetCode\n |||\n| --- | --- |\n"
	leetcode := leetcodeItems()
	readMe += leetcode
	f, _ := os.Create("ReadMe.md")
	defer func() {
		f.Close()
	}()
	if f == nil {
		log.Fatal()
	}
	_, err := f.WriteString(readMe)
	if err != nil {
		log.Fatal(err)
	}
}

func leetcodeItems() string {
	dir := "./"
	files, error := ioutil.ReadDir(dir)
	if error != nil {
		log.Fatal(error)
	}
	b := new(bytes.Buffer)

	index := 1
	for _, f := range files {
		item := modified(f, dir)
		if item != nil {
			b.WriteString(fmt.Sprintf("| %d | %s |\n", index, *item))
			index++
		}
	}
	return b.String()
}

//change direcrory name componented by number and description
func modified(f os.FileInfo, dir string) *string {
	name := f.Name()
	number := getLeetcodeNumber(name)
	if len(number) == 0 {
		return nil
	}
	files, error := ioutil.ReadDir(dir + "/" + name)
	if error != nil {
		log.Fatal(error)
	}
	if len(files) == 0 {
		return nil
	}
	questionName := strings.Replace(files[0].Name(), ".go", "", 1)
	finalName := number + "-" + questionName
	os.Rename(dir+"/"+name, dir+"/"+finalName)
	show := "[" + finalName + "]" + "(" + "https://github.com/jokerYellow/LeetCodeGo/tree/master/" + finalName + ")"
	return &show
}

func getLeetcodeNumber(name string) string {
	reg, err := regexp.Compile("^[0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.FindString(name)
}
