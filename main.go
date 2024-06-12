package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"ascii-art-fs/ascii"
)

func main() {
	lenArgs := len(os.Args)
	if lenArgs < 2 || lenArgs > 3 {
		fmt.Print(
			"Usage: go run . [STRING] [BANNER]\n\n" +

				"EX: go run . something standard\n")
		return
	}
	// Store the first argument in str and replace tab and newline characters
	str := os.Args[1]
	str = strings.ReplaceAll(str, "\\t", "    ")
	str = strings.ReplaceAll(str, "\n", "\\n")
	err := ascii.IsPrintableAscii(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set the second argument for the banner file name. the default has been set to standard.txt
	fileName := "standard.txt"
	if lenArgs == 3 {
		fileName = os.Args[2] + ".txt"
	}
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
	count := 0
	for _, str := range words {
		if str == "" {
			count++
			if count < len(words) {
				fmt.Println()
			}
		} else {
			// Print the ASCII representation of the word
			ascii.PrintAscii(str, contentSlice, 0)
		}
	}
}
