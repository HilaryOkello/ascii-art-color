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

//process args passed on terminal and returns our substr, str, filename, and error
// returns nil error if processed successfuly
func processArgs() (string, string, string, error) {
	var str, substr string
	fileName := "standard.txt" 
	if flag.NFlag() == 0 {//color flag not passed
		switch flag.NArg() {
		case 1: // just single string
			str = flag.Arg(0)
		case 2: // string + banner
			str = flag.Arg(0)
			fileName = flag.Arg(1) + ".txt"
		default:
			return "", "", "", fmt.Errorf("%s", usage)
		}
	} else if flag.NFlag() == 1 {//color flag passed
		switch flag.NArg() {
		case 1: //single string
			str = flag.Arg(0)
		case 2: // str + banner || substr + str
			if ascii.CheckFile(flag.Arg(1)) {
				fileName = flag.Arg(1) + ".txt"
				str = flag.Arg(0)
			} else {
				substr = flag.Arg(0)
				str = flag.Arg(1)
			}
		case 3: //substr + str + banner
			substr = flag.Arg(0)
			str = flag.Arg(1)
			fileName = flag.Arg(2) + ".txt"
		default:
			return "", "", "", fmt.Errorf("%s", usage)

		}
	} else {
		return "", "", "", fmt.Errorf("%s", usage)
	}
	//replace newline && tab characters
	str = strings.ReplaceAll(str, "\\t", "    ")
	str = strings.ReplaceAll(str, "\n", "\\n")
	substr = strings.ReplaceAll(substr, "\\t", "    ")
	substr = strings.ReplaceAll(substr, "\n", "\\n")
	return str, substr, fileName, nil
}


// reads from the banner file and return ascii-art characters as a []string
// returns non-nil error if it encounters errors
func ReadBannerFile(fileName string) ([]string, error) {
	// Set the second argument for the banner file name. the default has been set to standard.txt
	filePath := os.DirFS("./banner")
	contentByte, err := fs.ReadFile(filePath, fileName)
	if err != nil {
		return nil, err
	}
	if len(contentByte) == 0 {
		return nil, fmt.Errorf("Banner file is empty")
	}
	// check if the banner file has been tampered with
	if err = ascii.CheckFileTamper(fileName, contentByte); err != nil {
		return nil, err
	}
	contentString := string(contentByte[1:])
	if fileName == "thinkertoy.txt" {
		// convert all carriage returns to newlines
		contentString = strings.ReplaceAll(string(contentByte[2:]), "\r\n", "\n")
	}
	contentSlice := strings.Split(contentString, "\n\n")
	return contentSlice, nil
}

func main() {
	// define a color flag
	c := flag.String("color", "", "--color=<your color>")
	if err := ascii.ValidateFlag(); err != nil {
		fmt.Println(usage)
		return
	}
	flag.Parse()
	// process args passed
	str, substr, fileName, err := processArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	// check if characters are printable
	if err := ascii.IsPrintableAscii(str); err != nil {
		fmt.Println(err)
		return
	}
	// get characters from the banner file
	contentSlice, err := ReadBannerFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Split the str & substr by "\\n" to get the string section in each line
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
			//handles where substr has multiple lines
			if substr != "" && i < len(substrs) && strings.Contains(str, substr) {
				ascii.PrintAscii(s, substrs[i], *c, contentSlice, 0)
			} else {
				ascii.PrintAscii(s, substr, *c, contentSlice, 0)
			}
		}
	}
}
