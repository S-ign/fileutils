package fileutils

import "io/ioutil"

// FiletoByte converts a file to a byte string
func FiletoByte(filename string) ([]byte, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}
