package main

import (
	"fmt"
	"github.com/fatih/color"
	"regexp"
	"strings"
)

type Segment struct {
	text      string
	highlight bool
}

func main() {
	//b := regexp.MustCompile(`((?i)we)`).ReplaceAllString("We were all going West.", `%$1%`)
	//println(b)
	//
	//b2 := regexp.MustCompile("(?i)we").FindStringSubmatch("West East")
	//fmt.Printf("%q\n", b2)
	//
	//b3 := regexp.MustCompile("(?i)we").FindAllStringSubmatch("West East, we went out there for gold", -1)
	//fmt.Printf("%q\n", b3)

	//const SAMPLEINPUT = "We were there one day."
	//r := parse(SAMPLEINPUT, "we")
	//fmt.Println(SAMPLEINPUT, r)
	//printSegments(r)

	printHighlightText("I come FROM a land downunder", "down*own")
	printHighlightText("this was the file we were talking about", "thi*file")
}

func printHighlightText(text string, searchString string){
	printSegments(parse(text, searchString))
}

func printSegments(segments []Segment){
	for _, s:= range segments{
		if s.highlight {
			color.Set(color.FgBlack, color.BgHiGreen)
		} else {
			color.Unset()
		}
		fmt.Print(s.text)
	}
	fmt.Printf("\n")
}



func parse(inputText string, searchTerm string) (results []Segment) {
	const (
		SEARCHEXPRESSION = "((?i)%s)"
		PLACEHOLDER = "%%%"
	)

	terms := strings.Split(searchTerm, "*")
	// add markers around all found search strings
	markedString := inputText
	for _,term := range terms{
		searchRE := fmt.Sprintf(SEARCHEXPRESSION, term )
		markedString = regexp.MustCompile(searchRE).ReplaceAllString(markedString, fmt.Sprintf("%s$1%s", PLACEHOLDER, PLACEHOLDER) )
	}
	splits := strings.Split(markedString, PLACEHOLDER)
	for _,s:= range splits{
		if len(s)>0{
			// append any segments which are not empty
			termFound := CompareIn(terms, s)
			results = append(results, Segment{text: s, highlight: termFound})
		}
	}

	return
}


// Compare a given string to an array of terms. If the string is found in the array of terms, return true.
// The Comparison should ignore case.
func CompareIn(terms []string, inputString string) bool {
	for _, term:= range terms {
		if strings.EqualFold(term, inputString){ return true }
	}
	return false
}

