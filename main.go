package main

import (
	"fmt"
	"github.com/sergitopereira/csvappender/helpers"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func find_files(path string) map[string]int {

	//return a map with files in destination path and value counter
	dst_files, err := ioutil.ReadDir(path)
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
	f, err := os.OpenFile(prod_file, os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("\n")
	f.WriteString(string(body))
}

func main() {
	args := os.Args
	helpers.TerminalHelper(args)
    // start logger
	log_file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(log_file)
	log.Printf("Start csvappender v1.0")
	test_files, err := ioutil.ReadDir(args[1])
	if err != nil {
		log.Fatal(err)
	}
	m := find_files(args[2])
	log.Println("-> "+args[2]+ " has the following files:"  )
	log.Println(m)
	log.Println("-> "+args[1]+ " has the following files:"  )
	log.Println(find_files(args[1]))
	for _, file := range test_files {
		if file.IsDir() {
			continue
		} else {
			test := regexp.MustCompile("_test_data(.csv)?$")
			res := test.ReplaceAllString(file.Name(), "")
			//fmt.Println(res + ".csv")
			_, ok := m[res+".csv"]
			if ok {
				test_path := args[1] + "/" + file.Name()
				prod_path := args[2] + "/" + res + ".csv"
				append_csv(test_path, prod_path)
				log.Println("->" + prod_path + " has been appended with contents of " + test_path)
			}
		}
	}
	log.Println("End csvappender")
	fmt.Println("Files were successfully appended. Review logs.txt")
}
