package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	testCases := []struct {
		color  string
		substr string
		str    string
		banner string
		want   string
	}{
		{
			color:  "",
			substr: "",
			str:    "",
			banner: "standard",
			want:   "",
		},
		{
			color:  "red",
			substr: "",
			str:    "Hello",
			banner: "standard",
			want: "\x1b[31m _    _  \x1b[0m\x1b[31m       \x1b[0m\x1b[31m _  \x1b[0m\x1b[31m _  \x1b[0m\x1b[31m        \x1b[0m\n" +
				"\x1b[31m| |  | | \x1b[0m\x1b[31m       \x1b[0m\x1b[31m| | \x1b[0m\x1b[31m| | \x1b[0m\x1b[31m        \x1b[0m\n" +
				"\x1b[31m| |__| | \x1b[0m\x1b[31m  ___  \x1b[0m\x1b[31m| | \x1b[0m\x1b[31m| | \x1b[0m\x1b[31m  ___   \x1b[0m\n" +
				"\x1b[31m|  __  | \x1b[0m\x1b[31m / _ \\ \x1b[0m\x1b[31m| | \x1b[0m\x1b[31m| | \x1b[0m\x1b[31m / _ \\  \x1b[0m\n" +
				"\x1b[31m| |  | | \x1b[0m\x1b[31m|  __/ \x1b[0m\x1b[31m| | \x1b[0m\x1b[31m| | \x1b[0m\x1b[31m| (_) | \x1b[0m\n" +
				"\x1b[31m|_|  |_| \x1b[0m\x1b[31m \\___| \x1b[0m\x1b[31m|_| \x1b[0m\x1b[31m|_| \x1b[0m\x1b[31m \\___/  \x1b[0m\n" +
				"\x1b[31m         \x1b[0m\x1b[31m       \x1b[0m\x1b[31m    \x1b[0m\x1b[31m    \x1b[0m\x1b[31m        \x1b[0m\n" +
				"\x1b[31m         \x1b[0m\x1b[31m       \x1b[0m\x1b[31m    \x1b[0m\x1b[31m    \x1b[0m\x1b[31m        \x1b[0m\n",
		},
		{
			color:  "yellow",
			substr: "there",
			str:    "hello there",
			banner: "shadow",
			want: "                                       \x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\n" +
				"_|                _| _|                \x1b[33m  _|     \x1b[0m\x1b[33m_|       \x1b[0m\x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\n" +
				"_|_|_|     _|_|   _| _|   _|_|         \x1b[33m_|_|_|_| \x1b[0m\x1b[33m_|_|_|   \x1b[0m\x1b[33m  _|_|   \x1b[0m\x1b[33m_|  _|_| \x1b[0m\x1b[33m  _|_|   \x1b[0m\n" +
				"_|    _| _|_|_|_| _| _| _|    _|       \x1b[33m  _|     \x1b[0m\x1b[33m_|    _| \x1b[0m\x1b[33m_|_|_|_| \x1b[0m\x1b[33m_|_|     \x1b[0m\x1b[33m_|_|_|_| \x1b[0m\n" +
				"_|    _| _|       _| _| _|    _|       \x1b[33m  _|     \x1b[0m\x1b[33m_|    _| \x1b[0m\x1b[33m_|       \x1b[0m\x1b[33m_|       \x1b[0m\x1b[33m_|       \x1b[0m\n" +
				"_|    _|   _|_|_| _| _|   _|_|         \x1b[33m    _|_| \x1b[0m\x1b[33m_|    _| \x1b[0m\x1b[33m  _|_|_| \x1b[0m\x1b[33m_|       \x1b[0m\x1b[33m  _|_|_| \x1b[0m\n" +
				"                                       \x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\n" +
				"                                       \x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\x1b[33m         \x1b[0m\n",
		},
		{
			color:  "rainbow",
			substr: "not",
			str:    "not a valid color",
			banner: "shadow",
			want:   "invalid color/rgb/hex\n",
		},
		{
			color:  "red",
			substr: "yet",
			str:    "substr not in str",
			banner: "shadow",
			want:   "sub-string is not contained in the string\n",
		},
		{
			color:  "",
			substr: "",
			str:    "here",
			banner: "",
			want: ` _                           
| |                          
| |__     ___   _ __    ___  
|  _ \   / _ \ | '__|  / _ \ 
| | | | |  __/ | |    |  __/ 
|_| |_|  \___| |_|     \___| 
                             
                             
`,
		},
		{
			color:  "#766",
			substr: "Hello",
			str:    "Hello",
			banner: "standard",
			want: "\x1b[38;2;119;102;102m _    _  \x1b[0m\x1b[38;2;119;102;102m       \x1b[0m\x1b[38;2;119;102;102m _  \x1b[0m\x1b[38;2;119;102;102m _  \x1b[0m\x1b[38;2;119;102;102m        \x1b[0m\n" +
				"\x1b[38;2;119;102;102m| |  | | \x1b[0m\x1b[38;2;119;102;102m       \x1b[0m\x1b[38;2;119;102;102m| | \x1b[0m\x1b[38;2;119;102;102m| | \x1b[0m\x1b[38;2;119;102;102m        \x1b[0m\n" +
				"\x1b[38;2;119;102;102m| |__| | \x1b[0m\x1b[38;2;119;102;102m  ___  \x1b[0m\x1b[38;2;119;102;102m| | \x1b[0m\x1b[38;2;119;102;102m| | \x1b[0m\x1b[38;2;119;102;102m  ___   \x1b[0m\n" +
				"\x1b[38;2;119;102;102m|  __  | \x1b[0m\x1b[38;2;119;102;102m / _ \\ \x1b[0m\x1b[38;2;119;102;102m| | \x1b[0m\x1b[38;2;119;102;102m| | \x1b[0m\x1b[38;2;119;102;102m / _ \\  \x1b[0m\n" +
				"\x1b[38;2;119;102;102m| |  | | \x1b[0m\x1b[38;2;119;102;102m|  __/ \x1b[0m\x1b[38;2;119;102;102m| | \x1b[0m\x1b[38;2;119;102;102m| | \x1b[0m\x1b[38;2;119;102;102m| (_) | \x1b[0m\n" +
				"\x1b[38;2;119;102;102m|_|  |_| \x1b[0m\x1b[38;2;119;102;102m \\___| \x1b[0m\x1b[38;2;119;102;102m|_| \x1b[0m\x1b[38;2;119;102;102m|_| \x1b[0m\x1b[38;2;119;102;102m \\___/  \x1b[0m\n" +
				"\x1b[38;2;119;102;102m         \x1b[0m\x1b[38;2;119;102;102m       \x1b[0m\x1b[38;2;119;102;102m    \x1b[0m\x1b[38;2;119;102;102m    \x1b[0m\x1b[38;2;119;102;102m        \x1b[0m\n" +
				"\x1b[38;2;119;102;102m         \x1b[0m\x1b[38;2;119;102;102m       \x1b[0m\x1b[38;2;119;102;102m    \x1b[0m\x1b[38;2;119;102;102m    \x1b[0m\x1b[38;2;119;102;102m        \x1b[0m\n",
		},
		{
			color:  "yellow",
			substr: "invalid",
			str:    "invalid banner",
			banner: "shadw",
			want:   "open ./banner/shadw.txt: no such file or directory\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.str, func(t *testing.T) {
			var cmd *exec.Cmd
			if tc.color == "" {
				cmd = exec.Command("sh", "-c", fmt.Sprintf("go run . \"%s\"", tc.str))
			} else {
				if tc.substr == "" {
					cmd = exec.Command("sh", "-c", fmt.Sprintf("go run . --color=%s \"%s\" \"%s\"", tc.color, tc.str, tc.banner))
				} else {
					cmd = exec.Command("sh", "-c", fmt.Sprintf("go run . --color=%s \"%s\" \"%s\" \"%s\"", tc.color, tc.substr, tc.str, tc.banner))
				}
			}
			output, err := cmd.Output()
			if err != nil {
				t.Fatalf("Error running program: %v", err)
			}
			got := string(output)
			if got != tc.want {
				t.Errorf("\ngot:\n%v\nwant:\n%v\n", got, tc.want)
			}
		})
	}
}
