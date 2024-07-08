// Package args contains functions used to processes arguments passed in terminal.
package args

import (
	"flag"
	"fmt"
	"strings"

	"ascii-art-color/errs"
)

var Usage = fmt.Errorf(`Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> "something"`)

// ProcessArgs processes arguments passed on terminal and returns a substr,
// str, filename, color as strings and a nil error if processed successfuly.
func ProcessArgs() (string, string, string, string, error) {
	c := flag.String("color", "", "--color=<your color>")
	if err := errs.ValidateFlag(); err != nil {
		return "", "", "", "", Usage
	}
	flag.Parse()
	var str, substr string
	fileName := "standard.txt"
	if flag.NFlag() == 0 { // no color flag
		switch flag.NArg() {
		case 1:
			str = flag.Arg(0)
		case 2:
			str = flag.Arg(0)
			fileName = flag.Arg(1) + ".txt"
		default:
			return "", "", "", "", Usage
		}
	} else if flag.NFlag() == 1 { // color flag
		switch flag.NArg() {
		case 1:
			str = flag.Arg(0)
		case 2:
			if errs.CheckFile(flag.Arg(1)) {
				fileName = flag.Arg(1) + ".txt"
				str = flag.Arg(0)
			} else {
				substr = flag.Arg(0)
				str = flag.Arg(1)
			}
		case 3:
			substr = flag.Arg(0)
			str = flag.Arg(1)
			fileName = flag.Arg(2) + ".txt"
		default:
			return "", "", "", "", Usage

		}
	} else {
		return "", "", "", "", Usage
	}
	str = strings.ReplaceAll(str, "\\t", "    ")
	str = strings.ReplaceAll(str, "\n", "\\n")
	substr = strings.ReplaceAll(substr, "\\t", "    ")
	substr = strings.ReplaceAll(substr, "\n", "\\n")
	return str, substr, fileName, *c, nil
}
