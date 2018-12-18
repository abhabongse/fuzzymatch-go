# Fuzzy Match

[![GoDoc](https://godoc.org/github.com/abhabongse/fuzzymatch-go?status.svg)](https://godoc.org/github.com/abhabongse/fuzzymatch-go)

This repository contains a Go language implementation of an approximate string matching algorithm. Particularly,
the function `preset.PlainSimilarityScore` determines how the similarity score between two input strings. It has
the following type signature.

```go
func PlainSimilarityScore(fst, snd string) float64 { ... }
```

All source code for this project is released under the [GNU Lesser General Public License v3.0](LICENSE).

