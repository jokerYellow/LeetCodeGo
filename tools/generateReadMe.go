package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	dir := "./"
	files, error := ioutil.ReadDir(dir)
	if error != nil {
		log.Fatal(error)
	}
	for _, f := range files {
		modified(f, dir)
	}
}

//change direcrory name componented by number and description
func modified(f os.FileInfo, dir string) {
	name := f.Name()
	number := isLeetcodeQuestion(name)
	if len(number) == 0 {
		return
	}

	files, error := ioutil.ReadDir(dir + "/" + name)
	if error != nil {
		log.Fatal(error)
	}
	if len(files) == 0 {
		return
	}
	questionName := strings.Replace(files[0].Name(), ".go", "", 1)
	finalName := number + "-" + questionName
	fmt.Println(finalName)
	os.Rename(dir+"/"+name, dir+"/"+finalName)
}

func isLeetcodeQuestion(name string) string {
	reg, err := regexp.Compile("^[0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.FindString(name)
}
