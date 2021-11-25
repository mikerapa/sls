package fileTree

import "testing"

func TestDirList_Sort(t *testing.T) {
	//build the test data
	var dl DirList
	fileNames := []string{"v.txt", "a.txt", "g"}
	dirNames := []string{"c", "1", "x"}
	for _, dirName := range dirNames{
		newDir := MakeNewDir(dirName)
		for _, fileName := range fileNames{
			newDir.Files.Add(fileName)
		}
		dl.Add(newDir)
	}

	//call the sort
	dl.Sort()

	sortedFileNames := []string{ "a.txt", "g", "v.txt"}
	sortedDirNames := []string{"1", "c", "x"}
	for di, dv := range dl{
		if dv.Path != sortedDirNames[di] {
			t.Errorf("Sorting error: got %s, expected %s", dv.Path, sortedDirNames[di])
		}

		for fi, fv := range dv.Files{
			if fv != sortedFileNames[fi]{
				t.Errorf("Sorting error: got %s, expected %s", fv, sortedFileNames[fi])
			}
		}
	}

}
