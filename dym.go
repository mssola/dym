// Copyright (C) 2014 Miquel Sabaté Solà
// This file is licensed under the MIT license.
// See the LICENSE file.

package dym

// The default distance as used by the Similar function.
const DEFAULT_DISTANCE = 3

// Compute the distance between the given strings by using the
// Damerau-Levenshtein algorithm. It says how many letters need to be swapped,
// substituted, deleted from, or added to string1, at least, to get string2.
//
// Returns an integer with the distance value.
func levenshtein(string1, string2 string) int {
	var last, old, tmp int
	l1, l2 := len(string1), len(string2)
	d := make([]int, l1+1)

	for i := 1; i <= l1; i++ {
		d[i] = i
	}
	for i := 1; i <= l2; i++ {
		d[0] = i
		last = i - 1
		for j := 1; j <= l1; j++ {
			old = d[j]
			min := d[j] + 1
			if (d[j-1] + 1) < min {
				min = d[j-1] + 1
			}
			if (string1)[j-1] == (string2)[i-1] {
				tmp = 0
			} else {
				tmp = 1
			}
			if (last + tmp) < min {
				min = last + tmp
			}
			d[j] = min
			last = old
		}
	}
	return d[l1]
}

// Get all the words from the given dictionary that are close to the given word
// by a certain amount of distance.
//
// dict - A slice of strings containing the dictionary of words.
// word - The word that we are looking for.
// dist - The given distance.
//
// Returns a new slice of strings containing the results.
func SimilarDistance(dict []string, word string, dist int) []string {
	res := []string{}

	for _, r := range dict {
		if levenshtein(r, word) <= dist {
			res = append(res, r)
		}
	}
	return res
}

// Get all the words from the given dictionary that are close to the given
// word.
//
// dict - A slice of strings containing the dictionary of words.
// word - The word that we are looking for.
//
// Returns a new slice of strings containing the results.
func Similar(dict []string, word string) []string {
	return SimilarDistance(dict, word, DEFAULT_DISTANCE)
}
