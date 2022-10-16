package commandparser

import (
	"errors"
	"strconv"
)

var ErrNotEnoughArgs = errors.New("not enough args")
var ErrWrongNumberOfArgs = errors.New("wrong number of args")
var ErrTargetNotAPointer = errors.New("target must be a pointer")
var ErrTargetNilPointer = errors.New("target must be a non-null pointer")
var ErrTargetNotAStruct = errors.New("target must be a struct")
var ErrNamedArgMissing = errors.New("named arg is missing")
var ErrParsingInt = errors.New("parsing int error")

type CommanParser interface {
	Parse(args [][]byte) (args_reminder [][]byte, err error)
}

func parseInt(args [][]byte) (result int, next [][]byte, err error) {
	if len(args) < 1 {
		return 0, args, ErrNotEnoughArgs
	}
	result32, err := strconv.ParseInt(string(args[0]), 10, 32)
	if err != nil {
		return 0, args, ErrParsingInt
	}
	return int(result32), args[1:], nil
}

func parseString(args [][]byte) (result string, next [][]byte, err error) {
	if len(args) < 1 {
		return "", args, ErrNotEnoughArgs
	}

	return string(args[0]), args[1:], nil
}

func optionalPresent(args [][]byte, name string) (result bool, next [][]byte, err error) {
	if len(args) < 1 {
		return false, args, nil
	}

	if string(args[0]) == name {
		// present, progress args
		return true, args[1:], nil
	}
	// we don't progress args
	return false, args, nil
}
