package epub

import (
	"archive/zip"
	"encoding/xml"
	"io/ioutil"
	"log"
)

// Public function to retrieve metadata from epub file.
// Params: string - epub filename
// Returns: Metadata - struct containing metadata
func GetMetadata(filename string) Metadata {
	reader := openFile(filename)
	defer reader.Close()

	metadataFile := findMetadataFile(reader)
	return getMetadataFromContainerFile(reader, metadataFile)
}

// Check for all the files inside the epub to find the container.xml file which contains
// the path to the metadata file.
// Params: *zip.ReadCloser - reader object
// Returns: string - file path to metadata file
func findMetadataFile(reader *zip.ReadCloser) string {
	filePath := ""

	for _, f := range reader.File {
		if f.Name == "META-INF/container.xml" {
			container := parseContainerXML(f)
			filePath = metadataFilePath(container)
			break
		}
	}
	return filePath
}

// Extracts the full path to the metadata file from container.xml
// Params: Container - struct with XML data from container.xml
// Returns: string - file path to the metadata file
func metadataFilePath(container Container) string {
	return container.Rootfiles[0].FullPath
}

// Finds the metadata file and fills the Metadata struct with data
func getMetadataFromContainerFile(reader *zip.ReadCloser, metadataFile string) Metadata {
	var metadata Metadata
	for _, f := range reader.File {
		if f.Name == metadataFile {
			metadata = parseMetadata(f)
			break
		}
	}
	return metadata
}

// Parses a container.xml file which is a file in META-INF/ that contains the path to content.opf.
// content.opf contains the .epub file meta data. Returns the filled Container struct.
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

// Check's if the error is not nil
// Params: error
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
