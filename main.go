package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
	"sls/fileTree"
	"sls/segments"
)

func main() {

	var (
		showHelp bool
		path string
		filterPattern string
	)
	flag.BoolVar(&showHelp, "help", false, "show the help text")
	flag.BoolVar(&showHelp, "h", false, "show the help text (shorthand)")
	flag.StringVar(&path, "path", ".", "directory path for search")
	flag.StringVar(&path, "p", ".", "directory path for search (shorthand)")
	flag.StringVar(&filterPattern, "filter", "", "text to use as a filter pattern. Use quotes if text contains the * wildcard")
	flag.StringVar(&filterPattern, "f", "", "text to use as a filter pattern. Use quotes if text contains the * wildcard")


	flag.Parse()
	if showHelp{
		fmt.Println("sls [OPTION] [FILTER PATTERN]")
		flag.PrintDefaults()
	}

	//fmt.Printf("unused args: %v\n", flag.NArg())
	//for i,a:= range flag.Args(){
	//	fmt.Printf("argument %d: %s\n", i, a)
	//}

	// process positional argument for filter pattern
	if filterPattern=="" && flag.NArg()>0{
		filterPattern = flag.Args()[0]
	}
	//fmt.Printf("path: %s, pattern: %s\n", path, filterPattern)

	// show  the results
	tree := fileTree.GetFileTree(os.DirFS("."), path, filterPattern)
	for _, tv := range tree{
		printDirectory(tv, filterPattern)
	}


}


func printDirectory(directory fileTree.Directory, filterString string){
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
		printDirectory(dv, filterString)
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



