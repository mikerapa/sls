package fileTree

import (
	"errors"
	"fmt"
	"io/fs"
	"os/user"
	"path/filepath"
	"strings"
)

type Directory struct {
	Path string
	Files map[string]fs.DirEntry
	Dirs map[string]Directory
}

func MakeNewDir(dirPath string) Directory{
	newDir := Directory{Path: dirPath}
	newDir.Files = make(map[string]fs.DirEntry)
	newDir.Dirs = make(map[string]Directory)
	return newDir
}


func matchPatterns(fileName string, patternStrings ...string) bool {
	// handle a case in which there is no filter string
	if len(patternStrings)==1 && len(strings.Trim(patternStrings[0], " "))==0{
		return true
	}
	lowerFileName := strings.ToLower(fileName)

	for _, term:= range patternStrings{
		if len(term) == 0{ // skip empty filter strings
			continue
		}
		if strings.Contains(lowerFileName, strings.Trim(strings.ToLower(term), " ")) {
			return true
		}
	}
	return false
}

func RegularizePath(inputPath string) (outputPath string, err error) {
	outputPath = strings.Trim(inputPath, " ")

	// check for the tilde shortcut
	if strings.HasPrefix(outputPath, "~"){
		// the path provided contains a tilde which is a shortcut
		usr, err := user.Current()
		if err != nil {
			return "", errors.New("path provided contains a tilde but the user could not be resolved")
		}
		homeDir := usr.HomeDir
		// replace the tilde with the home directory
		outputPath = strings.Replace(outputPath, "~", homeDir, -1)
	}
	// clean the path before returning
	outputPath = filepath.Clean(outputPath)
	return
}

// TODO need a parameter to control the hidden file filter
// recursive function to read directory, apply filter criteria, and populate dir struct
func loadDir(fileSystem fs.FS, dirPath string, filterTerms []string) (dir Directory, fileCount int){
	dirEntries, err := fs.ReadDir(fileSystem, dirPath)
	if err != nil{
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}
	dir = MakeNewDir(dirPath)
	for _, entry := range dirEntries {
		p:=filepath.Join(dirPath, entry.Name())

		if entry.IsDir(){
			if !isHiddenFile(p){
				d,subFileCount := loadDir(fileSystem, p,filterTerms )
				// only add the dir to the results if there are files
				if subFileCount>0{
					dir.Dirs[p] = d
					fileCount = fileCount + subFileCount
				}

			}

		} else {
			if !isHiddenFile(p) && matchPatterns(entry.Name(), filterTerms... ) {
				dir.Files[entry.Name()] = entry
				fileCount ++
			}
		}
	}

	return
}

func GetFileTree(fileSystem fs.FS, rootPath string, filterPattern string) (dir Directory, fileCount int){
	// prepare the filter pattern here, because it should only be done once
	filterTerms := strings.Split(filterPattern, "*")

	dir, fileCount = loadDir(fileSystem, rootPath, filterTerms)
	return

}