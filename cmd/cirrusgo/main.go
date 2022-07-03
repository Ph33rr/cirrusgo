package main

import "github.com/Ph33rr/CirrusGo/internal/runner"

func main() {
	options := runner.ParseOptions()
	runner.New(options)
}
