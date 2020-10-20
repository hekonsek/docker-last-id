package main

import (
	"fmt"
	"os/exec"
	"strings"
)
import "github.com/hekonsek/osexit"

func main() {
	var output, err = exec.Command("docker", "ps", "-a", "-n500",
		"--format", "table {{.ID}} {{.Image}} {{.Command}}").CombinedOutput()
	osexit.ExitOnError(err)
	var outputLines = strings.Split(string(output), "\n")[1:]
	linesWithoutAlipine := []string{}
	for _, line := range outputLines[1 : len(outputLines)-1] {
		lineParts := strings.Split(line, " ")
		if lineParts[1] != "alpine" && lineParts[2] != `""echo""` {
			linesWithoutAlipine = append(linesWithoutAlipine, lineParts[0])
		}
	}
	fmt.Println(linesWithoutAlipine[0])
}
