package main

import (
	"io/ioutil"
	"path/filepath"
)

func main() {
	//list directory
	files, err := ioutil.ReadDir("./")
	check(err)
	writer := initCSVWriter("books.csv")

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) != ".epub" {
			continue
		}
		reader := openFile(file.Name())

		var container Container
		var metadataFile string
		var metadata Metadata

		for _, f := range reader.File {
			if f.Name == "META-INF/container.xml" {
				container = parseContainerXML(f)
				metadataFile = container.Rootfiles[0].FullPath
				break
			}
		}

		for _, f := range reader.File {
			if f.Name == metadataFile {
				metadata = parseMetadata(f)
				fields := getMetadata(metadata)
				writeLineToCSV(writer, fields)
				break
			}
		}
		reader.Close()
	}
}

func getMetadata(md Metadata) []string {
	fields := []string{md.Metadata.Title,
		md.Metadata.ISBN,
		md.Metadata.Description,
		md.Metadata.Creator,
		md.Metadata.Date,
		md.Metadata.Publisher,
		md.Metadata.Language,
		md.Metadata.Format}
	return fields
}
