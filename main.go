package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {
	var err error
	var out []byte
	if runtime.GOOS == "windows" {
		out, err = exec.Command("powershell", "echo \"\nHello world\n\"").Output()
	} else {
		out, err = exec.Command("ls").Output()
	}
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", out)
	}
}
