package main

import (
	"fmt"
	"os/exec"
	"strings"
)
import "github.com/hekonsek/osexit"

func main() {
	var output, err = exec.Command("docker", "ps", "-a", "-n1").CombinedOutput()
	osexit.ExitOnError(err)
	var outputLines = strings.Split(string(output), "\n")
	var lastId = strings.Split(outputLines[1], " ")[0]
	fmt.Println(lastId)
}
