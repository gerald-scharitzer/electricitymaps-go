package main

import (
	"os"
)

const Usage = `Usage: go run package [-]

"package" is the main package of this module.

"-" processes standard input.`

func main() {
	switch len(os.Args) {
	case 1:
	case 2:
		println("TODO stdin")
	default:
		println(Usage)
	}
}
