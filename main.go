package main

import (
	"os/exec"
	"strings"
)

const SEPARATOR = "    "

func setBarStatus(elements []string) {
	status := strings.Join(elements, SEPARATOR)
	exec.Command("xsetroot", "-name", status).Run()
}

func main() {
	setBarStatus([]string{"test", "mytest"})
}
