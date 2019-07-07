package main

import (
	"archive/zip"
	"encoding/csv"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

// Parses a container.xml file which is a file in META-INF/ that contains the path to content.opf.
// content.opf contains the meta data on the .epub. Returns the filled Container struct.
// Params: *zip.File - pointer to epub file object
// Returns: Container
func parseContainerXML(f *zip.File) Container {
	byteValue := readFileContent(f)

	return unmarshalContainerXML(byteValue)
}

// Fills a struct with XML data
// Params: []byte - XML bytes data
// Returns: Container - struct with XML data filled
func unmarshalContainerXML(byteValue []byte) Container {
	var container Container
	xml.Unmarshal(byteValue, &container)
	return container
}

// Parses the content.opf file which contains the metadata and returns the filled Metadata struct.
// Params: *zip.File - pointer to epub file object
// Returns: Metadata
func parseMetadata(f *zip.File) Metadata {
	content := readFileContent(f)

	return unmarshalMetaDataXML(content)
}

// Fills a struct with XML data
// Params: []byte - XML bytes data
// Returns: Metadata - struct with XML data filled
func unmarshalMetaDataXML(byteValue []byte) Metadata {
	var meta Metadata
	xml.Unmarshal(byteValue, &meta)
	return meta
}

// Open the epub file for reading. epub files are archives which can be opened with the 'zip' package.
// Params: string - file name
// Returns: *zip.ReadCloser - reader object
func openFile(name string) *zip.ReadCloser {
	reader, err := zip.OpenReader(name)
	check(err)

	return reader
}

// Reads a file in the epub and returns the content.
// Params: *zip.File - pointer to file object in epub archive
// Returns: []byte - file content in bytes
func readFileContent(f *zip.File) []byte {
	reader, err := f.Open()
	check(err)
	defer reader.Close()

	byteValue, err := ioutil.ReadAll(reader)
	check(err)

	return byteValue
}

// Creates the CSV file and initializes the CSV Writer
// Params: string - name of the file
// Returns: *csv.Writer
func initCSVWriter(filename string) *csv.Writer {
	f, err := os.Create(filename)
	check(err)
	w := csv.NewWriter(f)

	return w
}

// Writes an array of strings to a CSV file, comma delimited
// Params: *csv.Writer - pointer to object which writes to the file
// 		   []string - array of strings to write
// Returns: error
func writeLineToCSV(w *csv.Writer, fields []string) error {
	err := w.Write(fields)
	defer w.Flush()
	return err
}

// Check's if the error is not nil
// Params: error
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
