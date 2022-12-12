package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/planetscale/enumcheck/enumcheck"
)

func main() { singlechecker.Main(enumcheck.Analyzer) }
