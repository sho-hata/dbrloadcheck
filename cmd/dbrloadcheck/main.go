package main

import (
	"github/sho-hata/dbrloadcheck"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(dbrloadcheck.Analyzer) }
