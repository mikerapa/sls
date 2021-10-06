package main

import (
	"fmt"
	"github.com/fatih/color"
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

	println(markStrings("sample text", []string{"sample text"}))
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
		// set the highlight when evaluating the first char in the string
		if position == 0 {
			highlightOn = char == '*'
		}

		if (highlightOn && char !='*') || (!highlightOn && char=='*'){
			// end the existing segment
			segments = append(segments, Segment{text: originalString[segmentStartPosition:position], highlight:highlightOn})
			highlightOn = !highlightOn
			segmentStartPosition = position
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



func parse(inputText string, searchTerm string) (segments []Segment) {
	// prepare a list of search terms
	terms := strings.Split(searchTerm, "*")
	// add markers around all found search strings
	markedString := markStrings(inputText, terms)
	segments = convertMarkedStringToSegments(inputText, markedString)
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

