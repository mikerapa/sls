package main

import (
	"reflect"
	"testing"
)

func TestCompareIn(t *testing.T) {
	tests := []struct {
		name string
		terms       []string
		inputString string
		want bool
	}{
		{name: "mixed case one term", terms:[]string{"got"}, inputString : "goT", want: true},
		{name: "mixed case 3 terms", terms:[]string{"got", "haD", "PUT"}, inputString : "put", want: true},
		{name: "mixed case 3 terms no match", terms:[]string{"got", "haD", "PUT"}, inputString : "flub", want: false},
		{name: "no terms", terms:[]string{}, inputString : "flub", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareIn(tt.terms, tt.inputString); got != tt.want {
				t.Errorf("CompareIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	tests := []struct {
		name        string
		inputText  string
		searchTerm string
		wantResults []Segment
	}{
		{name: "no matching terms", inputText: "sample text", searchTerm: "txt", wantResults: []Segment{{text: "sample text", highlight: false} }},
		{name: "one matching term", inputText: "sample text", searchTerm: "text", wantResults: []Segment{{text: "sample ", highlight: false},{text: "text", highlight: true} }},
		{name: "match the entire text", inputText: "sample text", searchTerm: "sample text", wantResults: []Segment{{text: "sample text", highlight: true}}},
		{name: "match the entire text", inputText: "I have some sample text to share", searchTerm: "sample text", wantResults: []Segment{{text: "I have some ", highlight: false},{text: "sample text", highlight: true}, {text: " to share", highlight: false}}},
		{name: "one matching term, different case", inputText: "sample texT", searchTerm: "TExt", wantResults: []Segment{{text: "sample ", highlight: false},{text: "texT", highlight: true} }},
		{name:"matching term with wild card", inputText: "I was born in a land downunder", searchTerm: "born*under",
			wantResults: []Segment{{text: "I was ", highlight: false},
				{text: "born", highlight: true},
				{text: " in a land down", highlight: false},
				{text: "under", highlight: true}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := parse(tt.inputText, tt.searchTerm); !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("parse() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}