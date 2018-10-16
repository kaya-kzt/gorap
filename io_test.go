package handy

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenFile(t *testing.T) {
	p := "./dirtest/test.txt"
	f, err := OpenFile(p, 0700)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	_, err = f.WriteString("test string\n")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	_, err = os.Stat(p) // file exist?
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	f.Close()

	// remove file&dir
	d := filepath.Dir(p)
	err = os.RemoveAll(d)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
