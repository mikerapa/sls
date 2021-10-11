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
		{name: "no filter", fileName: "file name.txt", patternStrings: []string{""}, want: true},
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

func TestGetFileTree(t *testing.T) {
	const (testDirPath = "../testfiles"
		moreTestDirPath = "../testfiles/moretestfiles"
	)

	tests := []struct {
		name string
		filterPatter string
		wantTestFilesCount int
		wantMoreTestFilesCount int
	}{
		{name: "all files", filterPatter: ".txt", wantTestFilesCount: 2, wantMoreTestFilesCount: 2},
		{name: "no files", filterPatter: ".hg", wantTestFilesCount: 0, wantMoreTestFilesCount: 0},
		{name: "one file in the sub-directory", filterPatter: "two", wantTestFilesCount: 0, wantMoreTestFilesCount: 1},
		{"one file in the main directory", "1", 1, 0},
	}

	for _,currentTest := range tests {
		t.Run(currentTest.name, func (t *testing.T){
			gotDirs := GetFileTree(testDirPath, currentTest.filterPatter)

			// make sure there is a sub-directory
			if len(gotDirs[testDirPath].Dirs)==1{
				t.Errorf("there should be 1 directory under this folder, got %d. ", len(gotDirs[testDirPath].Dirs))
			}

			// test the number of files in the main dir
			gotFilesCountTestDir := len(gotDirs[testDirPath].Files)
			if gotFilesCountTestDir!=currentTest.wantTestFilesCount{
				t.Errorf("there should be %d files in %s, got %d", currentTest.wantTestFilesCount, testDirPath, gotFilesCountTestDir)
			}

			// test the files in the more directory
			gotFilesCountMoreTestDir := len(gotDirs[moreTestDirPath].Files)
			if gotFilesCountMoreTestDir != currentTest.wantMoreTestFilesCount{
				t.Errorf("there should be %d files in %s, got %d", currentTest.wantMoreTestFilesCount, moreTestDirPath, gotFilesCountMoreTestDir)
			}

		})
	}

}