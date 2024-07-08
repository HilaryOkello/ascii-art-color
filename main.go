// Package main is the entry point for the application.
// It prints the ascii representation based on arguments passed.
package main

import (
	"fmt"
	"strings"

	"ascii-art-color/args"
	"ascii-art-color/ascii"
	"ascii-art-color/banner"
	"ascii-art-color/color"
	"ascii-art-color/errs"
)

func main() {
	str, substr, filename, c, err := args.ProcessArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Check that str & substr are printable characters
	if err := errs.IsPrintableAscii(str); err != nil {
		fmt.Println(err)
		return
	}
	if err := errs.IsPrintableAscii(substr); err != nil {
		fmt.Println(err)
		return
	}
	color, err := color.ParseColor(strings.ToLower(c))
	if err != nil {
		fmt.Println(err)
		return
	}

	contentSlice, err := banner.ReadBannerFile(strings.ToLower(filename))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Split the str & substr by "\\n" to get the string section in each line
	strs := strings.Split(str, "\\n")
	substrs := strings.Split(substr, "\\n")
	count := 0 // tracks empty strings after splitting with \n

	for i, s := range strs {
		if s == "" {
			count++
			if count < len(strs) {
				fmt.Println()
			}
		} else {
			if substr != "" && i < len(substrs) && strings.Contains(str, substr) {
				substr = substrs[i]
			}
			args := &ascii.PrintArgs{
				Str:        s,
				Substr:     substr,
				Color:      color,
				Characters: contentSlice,
			}
			ascii.PrintAscii(args)
		}
	}
}
