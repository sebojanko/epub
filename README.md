# epub

E-books files are mostly epubs. Since epubs are essentially archives, we can read files inside and extract metadata to help us better organize our library.

When opening the .epub in an archive manager, we can see a folder named META-INF which contains `container.xml`. Inside this file we can find the location of the metadata file.
The metadata file will usually have the extension `.opf`, but it's also an XML file.

## Metadata

I've found most books have these fields available. Some fields may sometimes be empty, depends on the file creator.
```
type Metadata struct {
    ISBN        string `xml:"identifier"`
    Title       string `xml:"title"`
    Description string `xml:"description"`
    Creator     string `xml:"creator"`
    Date        string `xml:"date"`
    Publisher   string `xml:"publisher"`
    Language    string `xml:"language"`
    Format      string `xml:"format"`
}
```

## Example

Available in [examples/example.go](examples/example.go)
