package request

import (
	"errors"
	"io"
	"strings"
)

var (
	errInvalidMethod = errors.New("invalid method: must contain only uppercase letters")
	errInvalidHTTP   = errors.New("invalid HTTP version: only HTTP/1.1 supported")
	errInvalidData   = errors.New("invalid data request")
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	rawData, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(rawData), "\r\n")
	requestLineParts := strings.Split(lines[0], " ")

	if len(requestLineParts) != 3 {
		return nil, errInvalidData
	}

	method := requestLineParts[0]
	requestTarget := requestLineParts[1]
	httpVersionRaw := requestLineParts[2]

	// valid method contains only uppercase letter. For example: GET, POST, PUT, DELETE
	for _, char := range method {
		if char < 'A' || char > 'Z' {
			return nil, errInvalidMethod
		}
	}

	// validate HTTP version
	if !strings.HasPrefix(httpVersionRaw, "HTTP/") || httpVersionRaw != "HTTP/1.1" {
		return nil, errInvalidHTTP
	}

	httpVersion := strings.TrimPrefix(httpVersionRaw, "HTTP/")

	reqLine := RequestLine{
		HttpVersion:   httpVersion,
		RequestTarget: requestTarget,
		Method:        method,
	}

	return &Request{
		RequestLine: reqLine,
	}, nil
}
