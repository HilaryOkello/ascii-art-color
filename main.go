package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"ascii-art-color/ascii"
)

var usage = fmt.Errorf(`Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> "something"`)

func main() {
	var str string
	var substr string
	fileName := "standard.txt"

	c := flag.String("color", "", "--color=<your color>")
	if err := ascii.ValidateFlag(); err != nil {
		fmt.Println(usage)
		return
	}
	flag.Parse()

	if flag.NFlag() == 0 {
		switch flag.NArg() {
		case 1:
			str = flag.Arg(0)
		case 2:
			str = flag.Arg(0)
			fileName = flag.Arg(1) + ".txt"
		default:
			fmt.Println(usage)
			return

		}
	} else if flag.NFlag() == 1 {
		switch flag.NArg() {
		case 1:
			str = flag.Arg(0)
		case 2:
			if ascii.CheckFile(flag.Arg(1)) {
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
			fmt.Println(usage)
			return

		}
	} else {
		fmt.Print(usage)
	}
	str = strings.ReplaceAll(str, "\\t", "    ")
	str = strings.ReplaceAll(str, "\n", "\\n")
	err := ascii.IsPrintableAscii(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set the second argument for the banner file name. the default has been set to standard.txt
	filePath := os.DirFS("./banner")
	contentByte, err := fs.ReadFile(filePath, fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(contentByte) == 0 {
		fmt.Println("Banner file is empty")
		return
	}
	// check if the banner file has been tampered with
	er := ascii.CheckFileTamper(fileName, contentByte)
	if er != nil {
		fmt.Println(er)
		return
	}
	contentString := string(contentByte[1:])
	if fileName == "thinkertoy.txt" {
		// convert all carriage returns to newlines
		contentString = strings.ReplaceAll(string(contentByte[2:]), "\r\n", "\n")
	}
	contentSlice := strings.Split(contentString, "\n\n")
	// Split the input string by "\\n" to get individual words
	words := strings.Split(str, "\\n")
	substrs := strings.Split(substr, "\\n")
	count := 0
	for i, s := range words {
		if s == "" {
			count++
			if count < len(words) {
				fmt.Println()
			}
		} else {
			if strings.Contains(str, substr) {
				ascii.PrintAscii(s, substrs[i], *c, contentSlice, 0)
			} else {
				ascii.PrintAscii(s, substr, *c, contentSlice, 0)
			}
		}
	}
}
