package main

import (
	"github.com/pajicf/stanalizator/cmd"
)

func main() {
	rc := cmd.NewRootCommand()

	rc.Execute()
}
