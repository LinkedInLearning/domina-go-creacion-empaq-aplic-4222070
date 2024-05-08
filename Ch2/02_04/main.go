package main

import "github.com/linkedinlearning/domina-go/goreleaser/cmd"

var author string
var version string
var date string
var commit string

func main() {
	cmd.Execute(cmd.Metadata{Author: author, Version: version, Date: date, Commit: commit})
}
