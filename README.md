# Fuzzy Match

[![GoDoc](https://godoc.org/github.com/abhabongse/fuzzymatch-go?status.svg)](https://godoc.org/github.com/abhabongse/fuzzymatch-go)

This repository contains a Go language implementation of approximate string matching algorithms.

### Preset functions

Inside the package `github.com/abhabongse/fuzzymatch-go/factory/presets`, there is
the function `PlainSimilarityScore` which determines the similarity score between two generic
input strings. Another function, `ThaiNameSimilarityScore` is a customized version for the
similarity scoring function but it has extra pre-processings and string comparison logic for
names of Thai people. 

Both functions have output values between 0 and 1, where 0 means that two
strings are very distinct whereas 0 means that two strings are very similar. Here are the 
signatures for both functions.
 
```go
func PlainSimilarityScore(fst, snd string) float64 { ... }

func ThaiNameSimilarityScore(fst, snd string) float64 { ... }
```

### Customization

The logical flow for string similarity scoring functions may be customized by using the factory
(inside the package `github.com/abhabongse/fuzzymatch-go/factory`) to construct a new function. 
The factory accepts various options; see the documentation for more information.

### Notes

All source code for this project is released under the [GNU Lesser General Public License v3.0](LICENSE).

