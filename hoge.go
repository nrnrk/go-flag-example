package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Hoge struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *Hoge) String() string {
	return fmt.Sprintf("Hoge{id: %d, name: %s}", h.ID, h.Name)
}

// Hoge (Custom struct) example
type HogeValue struct {
	Hoge *Hoge
}

func (v HogeValue) String() string {
	if v.Hoge != nil {
		return v.Hoge.String()
	}
	return ``
}

func (v HogeValue) Set(s string) error {
	var hoge Hoge
	if err := json.Unmarshal([]byte(s), &hoge); err != nil {
		return errors.New(`Format is not acceptable`)
	}
	*v.Hoge = hoge
	return nil
}
