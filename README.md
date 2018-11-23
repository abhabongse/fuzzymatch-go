# Fuzzy Match

[![GoDoc](https://godoc.org/github.com/abhabongse/fuzzymatch-go/pkg/fuzzymatch?status.svg)](https://godoc.org/github.com/abhabongse/fuzzymatch-go/pkg/fuzzymatch)

This repository contains a Go language implementation of an approximate string matching algorithm. Particularly,
the function `fuzzymatch.SimilarityScore` determines how the similarity score between two input strings. It has
the following type signature.

```go
func SimilarityScore(fst, snd string) float64 { ... }
```

All source code for this project is released under the [GNU Lesser General Public License v3.0](LICENSE).

