package main

import (
	"errors"
	"net/url"
)

// URL example from https://golang.org/pkg/flag/#Value
type URLValue struct {
	URL *url.URL
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return errors.New(`Format is not acceptable`)
	} else {
		*v.URL = *u
	}
	return nil
}
