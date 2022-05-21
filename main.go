package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"github.com/sergitopereira/csvappender/helpers"
)

func dest_files(dest_path string) map[string]int {

	//return a map with files in destination path and value counter
	dst_files, err := ioutil.ReadDir(dest_path)
	if err != nil {
		log.Fatal(err)
	}
	m := make(map[string]int)
	for _, file := range dst_files {
		if file.IsDir() {
			continue
		}
		_, ok := m[file.Name()]
		if ok {
			m[file.Name()] += 1
		} else {
			m[file.Name()] = 1
		}

	}
	return m
}

func append_csv(test_file string, prod_file string) {
	body, err := ioutil.ReadFile(test_file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fmt.Println(string(body))

	f, err := os.OpenFile(prod_file, os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("\n")
	f.WriteString(string(body))
	return
}

func main() {
	args := os.Args
	helpers.TerminalHelper(args)
	test_files, err := ioutil.ReadDir(args[1])
	if err != nil {
		log.Fatal(err)
	}
	m := dest_files(args[2])

	for _, file := range test_files {
		if file.IsDir() {
			continue
		} else {

			test := regexp.MustCompile("_test_data(.csv)?$")
			res := test.ReplaceAllString(file.Name(), "")
			fmt.Println(res + ".csv")
			_, ok := m[res+".csv"]
			if ok {
				test_path := args[1] + "/" + file.Name()
				prod_path := args[2] + "/" + res+".csv"
				append_csv(test_path, prod_path )
			}
		}
	}
	fmt.Println("completed")
}
