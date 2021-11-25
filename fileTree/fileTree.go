package fileTree

import (
	"errors"
	"fmt"
	"io/fs"
	"os/user"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
)



func MakeNewDir(dirPath string) Directory{
	newDir := Directory{Path: dirPath}
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

// recursive function to read directory, apply filter criteria, and populate dir struct
func loadDir(fileSystem fs.FS, dirPath string, filterTerms []string, showHidden bool, dirChan chan Directory, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Printf("inside loadDir for %s\n", dirPath)
	dirEntries, err := fs.ReadDir(fileSystem, dirPath)
	if err != nil{
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}
	dir := MakeNewDir(dirPath)
	for _, entry := range dirEntries {
		p:=filepath.Join(dirPath, entry.Name())

		if entry.IsDir(){
			if showHidden || !isHiddenFile(p){
				wg.Add(1)
				//fmt.Printf("calling loadDir for %s\n", p)
				go loadDir(fileSystem, p,filterTerms, showHidden, dirChan, wg )
			}

		} else {
			if (showHidden || !isHiddenFile(p)) && matchPatterns(entry.Name(), filterTerms... ) {
				//fmt.Printf("adding file %s\n", p)
				dir.Files.Add(entry.Name())
			}
		}
	}

	// if this dir has files, send it back to the channel
	if len(dir.Files)>0{
		//fmt.Printf("sending %s to dirChan\n", dir.Path)
		wg.Add(1)
		dirChan <- dir
	}
	//fmt.Printf("done with loadDir for %s\n", dirPath)
}



func GetFileTree(fileSystem fs.FS, rootPath string, filterPattern string, showHidden bool) (dirs DirList, fileCount int){
	// prepare the filter pattern here, because it should only be done once
	filterTerms := strings.Split(filterPattern, "*")
	dirChan := make (chan Directory, 50)
	//dirs = make(map[string]Directory)
	wg := sync.WaitGroup{}
	var fileCountInt64 int64 = 0
	go func(){
		for newDir := range dirChan{
			//fmt.Printf("1- got a dir in the dirChan %s, new file count %d, total file cound %d\n", newDir.Path, len(newDir.Files), fileCountInt64)
			// TODO the atomic call below may not be needed.
			dirs.Add(newDir)
			var newFilesCount int64 = int64(len(newDir.Files))
			atomic.AddInt64(&fileCountInt64, newFilesCount)
			//fmt.Printf("2- got a dir in the dirChan %s, new file count %d, total file cound %d\n", newDir.Path, len(newDir.Files), fileCountInt64)
			wg.Done()
		}

	}()
	wg.Add(1)
	go loadDir(fileSystem, rootPath, filterTerms, showHidden, dirChan, &wg)


	wg.Wait()
	close(dirChan)
	// sort the directories by path
	dirs.Sort()
	//fmt.Printf("closing the channel\n")
	fileCount = int(fileCountInt64)
	return

}