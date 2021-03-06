# Did you mean? [![Build Status](https://travis-ci.org/mssola/dym.png?branch=master)](https://travis-ci.org/mssola/dym) [![GoDoc](https://godoc.org/github.com/mssola/dym?status.png)](http://godoc.org/github.com/mssola/dym)

This package gets the inspiration from git's "Did you mean this?" feature.
That is, it detects typos and it offers similar words that might be what the
user really wanted to type. This package offers two public methods:

~~~ go
func Similar(dict []string, word string) []string
func SimilarDistance(dict []string, word string, dist int) []string
~~~

The first method returns all the words from the given dictionary that are close
to the given word. The second method is identical to the first one but it also
accepts a distance (the number of swaps, additions, deletions, etc. allowed).
Finally, this package also provides a couple of helper methods:

~~~ go
func IsSimilar(string1, string2 string) bool
func IsSimilarDistance(string1, string2 string, dist int) bool
~~~

The first method returns whether the two given strings are similar. The second
method does the same but it also requires a distance.

## Usage

Here goes an example:

~~~ go
package main

import "github.com/mssola/dym"

func main() {
  a := "thing"
  dict := []string{"thingie", "lala", "ting"}

  results := dym.Similar(dict, a)
  // results now contains: ["thingie", "ting"]

  results = dym.SimilarDistance(dict, a, 1)
  // results now contains: ["ting"]
}
~~~

Copyright &copy; 2014-2016 Miquel Sabaté Solà, released under the MIT License.

