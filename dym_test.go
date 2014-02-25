// Copyright (C) 2014 Miquel Sabaté Solà <mikisabate@gmail.com>
// This file is licensed under the MIT license.
// See the LICENSE file.

package dym

import (
	"testing"
)

func TestLevenshtein(t *testing.T) {
	a, b, c := "thing", "thingie", "ting"

	if levenshtein(a, b) != 2 {
		t.Errorf("Levenshtein of %v and %v is not 2", a, b)
	}
	if levenshtein(a, c) != 1 {
		t.Errorf("Levenshtein of %v and %v is not 1", a, c)
	}
	if levenshtein(b, c) != 3 {
		t.Errorf("Levenshtein of %v and %v is not 3", b, c)
	}
	if levenshtein(a, a) != 0 {
		t.Errorf("Levenshtein of %v and %v is not 0", a, a)
	}
}

func TestIsSimilar(t *testing.T) {
	if !IsSimilar("thing", "ting") {
		t.Errorf("they're not really similar")
	}
	if IsSimilarDistance("thing", "tingu", 1) {
		t.Errorf("they're similar")
	}
}

func TestSimilar(t *testing.T) {
	a := "thing"
	dict := []string{"thingie", "lala", "ting"}

	results := Similar(dict, a)
	if len(results) != 2 {
		t.Errorf("wrong number of results")
	}
	if results[0] != "thingie" || results[1] != "ting" {
		t.Errorf("wrong results")
	}

	results = SimilarDistance(dict, a, 1)
	if len(results) != 1 {
		t.Errorf("wrong number of results")
	}
	if results[0] != "ting" {
		t.Errorf("wrong results")
	}

	results = SimilarDistance(dict, a, -1)
	if len(results) != 0 {
		t.Errorf("wrong number of results")
	}
}
