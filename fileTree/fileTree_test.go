package fileTree

import (
	"strings"
	"testing"
)

func Test_matchPatterns(t *testing.T) {

	tests := []struct {
		name string
		fileName       string
		patternStrings []string
		want bool
	}{
		{name: "simple match", fileName: "filename.go", patternStrings: []string{".go"}, want: true},
		{name: "case insensitive match", fileName: "filename.go", patternStrings: []string{"FILE"}, want: true},
		{name: "simple match, multiple terms", fileName: "filename.go", patternStrings: []string{"txt", ".go"}, want: true},
		{name: "no match", fileName: "filename.go", patternStrings: []string{".txt"}, want: false},
		{name: "extra space", fileName: " file name.txt", patternStrings: []string{" txt"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchPatterns(tt.fileName, tt.patternStrings...); got != tt.want {
				t.Errorf("matchPatterns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeNewDir(t *testing.T) {
	newDir := MakeNewDir("./temp")
	// make sure the path is correct
	if len(newDir.Path) ==0 || !strings.Contains(newDir.Path, "temp"){
		t.Errorf("MakeNewDir: path not set correctly. Got %s", newDir.Path)
	}

	// make sure the Files map has been made
	if newDir.Files == nil {
		t.Error("MakeNewDir: Files map was not initialized")
	}

	// make sure the Dirs map has been made
	if newDir.Dirs == nil {
		t.Error("MakeNewDir: Dirs map was not initialized")
	}
}