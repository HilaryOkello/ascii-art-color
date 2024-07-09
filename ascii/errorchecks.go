
// this package has functions that handles errors at different stages
package ascii

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// the function takes string and returns an error and returns
// a non-nil error if it finds non-printable characters
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
		// Check if the next character is an escape letter
		NextIsAnEscapeLetter := strings.ContainsAny(string(next), escapes)
		// Check if the current character is an escape character
		isAnEscape := (char == '\\' && NextIsAnEscapeLetter)
		// Check if the current character is non-printable
		isNonPrintable := ((char < ' ' || char > '~') && char != '\n')

		if isAnEscape {
			foundEscapes += "\\" + string(next)
		}
		if isNonPrintable {
			nonPrintables += string(char)
		}
	}
	// Construct error message if escape characters or non-printable characters are found
	if foundEscapes != "" {
		return fmt.Errorf("%s%s", foundEscapes, errMessage)
	} else if nonPrintables != "" {
		return fmt.Errorf("%s%s", nonPrintables, errMessage)
	}

	return nil
}

// function takes the file name and content of file name respectively and then returns an error message
// if the file has been tampered
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

func CheckFile(s string) bool {
	files := []string{"standard", "shadow", "thinkertoy"}
	for _, file := range files {
		if file == s {
			return true
		}
	}
	return false
}

func ValidateFlag() error {
	usage := fmt.Errorf(`Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> "something"`)
	seenFlags := make(map[string]bool)
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			if strings.HasPrefix(arg, "-color") {
				return usage
			} else if !strings.Contains(arg, "=") && strings.Contains(arg, "color") {
				return usage
			}
			flagName := strings.SplitN(arg[2:], "=", 2)[0]
			if seenFlags[flagName] {
				return usage
			}
			seenFlags[flagName] = true
		}
	}
	return nil
}
