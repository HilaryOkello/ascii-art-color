// Package ascii provides functions for printing ASCII art
// with optional color highlighting.
package ascii

import (
	"fmt"
	"strings"
)

// PrintArgs contains parameters for the PrintAscii function.
type PrintArgs struct {
	Str        string
	Substr     string
	Color      string
	Characters []string
}

// PrintAscii prints ASCII art based on the given PrintArgs configuration.
func PrintAscii(args *PrintArgs) {
	colorCode := "\033[" + args.Color + "m"
	reset := "\033[0m"
	index := 0

	// Loop through each line of ASCII art (up to 8 lines)
	for index < 8 {
		indices := GetIndices(args.Str, args.Substr)
		track := 0
		count := 0

		for i, char := range args.Str {
			character := args.Characters[int(char)-32]
			lines := strings.Split(character, "\n")
			if shouldPrintWithColor(args, indices, track, i) {
				fmt.Printf("%s%s%s", colorCode, lines[index], reset)
				count++
				if count == len(args.Substr) && track < len(indices)-1 {
					track++
					count = 0
				}
			} else {
				fmt.Print(lines[index])
			}
		}
		fmt.Println()
		index++
	}
}

// shouldPrintWithColor determines if a character should be printed
// with color highlighting based on the provided arguments.
func shouldPrintWithColor(args *PrintArgs, indices []int, track, i int) bool {
	if args.Color != "" && args.Substr == "" {
		return true
	} else if args.Substr != "" && len(indices) > 0 {
		if i >= indices[track] && i < indices[track]+len(args.Substr) {
			return true
		}
	}
	return false
}

// GetIndices returns the starting indices of all occurrences of substr in str.
// If substr is an empty string, it returns an empty slice.
func GetIndices(str, substr string) (indices []int) {
	if substr == "" {
		return
	}
	start := 0
	for {
		index := strings.Index(str[start:], substr)
		if index == -1 {
			break
		}
		indices = append(indices, start+index)
		start += index + len(substr)
	}
	return
}
