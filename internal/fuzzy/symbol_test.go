// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fuzzy_test

import (
	"testing"

	. "github.com/bafko/tools/internal/fuzzy"
)

func TestSymbolMatchIndex(t *testing.T) {
	tests := []struct {
		pattern, input string
		want           int
	}{
		{"test", "foo.TestFoo", 4},
		{"test", "test", 0},
		{"test", "Test", 0},
		{"test", "est", -1},
		{"t", "shortest", 7},
		{"", "foo", -1},
		{"", string([]rune{0}), -1}, // verify that we don't default to an empty pattern.
		{"anything", "", -1},
	}

	for _, test := range tests {
		matcher := NewSymbolMatcher(test.pattern)
		if got, _ := matcher.Match([]string{test.input}); got != test.want {
			t.Errorf("NewSymbolMatcher(%q).Match(%q) = %v, _, want %v, _", test.pattern, test.input, got, test.want)
		}
	}
}

func TestSymbolRanking(t *testing.T) {
	matcher := NewSymbolMatcher("test")

	// symbols to match, in ascending order of ranking.
	symbols := []string{
		"this.is.better.than.most",
		"test.foo.bar",
		"thebest",
		"test.foo",
		"test.foo",
		"atest",
		"testage",
		"tTest",
		"foo.test",
		"test",
	}
	prev := 0.0
	for _, sym := range symbols {
		_, score := matcher.Match([]string{sym})
		t.Logf("Match(%q) = %v", sym, score)
		if score < prev {
			t.Errorf("Match(%q) = _, %v, want > %v", sym, score, prev)
		}
		prev = score
	}
}

// Test that we strongly prefer exact matches.
//
// In golang/go#60027, we preferred "Runner" for the query "rune" over several
// results containing the word "rune" exactly. Following this observation,
// scoring was tweaked to more strongly emphasize sequential characters and
// exact matches.
func TestSymbolRanking_Issue60027(t *testing.T) {
	matcher := NewSymbolMatcher("rune")

	// symbols to match, in ascending order of ranking.
	symbols := []string{
		"Runner",
		"singleRuneParam",
		"Config.ifsRune",
		"Parser.rune",
	}
	prev := 0.0
	for _, sym := range symbols {
		_, score := matcher.Match([]string{sym})
		t.Logf("Match(%q) = %v", sym, score)
		if score < prev {
			t.Errorf("Match(%q) = _, %v, want > %v", sym, score, prev)
		}
		prev = score
	}
}

func TestChunkedMatch(t *testing.T) {
	matcher := NewSymbolMatcher("test")

	chunked := [][]string{
		{"test"},
		{"", "test"},
		{"test", ""},
		{"te", "st"},
	}

	for _, chunks := range chunked {
		offset, score := matcher.Match(chunks)
		if offset != 0 || score != 1.0 {
			t.Errorf("Match(%v) = %v, %v, want 0, 1.0", chunks, offset, score)
		}
	}
}
