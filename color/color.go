// Package color provides functions for converting color strings
// (named colors, RGB, hexadecimal) to ANSI color codes.
package color

import (
	"fmt"
	"strconv"
	"strings"
)

// ansiCodes maps common color names to their corresponding ANSI color codes.
// The ANSI color codes are used to set text color in terminal output.
var ansiCodes = map[string]string{
	"black":          "30",
	"red":            "31",
	"green":          "32",
	"yellow":         "33",
	"blue":           "34",
	"magenta":        "35",
	"cyan":           "36",
	"white":          "37",
	"gray":           "90",
	"bright red":     "91",
	"bright green":   "92",
	"bright yellow":  "93",
	"bright blue":    "94",
	"bright magenta": "95",
	"bright cyan":    "96",
	"bright white":   "97",
	"orange":         "38;2;255;165;0",
	"brown":          "38;2;165;42;42",
	"purple":         "38;2;128;0;128",
	"pink":           "38;2;255;192;203",
	"olive":          "38;2;128;128;0",
	"teal":           "38;2;0;128;128",
	"navy":           "38;2;0;0;128",
}

// ParseColor converts a color string to its corresponding ANSI color code.
// The input color string can be in one of the following formats:
// - Named color (e.g., "red", "blue", "orange")
// - RGB format (e.g., "rgb(255, 0, 0)")
// - Hex format (e.g., "#FF0000")
// If the color string is valid, the function returns the ANSI color code
// in the format "38;2;r;g;b" or a predefined ANSI code for named colors.
// If the color string is invalid, it returns an error.
func ParseColor(c string) (string, error) {
	if strings.HasPrefix(c, "rgb(") {
		code, err := rgbToANSI(c)
		if err != nil {
			return "", err
		}
		return code, nil
	} else if strings.HasPrefix(c, "#") {
		code, err := hexToANSI(c)
		if err != nil {
			return "", err
		}
		return code, nil
	} else if ansiCodes[c] == "" && c != "" {
		return "", fmt.Errorf("invalid color/rgb/hex")
	} else {
		return ansiCodes[c], nil
	}
}

// rgbToANSI converts an RGB color string in the format "rgb(r, g, b)"
// to its corresponding ANSI color code in the format "38;2;r;g;b".
// If the input string is not a valid RGB format, it returns an error.
func rgbToANSI(c string) (string, error) {
	trimmed := strings.TrimPrefix(strings.TrimSuffix(c, ")"), "rgb(")
	parts := strings.Split(trimmed, ",")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid rgb")
	}
	r, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return "", err
	}
	g, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return "", err
	}
	b, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		return "", err
	}
	if (r < 0 || r > 255) || (g < 0 || g > 255) || (b < 0 || b > 255) {
		return "", fmt.Errorf("invalid rgb")
	}
	return fmt.Sprintf("38;2;%d;%d;%d", r, g, b), nil
}

// hexToANSI converts a hexadecimal color string in the format "#RRGGBB"
// to its corresponding ANSI color code in the format "38;2;r;g;b".
// If the input string is not a valid hexadecimal color, it returns an error.
func hexToANSI(c string) (string, error) {
	hex := strings.TrimPrefix(c, "#")
	if len(hex) == 3 {
		hex = fmt.Sprintf("%c%c%c%c%c%c", hex[0], hex[0], hex[1], hex[1], hex[2], hex[2])
	}
	if len(hex) != 6 {
		return "", fmt.Errorf("invalid color hexadecimal")
	}
	r, err := strconv.ParseInt(hex[0:2], 16, 64)
	if err != nil {
		return "", err
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 64)
	if err != nil {
		return "", err
	}
	b, err := strconv.ParseInt(hex[4:], 16, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("38;2;%d;%d;%d", r, g, b), nil
}
