// ascii has functions that handles errors at different stages of the program.
package ascii

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// expectedChecksum contains the hash of the contents of our banner files
// calculated using the SHA256 algorithm.
var expectedChecksum = map[string]string{
	"standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	"thinkertoy.txt": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
	"shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
}

// usage is the correct usage of the
var usage = fmt.Errorf(`Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> "something"`)

// IsPrintableAscii takes a string and returns a non-nil error
// if it finds non-printable characters.
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

// CheckFileTamper takes the file name and content of file name
// and returns an error if the file has been tampered.
func CheckFileTamper(file string, content []byte) error {
	checksum := sha256.Sum256(content)
	computedChecksum := hex.EncodeToString(checksum[:])
	if computedChecksum != expectedChecksum[file] {
		return fmt.Errorf("%s tampered", file)
	}
	return nil
}

// CheckFile returns true if the string passed is either of our filenames:
// standard, shadow, or thinkerty.
func CheckFile(s string) bool {
	files := []string{"standard", "shadow", "thinkertoy"}
	for _, file := range files {
		if file == s {
			return true
		}
	}
	return false
}

// ValidateFlag returns an error if the flag passed is
// not in the format "--color=<value>".
func ValidateFlag() error {
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
