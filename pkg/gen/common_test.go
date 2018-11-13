package gen_test

import (
	"bytes"
	"testing"

	"github.com/joincivil/civil-events-crawler/pkg/gen"
)

var badTestTmpl = `{{.BadVarUnknown}}`
var badTestTmpl2 = `
package test

import (
	"testing"

func BadSyntax( string {

}
`

func TestOutputTemplateData(t *testing.T) {
	bys := &bytes.Buffer{}

	testTmplData := struct {
		TestData string
	}{
		TestData: "data",
	}

	err := gen.OutputTemplatedData(bys, "tmp.tmpl", badTestTmpl, testTmplData, true)
	if err == nil {
		t.Errorf("Should have failed to output templated data: %v", err)
	}
	if bys.String() != "" {
		t.Errorf("Should have returned empty templated data: %v", bys.String())
	}

	err = gen.OutputTemplatedData(bys, "tmp.tmpl", badTestTmpl2, testTmplData, true)
	if err == nil {
		t.Errorf("Should have failed to output templated data: %v", err)
	}
	if bys.String() != "" {
		t.Errorf("Should have returned empty templated data: %v", bys.String())
	}
}
