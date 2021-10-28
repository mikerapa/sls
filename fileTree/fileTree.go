package fileTree

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
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

func GetFileTree(fileSystem fs.FS, rootPath string, filterPattern string) (dirs map[string]Directory, fileCount int){
	dirs = make(map[string]Directory)
	// prepare the filter pattern here, because it should only be done once
	filterTerms := strings.Split(filterPattern, "*")

	// walk the directory and files
	err := fs.WalkDir(fileSystem, rootPath,
		func(path string, dirEntry fs.DirEntry,  err2 error) error {
			if err2 != nil {
				// if there is an error here, it's likely because a directory is inaccessible. Print the error
				// but do not return it. If the error is returned, the WalkDir function will not process remaining
				// directories.
				fmt.Printf("ERROR: %s\n", err2.Error())
				return nil
			}
			if dirEntry.IsDir(){
				dirs[path] = MakeNewDir(path)
				//println("Adding a new directory " , path)
			} else {
				dirPath := filepath.Dir(path)
				d, ok := dirs[dirPath]
				if !ok {
					println("Error: directory is not in the map")
				}

				// if the file matches the filter pattern, add it to the dictionary
				if matchPatterns(dirEntry.Name(), filterTerms... ) {
					d.Files[dirEntry.Name()] = dirEntry
					fileCount++
				}

			}
			return nil
		})
	if err != nil {

		// Show any error. If the error came from WalkDir, it would not have been displayed previously.
		log.Println("ERROR: ", err.Error())
	}
	return
}