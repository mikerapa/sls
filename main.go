package main

import (
	"fmt"
	"github.com/fatih/color"
	"sls/fileTree"
	"sls/segments"
)

func main() {
	const (initialPath = "/home/mike/go/src"
	pattern = "*.go")
	tree := fileTree.GetFileTree(initialPath, pattern)
	for _, tv := range tree{
		printDirectory(tv)
	}

	printHighlightText("I come FROM a land downunder", "down*own")
	printHighlightText("this was the file we were talking about", "thi*file")

	//println(markStrings("sample text", []string{"sample text"}))
}

func printDirectory(directory fileTree.Directory){
	println("Folder: ", directory.Path)
	for _, fv := range directory.Files{
		println("File: ", fv.Name())
	}

	for _,dv:= range directory.Dirs{
		printDirectory(dv)
	}

}

func printHighlightText(text string, searchString string){
	printSegments(segments.Parse(text, searchString))
}

func printSegments(segments []segments.Segment){
	for _, s:= range segments{
		if s.Highlight {
			color.Set(color.FgBlack, color.BgHiGreen)
		} else {
			color.Unset()
		}
		fmt.Print(s.Text)
	}
	fmt.Printf("\n")
}

//
//// Compare a given string to an array of terms. If the string is found in the array of terms, return true.
//// The Comparison should ignore case.
//func CompareIn(terms []string, inputString string) bool {
//	for _, term:= range terms {
//		if strings.EqualFold(term, inputString){ return true }
//	}
//	return false
//}

