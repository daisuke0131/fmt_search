package fmt_search

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"strconv"
)

const doc = "fmt_search is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "fmt_search",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files{
		for _,imp := range  f.Imports{
			v, err := strconv.Unquote(imp.Path.Value)
			if err != nil{
				return nil, err
			}
			if "fmt" == v {
				pass.Reportf(imp.Pos(), "%s is imported.", v)
			}
		}
	}

	return nil, nil
}
