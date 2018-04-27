package templateUtil

import (
	"bytes"
	"log"
	"text/template"
)

func TmplExecuteMust(tp *template.Template, buf *bytes.Buffer, data interface{}) error {
	err := tp.Execute(buf, data)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

