package main

import (
	"github.com/daisuke0131/fmt_search"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(fmt_search.Analyzer) }
