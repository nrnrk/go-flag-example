package main

import (
	"flag"
	"fmt"
	"net/url"
	"time"
)

var u = &url.URL{}
var t = &time.Time{}
var d time.Duration
var h = &Hoge{}

func main() {
	fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	fs.Var(&URLValue{u}, "url", "URL to parse")
	fs.Var(&TimeValue{t}, "time", "Time to parse (RFC3339, RFC1123, RFC1123Z)")
	fs.Var(&DurationValue{&d}, "duration", "Duration to parse (supported units: 'ns', 'ms', 's', 'm', 'h') ")
	fs.Var(&HogeValue{h}, "hoge", "Hoge to parse (json format)")

	fs.Parse([]string{
		"-url",
		"https://golang.org/pkg/flag/",
		"-time",
		"Mon, 02 Jan 2006 15:04:05 MST",
		"-duration",
		"75s",
		"-hoge",
		`{"id": 352, "name": "Diego"}`,
	})

	// URL
	fmt.Printf("url: {scheme: %q, host: %q, path: %q}\n", u.Scheme, u.Host, u.Path)
	// Time
	fmt.Printf("time: %q\n", t.Format(defaultLayout))
	// Duration
	fmt.Printf("duration: %d\n", d)
	// Hoge (Custom struct)
	fmt.Printf("hoge: %#v\n", h)
}
