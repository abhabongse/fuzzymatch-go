# Fuzzy Match

[![pkg.go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/abhabongse/fuzzymatch-go)

This repository contains a Go language implementation of approximate string matching algorithms.

### Preset functions

Inside the package [`github.com/abhabongse/fuzzymatch-go/factory/preset`](https://godoc.org/github.com/abhabongse/fuzzymatch-go/factory/preset),
there are three examples of preset string similarity score functions with different customizations:

- [SimpleSimilarityScore](factory/preset/simple.go)
- [PlanSimilarityScore](factory/preset/plain.go)
- [ThaiNameSimilarityScore](factory/preset/thai_name.go)

All of these functions have output values between 0 and 1 (inclusive),
where 1 indicates that both input strings are identical under some criteria
and 0 indicates that both strings are totally distinct.
The signature of these functions is
```go
func(string, string) float64
```

### Customization

New string similarity score functions may be constructed via various higher-order functions provided in this module.
Look at the construction of the preset functions above for some ideas
of how to introduce addition functionalities to your string similarity score functions.

### Notes

All source code for this project is released under the [GNU Lesser General Public License v3.0](LICENSE).

