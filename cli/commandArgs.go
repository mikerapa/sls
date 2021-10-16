package cli

import (
	"flag"
	"os"
	"path/filepath"
)

func ProcessCommandLine() (showHelp bool, path string, filterPattern  string,command *flag.FlagSet, err error){
	showHelp, path, filterPattern, command, err = parseCommandLineArgs(os.Args[1:])
	path = filepath.Clean(path)
	return
}

func parseCommandLineArgs(args []string) (showHelp bool, path string, filterPattern  string, commandLine *flag.FlagSet, err error){
	commandLine = flag.NewFlagSet("sls", flag.ContinueOnError)

	commandLine.BoolVar(&showHelp, "help", false, "show the help text")
	commandLine.BoolVar(&showHelp, "h", false, "show the help text (shorthand)")
	commandLine.StringVar(&path, "path", ".", "directory path for search")
	commandLine.StringVar(&path, "p", ".", "directory path for search (shorthand)")
	commandLine.StringVar(&filterPattern, "filter", "", "text to use as a filter pattern. Use quotes if text contains the * wildcard")
	commandLine.StringVar(&filterPattern, "f", "", "text to use as a filter pattern. Use quotes if text contains the * wildcard")

	// parse the flag values
	err= commandLine.Parse(args)

	if err!= nil {
		// if there is an error in parsing the args, stop processing
		showHelp = true
		return
	}

	// process positional argument for filter pattern
	if filterPattern=="" && commandLine.NArg()>0{
		filterPattern = commandLine.Args()[0]
	}

	return
}
