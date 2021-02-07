package main

import (
	"fmt"
	"strings"
)

const arraySeparator = `,`

// Array example
type ArrayValue struct {
	Array *[]string
}

func (v ArrayValue) String() string {
	if v.Array != nil {
		return fmt.Sprintf("%#v", *v.Array)
	}
	return ``
}

func (v ArrayValue) Set(s string) error {
	*v.Array = strings.Split(s, arraySeparator)
	return nil
}
