package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version = ""
	commit  = ""
	date    = ""
	builtBy = ""
)

var (
	cmdname = filepath.Base(os.Args[0])
)

func main() {
	// args := kingpin.MustParse(app.Parse(os.Args[1:]))
	kingpin.MustParse(app.Parse(os.Args[1:]))

	// Handle updating to a new version
	fmt.Print("Attempting update of " + cmdname + "...")
	update_result, err := doSelfUpdate()
	if err != nil {
		fmt.Println("Couldn't update at this time. Please try again later. Continuing...")
	}
	if update_result {
		fmt.Println("Please run " + cmdname + " again.")
		os.Exit(0)
	}

	if *cmd_commit {
		fmt.Println("version:", version)
		fmt.Println("Commit:", commit)
		fmt.Println("Date:", date)
		fmt.Println("built by:", builtBy)
	}

	// main loop
	// switch args {
	// case example.FullCommand():
	// 	fmt.Println("Yay!")
	// }
	fmt.Println("Pretend this does something cool.")
}
