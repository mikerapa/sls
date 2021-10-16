package main

import (
	"fmt"
	"os"
	"sls/cli"
	"sls/fileTree"
)

func main() {

	showHelp, path, filterPattern, command, err := cli.ProcessCommandLine()

	if err!=nil {
		// if there is an error, show the error message to the user
		fmt.Printf("Command line error: %s", err.Error())
	}

	if showHelp{
		// show the help text and exit the application
		cli.ShowHelp(command)
		os.Exit(0)
	}

	fmt.Printf("path: %s, pattern: %s\n", path, filterPattern)

	// get the file tree and show  the results
	tree := fileTree.GetFileTree(os.DirFS("."), path, filterPattern)
	for _, tv := range tree{
		cli.PrintDirectory(tv, filterPattern)
	}

}




