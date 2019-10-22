package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	leetcode, count := leetcodeItems()

	readMe := fmt.Sprintf("# LeetCode\n >count:%d \n\n|#|Title|\n| --- | --- |\n", count)

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

func leetcodeItems() (string, int) {
	dir := "./"
	files, error := ioutil.ReadDir(dir)
	if error != nil {
		log.Fatal(error)
	}
	b := new(bytes.Buffer)
	sort.Slice(files, func(left, right int) bool {
		leftNumber := getLeetcodeNumber(files[left].Name())
		rightNumber := getLeetcodeNumber(files[right].Name())
		l, _ := strconv.Atoi(leftNumber)
		r, _ := strconv.Atoi(rightNumber)
		return l < r
	})
	count := 0
	for _, f := range files {
		item := modified(f, dir)
		if item != nil {
			count++
			b.WriteString(fmt.Sprintf("%s\n", *item))
		}
	}
	return b.String(), count
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
	show := fmt.Sprintf("|%s|[%s](https://github.com/jokerYellow/LeetCodeGo/tree/master/%s)|", number, questionName, finalName)
	return &show
}

func getLeetcodeNumber(name string) string {
	reg, err := regexp.Compile("^[0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.FindString(name)
}
