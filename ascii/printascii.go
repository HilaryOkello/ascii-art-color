package ascii

import (
	"fmt"
	"strings"
)

/*
function takes str which is string passed at argument one ,contentslice which is characters from filaname,
and index which tracks the number of lines per character.
it recursively print the provided string up to last line
*/
func PrintAscii(str string, contentSlice []string, index int) {
	if index == 8 {
		return
	}
	// loop through each character in a str and prints it line by line.
	for _, char := range str {
		character := contentSlice[int(char)-32]                 // obtain char from contentslice
		character = strings.ReplaceAll(character, "\r\n", "\n") // thinkertoy
		lines := strings.Split(character, "\n")
		fmt.Print(lines[index])
	}
	fmt.Println()
	PrintAscii(str, contentSlice, index+1)
}
