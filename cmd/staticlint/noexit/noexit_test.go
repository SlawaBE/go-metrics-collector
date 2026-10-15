package noexit_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/SlawaBE/go-metrics-collector/cmd/staticlint/noexit"
)

// TestAnalyzer verifies that the noexit analyzer reports only direct os.Exit
// calls in the main function of the main package and ignores calls in other
// functions.
func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), noexit.Analyzer, "exit")
}
