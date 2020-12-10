package util

import (
	"io/ioutil"
	"testing"
)

func TestWriteByteFile_ByteContent(t *testing.T) {
	// ioutil.ReadFile("")
	ioutil.WriteFile("newfile", []byte("my data"), 0755)
}
