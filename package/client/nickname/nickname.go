package nickname

import (
	"errors"
	"regexp"
)

type Nickname string

var EmptyNicknameError = errors.New("user nickname cannot be empty")
var ForbiddenCharError = errors.New("user nickname cannot contain forbidden character")
var ForbiddenStartingCharError = errors.New("user nickname cannot begin with forbidden charater")

var forbiddenCharsRegex = regexp.MustCompile(`[ ,*?!@.]`)
var forbiddenStartingCharsRegex = regexp.MustCompile(`^([$:#&~@%+]|\+q|\+a|\+o|\+h|\+v)`)

// validate if nick follows all rules given in the RFC:
// https://modern.ircdocs.horse/#clients
// following the pattern "Parse don't validate" from:
// https://lexi-lambda.github.io/blog/2019/11/05/parse-don-t-validate/
func New(nickname string) (Nickname, error) {
	if nickname == "" {
		return Nickname(""), EmptyNicknameError
	}

	if forbiddenCharsRegex.MatchString(nickname) {
		return Nickname(""), ForbiddenCharError
	}

	if forbiddenStartingCharsRegex.MatchString(nickname) {
		return Nickname(""), ForbiddenStartingCharError
	}

	return Nickname(nickname), nil
}
