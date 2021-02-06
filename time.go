package main

import (
	"errors"
	"time"
)

const defaultLayout = time.RFC3339

// define acceptable layouts
var expectedLayouts = []string{
	defaultLayout,
	time.RFC1123,
	time.RFC1123Z,
}

// Time example
type TimeValue struct {
	Time *time.Time
}

func (v TimeValue) String() string {
	if v.Time != nil {
		return v.Time.Format(defaultLayout)
	}
	return ``
}

func (v TimeValue) Set(s string) error {
	success := false
	for _, layout := range expectedLayouts {
		t, err := time.Parse(layout, s)
		if err != nil {
			continue
		}
		*v.Time = t
		success = true
	}
	if !success {
		return errors.New(`Format is not acceptable`)
	}
	return nil
}
