package main

import (
	"errors"
	"time"
)

// Duration example
type DurationValue struct {
	Duration *time.Duration
}

func (v DurationValue) String() string {
	if v.Duration != nil {
		return v.Duration.String()
	}
	return ``
}

func (v DurationValue) Set(s string) error {
	d, err := time.ParseDuration(s)
	if err != nil {
		return errors.New(`Format is not acceptable`)
	}
	*v.Duration = d
	return nil
}
