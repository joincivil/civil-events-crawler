// Package gen contains all the components for code generation.
package gen

import (
	"bytes"
	log "github.com/golang/glog"
	"go/format"
	"io"
	"text/template"
)

// OutputTemplatedData is a shared function to output templated data given
// the template and template data to the writer.
func OutputTemplatedData(writer io.Writer, tmplName string, tmpl string,
	tmplData interface{}, gofmt bool) error {
	t := template.Must(template.New(tmplName).Parse(tmpl))
	buf := &bytes.Buffer{}
	err := t.Execute(buf, tmplData)
	if err != nil {
		return err
	}
	output := buf.Bytes()
	if gofmt {
		output, err = format.Source(buf.Bytes())
		if err != nil {
			log.Errorf("ERROR Gofmt: err:%v\ntemplate generated:\n%v", err, buf.String())
			return err
		}
	}
	_, err = writer.Write(output)
	return err
}
