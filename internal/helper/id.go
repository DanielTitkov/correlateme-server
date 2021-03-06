package helper

import (
	"errors"
	"strconv"
	"strings"
)

const sep = "_"

func PairOfIDsToString(left, right int) string {
	return strings.Join([]string{
		strconv.Itoa(left),
		strconv.Itoa(right),
	}, sep) // separator is for readability
}

func StringToPairOfIDs(s string) (int, int, error) {
	subs := strings.Split(s, sep)
	if len(subs) != 2 {
		return 0, 0, errors.New("failed to split string into 2 substrings")
	}
	left, err := strconv.Atoi(subs[0])
	if err != nil {
		return 0, 0, err
	}
	right, err := strconv.Atoi(subs[1])
	if err != nil {
		return 0, 0, err
	}
	return left, right, nil
}
