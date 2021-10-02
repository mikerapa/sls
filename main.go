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


	fmt.Println(markStrings("I come from a land down under.", []string{"own", "FRO"}))
	fmt.Println(markStrings("I come from a land down under.", []string{"om", "FRO"}))
	originalText := "I come from a land downunder"
	markedString := markStrings(originalText, []string{"om", "FRO"})
	segments := convertMarkedStringToSegments(originalText, markedString)
	printSegments(segments)

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

func convertMarkedStringToSegments(originalString string, markedString string) (segments []Segment){
	highlightOn := false
	segmentStartPosition := 0
	for position, char := range markedString{
		if highlightOn {
			if char != '*'{
				segments = append(segments, Segment{text: originalString[segmentStartPosition:position], highlight:true})
				highlightOn = false
				segmentStartPosition = position
			}
		} else {
			if char == '*'{
				segments = append(segments, Segment{text: originalString[segmentStartPosition:position], highlight:false})
				highlightOn = true
				segmentStartPosition = position
			}
		}
	}
	// get the last segment
	segments = append(segments, Segment{text: originalString[segmentStartPosition:], highlight:highlightOn})
	return
}

func markStrings(originalText string, matchStrings []string) (markedString string) {
	lowerText := strings.ToLower(originalText)
	markedString = originalText
	for _, matchString:= range matchStrings{
		currentPosition :=0
		for currentPosition < len(lowerText){
			foundIndex := strings.Index(lowerText[currentPosition:], strings.ToLower(matchString))
			// If the substring is not found, get out of the loop
			if foundIndex ==-1 {
				break
			}
			foundIndex = foundIndex + currentPosition
			markedString = markedString[:foundIndex] + strings.Repeat("*", len(matchString)) + markedString[foundIndex+len(matchString):]
			currentPosition = foundIndex + len(matchString)

		}
	}

	return
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

