package nametitle

import (
	"github.com/abhabongse/fuzzymatch-go/candidate"
	"regexp"
)

// GenerateNamesWithoutTitles is a preset function which generates a sequence of
// all possible bare names (or names without prefixed or suffixed titles).
// It is built upon the function DecomposeName with the patterns from DefaultNameTitlePatterns.
var GenerateNamesWithoutTitles = candidate.MakeRegexpGenerator(DefaultNameTitlePatterns)

// DefaultNameTitlePatterns is a sequence of all compiled regular expression objects
// which separates bare names from their prefixed and suffixed titles in a full name string.
// As of current, only common English and Thai prefixed titles are handled.
var DefaultNameTitlePatterns = []*regexp.Regexp{
	regexp.MustCompile("^(.*)$"),
	// English full prefixed titles: space separator required
	regexp.MustCompile("^(?:mister )(.*)$"),
	regexp.MustCompile("^(?:miss )(.*)$"),
	regexp.MustCompile("^(?:master )(.*)$"),
	// English abbreviated prefixed titles: separator required
	regexp.MustCompile("^(?:mr(?: |\\. |\\.))(.*)$"),
	regexp.MustCompile("^(?:mrs(?: |\\. |\\.))(.*)$"),
	regexp.MustCompile("^(?:ms(?: |\\. |\\.))(.*)$"),
	// Thai full prefixed titles: separator optional
	regexp.MustCompile("^(?:นาย ?)(.*)$"),
	regexp.MustCompile("^(?:นาง ?)(.*)$"),
	regexp.MustCompile("^(?:นางสาว ?)(.*)$"),
	regexp.MustCompile("^(?:เด็กชาย ?)(.*)$"),
	regexp.MustCompile("^(?:เด็กหญิง ?)(.*)$"),
	// Thai abbreviated prefixed titles: separator required
	regexp.MustCompile("^(?:ดช(?: |\\. |\\.))(.*)$"),
	regexp.MustCompile("^(?:ดญ(?: |\\. |\\.))(.*)$"),
	// Thai dot-abbreviated prefixed titles: separator optional
	regexp.MustCompile("^(?:น\\.ส\\. ?)(.*)$"),
	regexp.MustCompile("^(?:ด\\.ช\\. ?)(.*)$"),
	regexp.MustCompile("^(?:ด\\.ญ\\. ?)(.*)$"),
}
