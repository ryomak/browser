package main

import (
	"flag"
	"github.com/ryomak/browser/src/browser"
)

func main() {
	flag.Parse()
	browser.Browse(flag.Arg(0))
}
