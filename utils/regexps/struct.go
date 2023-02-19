package regexps

import "regexp"

type Flags string
type Flag uint8

const (
	Global Flag = 'g'
	Force  Flag = 'f'
)

type Regexp struct {
	exp    *regexp.Regexp
	source *string
	Flags  *Flags
}
