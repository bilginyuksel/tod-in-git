package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/bilginyuksel/executor"
)

func main() {
	var err error
	var out []byte
	if runtime.GOOS == "windows" {
		out, err = exec.Command("powershell", "ls").Output()
	} else {
		out, err = exec.Command("ls").Output()
	}
	executor.Hello()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", out)
	}
}
