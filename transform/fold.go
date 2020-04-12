package transform

import "golang.org/x/text/cases"

// CaseFoldingTransformer is a Unicode stream transformer object
// which performs case folding on all unicode characters.
// Characters will generally be transformed into its lowercase forms.
// Special characters such as 'ß' will also be converted (to 'ss').
var CaseFoldingTransformer = cases.Fold(cases.HandleFinalSigma(true))
