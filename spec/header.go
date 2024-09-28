package spec

import (
	"fmt"
	"regexp"
)

type HeaderField struct {
	Name  			string
	Value 			string

	isCaseSensitive bool
}

func ParseHeaderField(line string) (*HeaderField, error) {
	headerRegex := regexp.MustCompile(`^([!#$%&'*+\-\.^_`+"`"+`|~0-9A-Za-z]+):\s*(.*?)\s*$`)

	matches := headerRegex.FindStringSubmatch(line)
	if len(matches) != 2 {
		return nil, fmt.Errorf("invalid header field: %s", line)
	}
	return &HeaderField{
		Name: matches[1],
		Value: matches[2],
	}, nil
}
