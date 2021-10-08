package fileTree

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Directory struct {
	Path string
	Files map[string]os.FileInfo
	Dirs map[string]Directory
}

func MakeNewDir(dirPath string) Directory{
	newDir := Directory{Path: dirPath}
	newDir.Files = make(map[string]os.FileInfo)
	newDir.Dirs = make(map[string]Directory)
	return newDir
}


func matchPatterns(fileName string, patternStrings ...string) (bool) {
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

func GetFileTree(topPath string, searchPatter string) (dirs map[string]Directory){
	dirs = make(map[string]Directory)
	// prepare the filter pattern here, because it should only be done once
	filterTerms := strings.Split(searchPatter, "*")

	err := filepath.Walk(topPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//fmt.Println(path, info.Size())
			if info.IsDir(){
				dirs[path] = MakeNewDir(path)
				//println("Adding a new directory " , path)
			} else {
				dirPath := filepath.Dir(path)
				d, ok := dirs[dirPath]
				if !ok {
					println("Error: directory is not in the map")
				}
				//println("Adding a new file " , dirPath, info.Name())
				if matchPatterns(info.Name(), filterTerms... ) {
					d.Files[info.Name()] = info
				}


			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return
}