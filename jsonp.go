package main

import (
	"fmt"
	"io"
)

type JsonpWrapper struct {
	Prefix     string
	Underlying io.Reader

	gotPrefix bool
}

func (jpw *JsonpWrapper) Read(b []byte) (int, error) {
	if jpw.gotPrefix {
		return jpw.Underlying.Read(b)
	}

	prefix := make([]byte, len(jpw.Prefix))
	n, err := io.ReadFull(jpw.Underlying, prefix)
	if err != nil {
		return n, err
	}

	if string(prefix) != jpw.Prefix {
		return n, fmt.Errorf("JSONP prefix mismatch: expected %q, got %q",
			jpw.Prefix, prefix)
	}

	char := make([]byte, 1)
	for char[0] != '(' {
		n, err = jpw.Underlying.Read(char)
		if n == 0 || err != nil {
			return n, err
		}
	}

	jpw.gotPrefix = true
	return jpw.Underlying.Read(b)
}
