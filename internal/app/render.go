package app

import (
	"regexp"
	"text/template"
)

var reErr = regexp.Compile(`template snippet:(\d+):(\d+)?:? (.+$)`)

type Render struct {
	vars     map[string]interface{}
	funcMap  template.FuncMap
	template *template.Template
}
