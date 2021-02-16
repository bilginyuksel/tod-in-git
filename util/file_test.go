package util

import (
	"testing"
    "os"
    "log"
)

func TestWriteByteFile_ByteContent(t *testing.T) {
	// ioutil.ReadFile("")

    f, err := os.OpenFile("newfile", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        log.Println(err)
    }

    defer f.Close()
    _, err = f.Write([]byte("my byte string asdsa"))
    if err != nil {
        log.Println(err)
    }
}
