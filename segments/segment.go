package segments

import "strings"

type Segment struct {
	Text      string
	Highlight bool
}

func Parse(inputText string, searchTerm string) (segments []Segment) {
	// prepare a list of search terms
	terms := strings.Split(searchTerm, "*")
	// add markers around all found search strings
	markedString := markStrings(inputText, terms)
	segments = convertMarkedStringToSegments(inputText, markedString)
	return
}


func markStrings(originalText string, matchStrings []string) (markedString string) {
	lowerText := strings.ToLower(originalText)
	markedString = originalText
	for _, matchString:= range matchStrings{
		if len(matchString) ==0 {
			continue
		}
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

func convertMarkedStringToSegments(originalString string, markedString string) (segments []Segment){
	highlightOn := false
	segmentStartPosition := 0
	for position, char := range markedString{
		// set the Highlight when evaluating the first char in the string
		if position == 0 {
			highlightOn = char == '*'
		}

		if (highlightOn && char !='*') || (!highlightOn && char=='*'){
			// end the existing segment
			segments = append(segments, Segment{Text: originalString[segmentStartPosition:position], Highlight:highlightOn})
			highlightOn = !highlightOn
			segmentStartPosition = position
		}
	}
	// get the last segment
	segments = append(segments, Segment{Text: originalString[segmentStartPosition:], Highlight:highlightOn})
	return
}
