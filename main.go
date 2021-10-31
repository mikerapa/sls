package main

import (
	"fmt"
	"os"
	"sls/cli"
	"sls/fileTree"
)

func main() {
	showHelp, path, filterPattern, showHidden,  command, err := cli.ProcessCommandLine()

	if err!=nil {
		// if there is an error, show the error message to the user
		fmt.Printf("Command line error: %s", err.Error())
	}

	if showHelp{
		// show the help text and exit the application
		cli.ShowHelp(command)
		os.Exit(0)
	}

	// fix the path
	path, err = fileTree.RegularizePath(path)
	if err != nil {
		fmt.Printf("path error: %s, path = %s", err.Error(), path)
		os.Exit(0)
	}

	//fmt.Printf("path: %s, pattern: %s\n", path, filterPattern)

	// get the file tree and show  the results
	tree, fileCount := fileTree.GetFileTree(os.DirFS(path), ".", filterPattern, showHidden)
	cli.PrintDirectory(tree, filterPattern)
	cli.PrintFileCount(fileCount)
}




