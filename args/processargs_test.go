// Package args contains functions used to processes arguments passed in terminal.
package args

import (
	"flag"
	"os"
	"testing"
)

func TestProcessArgs(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		ExpStr      string
		ExpSubstr   string
		ExpFilename string
		ExpCoolor   string
		wantErr     bool
	}{
		{
			name:        "valid Arguments - All",
			args:        []string{"cmd", "--color=blue", "World", "Hello World", "shadow"},
			ExpStr:      "Hello World",
			ExpSubstr:   "World",
			ExpFilename: "shadow.txt",
			ExpCoolor:   "blue",
			wantErr:     false,
		},
		{
			name:        "valid Arguments Without substr",
			args:        []string{"cmd", "--color=blue", "Hello World", "shadow"},
			ExpStr:      "Hello World",
			ExpSubstr:   "",
			ExpFilename: "shadow.txt",
			ExpCoolor:   "blue",
			wantErr:     false,
		},
		{
			name:        "Invalid flag",
			args:        []string{"cmd", "-color=blue", "Hello World", "shadow"},
			ExpStr:      "",
			ExpSubstr:   "",
			ExpFilename: "",
			ExpCoolor:   "",
			wantErr:     true,
		},
		{
			name:        "not enough arguments",
			args:        []string{"cmd"},
			ExpStr:      "",
			ExpSubstr:   "",
			ExpFilename: "",
			ExpCoolor:   "",
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestWithArgs(t, tt.args, tt.ExpStr, tt.ExpSubstr, tt.ExpFilename, tt.ExpCoolor, tt.wantErr)
		})
	}
}

func runTestWithArgs(t *testing.T, args []string, ExpStr, ExpSubstr, ExpFilename, ExpCoolor string, wantErr bool) {
	// Save the original flags and restore them after the test
	originalArgs := os.Args
	// Resetting flag.CommandLine for clean parsing
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // Reset flag.CommandLine
	defer func() {
		os.Args = originalArgs
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	}()

	// Set os.Args to the test case's arguments
	os.Args = args
	got, got1, got2, got3, err := ProcessArgs()
	if (err != nil) != wantErr {
		t.Errorf("ProcessArgs() error = %v, wantErr %v", err, wantErr)
		return
	}
	if got != ExpStr {
		t.Errorf("ProcessArgs() got = %v, want %v", got, ExpStr)
	}
	if got1 != ExpSubstr {
		t.Errorf("ProcessArgs() got1 = %v, want %v", got1, ExpSubstr)
	}
	if got2 != ExpFilename {
		t.Errorf("ProcessArgs() got2 = %v, want %v", got2, ExpFilename)
	}
	if got3 != ExpCoolor {
		t.Errorf("ProcessArgs() got3 = %v, want %v", got3, ExpCoolor)
	}
}
