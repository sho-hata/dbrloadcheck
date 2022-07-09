package dbrloadcheck_test

import (
	"testing"

	"github/sho-hata/dbrloadcheck"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, dbrloadcheck.Analyzer, "a")
}
