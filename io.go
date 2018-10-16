package handy

import (
	"log"
	"os"
	"path/filepath"
)

// OpenFile function opens given path file.
// if a file doesn't exist, it will be created.
func OpenFile(path string, perm os.FileMode) (f *os.File, err error) {
	// make directories at first
	d := filepath.Dir(path)
	if err := os.MkdirAll(d, 0755); err != nil {
		return nil, err
	}
	f, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR, perm)
	if err != nil {
		return nil, err
	}
	return f, err
}

// OpenFileAppend function opens given path file with append mode.
// if a file doesn't exist, it will created.
func OpenFileAppend(path string, perm os.FileMode) (f *os.File, err error) {
	// make directories at first
	d := filepath.Dir(path)
	if err := os.MkdirAll(d, 0755); err != nil {
		return nil, err
	}
	f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return nil, err
	}
	return f, err
}

// PushStringToFile function pushes given string to given path file
func PushStringToFile(s, path string) error {
	f, err := OpenFileAppend(path, 0600)
	if err != nil {
		log.Println("can't open file: ", path)
		return err
	}
	defer f.Close()
	_, err = f.WriteString(s)
	if err != nil {
		log.Println("failed to write string to ", path)
		return err
	}
	return nil
}
