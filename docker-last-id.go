package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)
import "github.com/hekonsek/osexit"

func main() {
	offset := 1
	dockerized := len(os.Args) > 1 && os.Args[1] == "--dockerized"
	if dockerized {
		offset = 2
	}
	var output, err = exec.Command("docker", "ps", "-a", "-n2").CombinedOutput()
	osexit.ExitOnError(err)
	var outputLines = strings.Split(string(output), "\n")
	var lastId = strings.Split(outputLines[offset], " ")[0]
	fmt.Println(lastId)
}
