package convert

import (
	"encoding/json"
	"strconv"
)

// IntArrToStringArr is func to convert array int to array string
func IntArrToStringArr(a []int) []string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return b
}

// ParseToDestination is func to parse from interface to destination struct
func ParseToDestination(resp interface{}, dest interface{}) error {
	bytes, err := json.Marshal(resp)

	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, dest)

	if err != nil {
		return err
	}

	return err
}
