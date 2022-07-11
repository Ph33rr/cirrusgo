package main

import "github.com/Ph33rr/cirrusgo/internal/runner"

func main() {
	options := runner.ParseOptions()
	runner.New(options)
}
