package enumcheck_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/planetscale/enumcheck/enumcheck"
)

func TestFromFileSystem(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, enumcheck.Analyzer,
		"enumbyte",
		"enumpartial",
		"enumstring",
		"enumstring2",
		"enumstruct",
		"enumtype",
		"indirect",
	)
}
