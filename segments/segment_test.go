package segments

import (
	"reflect"
	"testing"
)

func BenchmarkParse(b *testing.B) {
	for i:=0;i<b.N;i++{
		_ = Parse("fakefilenamestring.txt", "file*ake")
	}
}

func Test_parse(t *testing.T) {
	tests := []struct {
		name        string
		inputText  string
		searchTerm string
		wantResults []Segment
	}{
		{name: "no matching terms", inputText: "sample Text", searchTerm: "txt", wantResults: []Segment{{Text: "sample Text", Highlight: false} }},
		{name: "one matching term", inputText: "sample Text", searchTerm: "Text", wantResults: []Segment{{Text: "sample ", Highlight: false},{Text: "Text", Highlight: true} }},
		{name: "match the entire Text", inputText: "sample Text", searchTerm: "sample Text", wantResults: []Segment{{Text: "sample Text", Highlight: true}}},
		{name: "partial match in the middle of the Text", inputText: "I have some sample Text to share", searchTerm: "sample Text", wantResults: []Segment{{Text: "I have some ", Highlight: false},{Text: "sample Text", Highlight: true}, {Text: " to share", Highlight: false}}},
		{name: "one matching term, different case", inputText: "sample texT", searchTerm: "TExt", wantResults: []Segment{{Text: "sample ", Highlight: false},{Text: "texT", Highlight: true} }},
		{name:"matching term with wild card", inputText: "I was born in a land downunder", searchTerm: "born*under",
			wantResults: []Segment{{Text: "I was ", Highlight: false},
				{Text: "born", Highlight: true},
				{Text: " in a land down", Highlight: false},
				{Text: "under", Highlight: true}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := Parse(tt.inputText, tt.searchTerm); !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("parse() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}