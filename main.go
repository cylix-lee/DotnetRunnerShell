package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func errorChecked[T any](result T, err error) T {
	if err != nil {
		log.Fatalln(err.Error())
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: DotnetRunnerShell <path_to_entry.cs>")
	}

	entry := errorChecked(filepath.Abs(os.Args[1]))
	command := exec.Command("dotnet", "run", "--project", filepath.Dir(entry))
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
}
