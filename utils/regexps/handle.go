package regexps

import (
	"regexp"
	"strings"
)

// Compile makes the world more complicated by adopting
// JavaScript-style for Regex expressions.
// This shit is easily scalable. Regex mode will
// be used only if the source starts with /
func Compile(source string) (*Regexp, error) {
	var flags Flags

	// get Flags
	if source[0] == '/' {
		flagEnd := -1
		for i := len(source) - 1; i > 0; i-- {
			if source[i] == '/' {
				flagEnd = i
				break
			}
		}

		if flagEnd > 0 {
			flags = Flags(source[flagEnd+1:])
			source = source[1:flagEnd]
		}
		regex, err := regexp.Compile(source)
		return &Regexp{regex, nil, &flags}, err
	} else {
		src := source
		return &Regexp{nil, &src, nil}, nil
	}
}

func (exp *Regexp) HasFlag(flag Flag) bool {
	if !exp.HasFlags() {
		return false
	}
	for _, f := range *exp.Flags {
		if Flag(f) == flag {
			return true
		}
	}
	return false
}

// HasFlags determines whether an expression is
// in Regex mode.
func (exp *Regexp) HasFlags() bool {
	return exp.Flags != nil
}

// Match matches a string with the given expression.
func (exp *Regexp) Match(str string) bool {
	if !exp.HasFlags() {
		return strings.Contains(str, *exp.source)
	}
	return exp.exp.MatchString(str)
}
