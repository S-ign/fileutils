package fileutils

import (
	"fmt"
	"io/ioutil"

	"github.com/S-ign/stringutils"
)

// FiletoByte converts a file to a byte string
func FiletoByte(filename string) ([]byte, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// WriteFileWithHashedName writes file name with a hash of byte array
func WriteFileWithHashedName(b []byte, dir string, extension string) (string, error) {

	hash, err := stringutils.GenerateHash(b)
	if err != nil {
		return "", err
	}

	// create tmp file info
	tmpFilename := fmt.Sprintf("%s.%s", hash, extension)
	fullpath := fmt.Sprintf("%s/%s", dir, tmpFilename)

	return fullpath, nil
}
