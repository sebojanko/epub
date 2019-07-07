package main

import (
	"epub"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	//list directory
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".epub" {
			continue
		}

		metadata := epub.GetMetadata(file.Name())
		printMetadata(metadata)
	}
}

func printMetadata(md epub.Metadata) {
	fmt.Println(md.Metadata.Title)
	fmt.Println(md.Metadata.ISBN)
	fmt.Println(md.Metadata.Description)
	fmt.Println(md.Metadata.Creator)
	fmt.Println(md.Metadata.Date)
	fmt.Println(md.Metadata.Publisher)
	fmt.Println(md.Metadata.Language)
	fmt.Println(md.Metadata.Format)
}
