package ascii

import (
	"fmt"
	"strconv"
	"strings"
)

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
}

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
	} else {
		return ansiCodes[c], nil
	}
}

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

func hexToANSI(c string) (string, error) {
	hex := strings.TrimPrefix(c, "#")
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
