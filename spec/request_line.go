package spec

import (
	"fmt"
	"regexp"
	"strconv"
)

type RequestLine struct {
	Method			string
	URI   			string
	MajorVersion 	int
	MinorVersion 	int
}

func ParseRequestLine(line string) (*RequestLine, error) {
	requestLineRegex := regexp.MustCompile(`^([A-Za-z0-9-]+) +([^ ]+) +HTTP/([0-9]+)\.([0-9]+)\r\n$`)
	matches := requestLineRegex.FindStringSubmatch(line)

	if len(matches) != 4 {
		return nil, fmt.Errorf("invalid request line: %s", line)
	}

	majorVer, err := strconv.Atoi(matches[3])
	if err != nil {
		return nil, fmt.Errorf("invalid major version: %s", matches[3])
	}

	minorVer, err := strconv.Atoi(matches[4])
	if err != nil {
		return nil, fmt.Errorf("invalid minor version: %s", matches[4])
	}

	return &RequestLine{
		Method: matches[1],
		URI: matches[2],
		MajorVersion: majorVer,
		MinorVersion: minorVer,
	}, nil
}