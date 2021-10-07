package fileTree

import (
	"log"
	"os"
	"path/filepath"
)

type Directory struct {
	Path string
	Files map[string]os.FileInfo
	Dirs map[string]Directory
}

func GetFileTree(topPath string, searchPatter string) (dirs map[string]Directory){
	dirs = make(map[string]Directory)
	err := filepath.Walk(topPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//fmt.Println(path, info.Size())
			if info.IsDir(){
				// TODO this should be extracted into a function which creates a Dir
				newDir := Directory{Path: path}
				newDir.Files = make(map[string]os.FileInfo)
				newDir.Dirs = make(map[string]Directory)
				dirs[path] = newDir
				//println("Adding a new directory " , path)
			} else {
				dirPath := filepath.Dir(path)
				d, ok := dirs[dirPath]
				if !ok {
					println("Error: directory is not in the map")
				}
				//println("Adding a new file " , dirPath, info.Name())
				d.Files[info.Name()] = info


			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return
}