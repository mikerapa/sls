package cli

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"sls/fileTree"
	"sls/segments"
)

func ShowHelp(command *flag.FlagSet){
	// show the help text
	fmt.Println("sls [OPTION] [FILTER PATTERN]")
	command.PrintDefaults()
}

func PrintFileCount(fileCount int){
	//fmt.Printf("\t")
	color.Set(color.BgYellow, color.FgBlack)
	fmt.Printf("\t%d files found\t", fileCount)
	color.Unset()
	fmt.Printf("\n")

}

func PrintDirectory(directory fileTree.Directory, filterString string){
	// only print the directory name if there are files under it
	if len(directory.Files) > 0 {
		color.Set(color.FgYellow)
		println(directory.Path)
		color.Unset()
		for _, fv := range directory.Files{
			fmt.Printf("\t")
			printHighlightText(fv.Name(), filterString)
		}

	}

	for _,dv:= range directory.Dirs{
		PrintDirectory(dv, filterString)
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
	color.Unset()
	fmt.Printf("\n")
}
