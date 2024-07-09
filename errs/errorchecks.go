// Package errs provides error handling utilities and validation functions.
package errs

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// IsPrintableAscii checks if the given string contains only printable ASCII characters.
// It returns an error indicating non-printable characters or escape sequences found.
func IsPrintableAscii(str string) error {
	var nonPrintables string
	var foundEscapes string
	errMessage := ": Not within the printable ascii range"
	for index, char := range str {

		escapes := "avrfb"
		var next byte
		if index < len(str)-1 {
			next = str[index+1]
		}
		NextIsAnEscapeLetter := strings.ContainsAny(string(next), escapes)
		isAnEscape := (char == '\\' && NextIsAnEscapeLetter)
		isNonPrintable := ((char < ' ' || char > '~') && char != '\n')

		if isAnEscape {
			foundEscapes += "\\" + string(next)
		}
		if isNonPrintable {
			nonPrintables += string(char)
		}
	}
	if foundEscapes != "" {
		return fmt.Errorf("%s%s", foundEscapes, errMessage)
	} else if nonPrintables != "" {
		return fmt.Errorf("%s%s", nonPrintables, errMessage)
	}

	return nil
}

// CheckFileTamper verifies if the file content checksum matches the expected checksum for a file.
// It returns an error if the checksum does not match, indicating possible tampering.
func CheckFileTamper(file string, content []byte) error {
	expectedChecksum := map[string]string{
		"standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
		"thinkertoy.txt": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
		"shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	}
	checksum := sha256.Sum256(content)
	computedChecksum := hex.EncodeToString(checksum[:])
	if computedChecksum != expectedChecksum[file] {
		return fmt.Errorf("%s tampered", file)
	}
	return nil
}

// CheckFile checks if the given filename corresponds to know banner file.
// It returns true if the string matches one of the files and false otherwise.
func CheckFile(s string) bool {
	files := []string{"standard", "shadow", "thinkertoy"}
	for _, file := range files {
		if file == strings.ToLower(s) {
			return true
		}
	}
	return false
}

// ValidateFlag validates command-line flags passed to the program.
// It returns an error if the flags are not used according to the expected usage.
func ValidateFlag() error {
	usage := fmt.Errorf(`Usage: go run . [OPTION] [STRING]

	EX: go run . --color=<color> <substring to be colored> "something"`)
	seenFlags := make(map[string]bool) // tracks duplication
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			if strings.HasPrefix(arg, "-color") {
				return usage
			} else if strings.Contains(arg, "color") && !strings.Contains(arg, "=") {
				return usage
			} else if strings.HasPrefix(arg, "--color=") {
				flagName := "--color="
				if seenFlags[flagName] {
					return usage
				}
				seenFlags[flagName] = true
			}
		}
	}
	return nil
}
