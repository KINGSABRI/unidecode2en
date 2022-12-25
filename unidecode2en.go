package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mozillazg/go-unidecode"
)

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %v", e)
		os.Exit(0)
	}
}

func main() {

	var (
		filePtr string
		strPtr  string
	)

	flag.StringVar(&filePtr, "file", "", "a file with unicode to decode into english")
	flag.StringVar(&filePtr, "f", "", "a file with unicode to decode into english")
	flag.StringVar(&strPtr, "string", "", "a string with unicode to decode into english")
	flag.StringVar(&strPtr, "s", "", "a string with unicode to decode into english")
	flag.Parse()

	if len(filePtr) == 0 && len(strPtr) == 0 {
		fmt.Println("Usage: ")
		fmt.Println("unidecode2en -file <FILE>")
		fmt.Println("unidecode2en -string <STRING>")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if filePtr != "" {
		file, err := os.ReadFile(filePtr)
		check(err)
		fmt.Print(unidecode.Unidecode(string(file)))
	}

	if strPtr != "" {
		fmt.Print(unidecode.Unidecode(strPtr))
	}

}
