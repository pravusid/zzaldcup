package helper

import "strconv"

func ParseInt(param string, defaultValue uint64) uint64 {
	str, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return defaultValue
	}
	return str
}
