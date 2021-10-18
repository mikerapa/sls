package cli

import (
	"testing"
)

func TestParseCommandLineArgs(t *testing.T) {
	tests := []struct {
		name              string
		args []string
		wantShowHelp      bool
		wantPath          string
		wantFilterPattern string
		wantErr           bool
	}{
		{"no inputs", []string{}, false, ".",  "", false},
		{"just a path", []string{"-p", "testfolder"}, false, "testfolder", "", false},
		{"path and positional filter2", []string{"--path", "testfolder", ".txt"}, false, "testfolder", ".txt", false},
		{"path and positional filter", []string{"-p", "testfolder", ".txt"}, false, "testfolder", ".txt", false},
		{"path and filter", []string{"-f", ".txt", "-p", "testfolder"}, false, "testfolder", ".txt", false},
		{"help", []string{"-h"}, true, ".", "", false},
		{"help 2", []string{"--help"}, true, ".", "", false},
		{"invalid input", []string{"--fake flag "}, true, ".", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShowHelp, gotPath, gotFilterPattern, gotCommand, err := parseCommandLineArgs(tt.args)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseCommandLineArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotShowHelp != tt.wantShowHelp {
				t.Errorf("parseCommandLineArgs() gotShowHelp = %v, want %v", gotShowHelp, tt.wantShowHelp)
			}
			if gotPath != tt.wantPath {
				t.Errorf("parseCommandLineArgs() gotPath = %v, want %v", gotPath, tt.wantPath)
			}
			if gotFilterPattern != tt.wantFilterPattern {
				t.Errorf("parseCommandLineArgs() gotFilterPattern = %v, want %v", gotFilterPattern, tt.wantFilterPattern)
			}
			if gotCommand==nil{
				t.Errorf("the command (type FlagSet) is nil")
			}
		})
	}
}
