package client

import (
	"errors"
	"log/slog"
	"regexp"
)

type Nickname string

var EmptyNicknameError = errors.New("client nickname cannot be empty")
var ForbiddenCharError = errors.New("client nickname cannot contain forbidden character")
var ForbiddenStartingCharError = errors.New("client nickname cannot begin with forbidden charater")

var forbiddenCharsRegex = regexp.MustCompile(`[ ,*?!@.]`)
var forbiddenStartingCharsRegex = regexp.MustCompile(`^([$:#&~@%+]|\+q|\+a|\+o|\+h|\+v)`)

// validate if nick follows all rules given in the RFC:
// https://modern.ircdocs.horse/#clients
func ValidateNickname(nickname string) error {
	slog.Info("walidating nickname", "nick", nickname)
	if nickname == "" {
		return EmptyNicknameError
	}

	if forbiddenCharsRegex.MatchString(nickname) {
		return ForbiddenCharError
	}

	if forbiddenStartingCharsRegex.MatchString(nickname) {
		return ForbiddenStartingCharError
	}
	slog.Info("nick name was correct", "nick", nickname)

	return nil
}
