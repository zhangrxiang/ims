package utils

import (
	"strconv"
	"strings"
)

func StrToIntAlice(str, sep string) []int {
	var intStr []int
	split := strings.Split(str, sep)
	for _, v := range split {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil
		}
		intStr = append(intStr, i)
	}

	return intStr
}
