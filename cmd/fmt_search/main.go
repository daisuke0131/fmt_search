package main

import (
	"fmt_search"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(fmt_search.Analyzer) }
