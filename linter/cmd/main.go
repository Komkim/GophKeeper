package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
	"linter/internal"
)

func main() {
	checks := internal.NewAnalyzer()

	multichecker.Main(
		checks...,
	)
}
