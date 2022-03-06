package fileutils

import (
	"fmt"
	"strings"
)

// fileNameCollection is a slice of base filenames ex.
// []string{"data1.pdf", "data2.pdf"}
// []string{"image1.png", "image2.png"}
// each collection must be of the same file extension
// suggestion: name varaibes of this type the name of file extension ex.
// pdfs := fileNameCollection {
//   Files: []string{"data1.pdf", "data2.pdf"}
// }
type fileNameCollection struct {
	// name of the collection, used to create key for FileNameCollections
	name string

	// a slice of filenames collected from a given path
	filenames []string

	// length of slice
	length int

	// The directory in which the filenames will be collected
	dir string

	// used to search for filenames in a directory
	subStrings []string
}

// FileNameCollection a slice of filenames with the same type
// extension
func FileNameCollection(name string, files []string, dir string, subStrings []string) (fileNameCollection, error) {

	// returns error if files slice is empty
	if len(files) == 0 {
		return fileNameCollection{}, fmt.Errorf(
			"FileNameCollectionSetter: can not set empty slice, files: %s", files)
	}
	if len(dir) == 0 {
		return fileNameCollection{}, fmt.Errorf(
			"FileNameCollectionSetter: blank dir: %s", dir)
	}

	length := len(files)

	// checks collection for consistent file extensions
	extension := strings.Split(files[0], ".")[1]
	for _, f := range files {
		if extension != strings.Split(f, ".")[1] {
			return fileNameCollection{}, fmt.Errorf(
				"error: collection must be of the same file extension, files %s", files)
		}
		if len(strings.Split(f, ".")) < 2 {
			return fileNameCollection{}, fmt.Errorf(
				"error: filenames must contain file extension, files %s", files)
		}
	}

	return fileNameCollection{
		name:       name,
		filenames:  files,
		length:     length,
		dir:        dir,
		subStrings: subStrings,
	}, nil
}

// ChangeFilenamesExtension creates a duplicate of this fileNameCollection but changes
// f.filenames extension
func (f fileNameCollection) ChangeFilenamesExtension(
	extension string) (fileNameCollection, error) {

	filenames := []string{}

	for _, filename := range f.filenames {
		baseFilename := strings.Split(filename, ".")[0]
		filenames = append(filenames, baseFilename+"."+extension)
	}

	nf, err := FileNameCollection(extension, filenames, f.dir, f.subStrings)
	if err != nil {
		return fileNameCollection{},
			fmt.Errorf("ChangeFilenamesExtension: unable to create type fileNameCollection: %s", err)
	}

	return nf, nil
}

// fileNameCollections a map of fileNameCollections, allowing for multiple sets of filenames, each collection may have a different extensiosn
type fileNameCollections struct {
	keys        []string
	values      []fileNameCollection
	collections map[string]fileNameCollection
}

// FileNameCollections creates a map of fileNameCollection and sets all
// other attributes
func FileNameCollections(fnc ...fileNameCollection) (fileNameCollections,
	error) {

	// constructs map of ...fileNameCollection
	collectionsMap := map[string]fileNameCollection{}
	for _, f := range fnc {
		collectionsMap[f.name] = f
	}

	// extracts key and value of each fileNameCollection
	keys := []string{}
	values := []fileNameCollection{}
	for k, v := range collectionsMap {
		keys = append(keys, k)
		values = append(values, v)
	}

	// ensures that empty fileNameCollection is not added
	for _, v := range values {
		if len(v.filenames) == 0 {
			return fileNameCollections{}, fmt.Errorf(
				"FileNameCollectionsSetter: can not assign empty fileNameCollection.filenames: %s",
				v.filenames)
		}
	}

	return fileNameCollections{
		keys:        keys,
		values:      values,
		collections: collectionsMap,
	}, nil
}
