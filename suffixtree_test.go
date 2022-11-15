package suffixtree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestSuffixTree(t *testing.T) {
	tests := []struct {
		name               string
		words              []string
		matchString        string
		numElementsToMatch int
		wantNumMatches     int
		wantMatchedWords   []string
	}{
		{
			name:               "expect all matches (3) since match string is common in all words",
			words:              []string{"banana", "apple", "中文app"},
			matchString:        "a",
			numElementsToMatch: -1,
			wantNumMatches:     3,
			wantMatchedWords:   []string{"banana", "apple", "中文app"},
		},
		{
			name:               "expect first word to be a match",
			words:              []string{"bananal", "applel", "中文appz"},
			matchString:        "al",
			numElementsToMatch: -1,
			wantNumMatches:     1,
			wantMatchedWords:   []string{"bananal"},
		},
		{
			name:               "expect last word to be a match",
			words:              []string{"bananal", "applel", "中文appz"},
			matchString:        "pz",
			numElementsToMatch: -1,
			wantNumMatches:     1,
			wantMatchedWords:   []string{"中文appz"},
		},
		{
			name:               "expect a match when substring is multiple words",
			words:              []string{"banana is cool", "apple is cold", "中文app is warm"},
			matchString:        "is cold",
			numElementsToMatch: -1,
			wantNumMatches:     1,
			wantMatchedWords:   []string{"apple is cold"},
		},
		{
			name:               "expect no match when substring is an empty string",
			words:              []string{"banana is cool", "apple is cold", "中文app is warm"},
			matchString:        "",
			numElementsToMatch: -1,
			wantNumMatches:     0,
			wantMatchedWords:   []string{},
		},
		{
			name:               "expect a match when substring is a number",
			words:              []string{"banana is cool 100", "apple is cold 99", "中文app is warm 98"},
			matchString:        "100",
			numElementsToMatch: -1,
			wantNumMatches:     1,
			wantMatchedWords:   []string{"banana is cool 100"},
		},
		{
			name:               "expect 1 match when numElementsToMatch is 1",
			words:              []string{"banana is cool 999", "apple is cold 998", "中文app is warm 997"},
			matchString:        "99",
			numElementsToMatch: 1,
			wantNumMatches:     1,
			wantMatchedWords:   []string{"banana is cool 999"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := NewGeneralizedSuffixTree()
			for k, word := range test.words {
				tree.Put(word, k)
			}

			indexes := tree.Search(test.matchString, test.numElementsToMatch)
			gotNumMatches := len(indexes)
			if gotNumMatches != test.wantNumMatches {
				t.Fatalf("expected %d matches, got %d matches", test.wantNumMatches, gotNumMatches)
			}

			gotMatchedWords := []string{}
			for _, index := range indexes {
				gotMatchedWords = append(gotMatchedWords, test.words[index])
			}

			less := func(a, b string) bool { return a < b }
			equalIgnoreOrder := cmp.Equal(gotMatchedWords, test.wantMatchedWords, cmpopts.SortSlices(less))
			if !equalIgnoreOrder {
				t.Fatalf("expected match words: %v, got match words: %v", test.wantMatchedWords, gotMatchedWords)
			}
		})
	}
}
