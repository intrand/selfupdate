package main

import "gopkg.in/alecthomas/kingpin.v2"

var (
	app = kingpin.New(cmdname, "updates itself").Version(version).Author("intrand")
	// example = app.Command("example", "some example cmd")
	cmd_commit = app.Flag("commit", "shows commit and exits").Default("false").Bool()
)
