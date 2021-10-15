package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Application struct {
	baseDir     string
	log         *log.Logger
	engineText  *text.Template
	fileContent string
}

func NewApplication(dir) *Application {
	tpl, err := template.New()
	return &Application{
		baseDir: dir,
	}
}

func (a *Application) readFile(fName string) map[string]string {
	if _, err := os.Stat(fName); err == nil {
		text, err := ioutil.ReadAll(file)
		lines := strings.Split(string(text), "\n")
		if err != nil {
			log.Fatal(err)
		}
		var template strings.Builder
		var config strings.Builder
		var isYaml = true
		for line := range lines {
			if line == "----" {
				isYaml = false
				continue
			}
			if isYaml {
				config.WriteString(line + "\n")
			} else {
				template.WriteString(line + "\n")
			}
		}
		ret := make(map[string]string)
	} else if os.IsNotExist(err) {
		log.Fatal(fmt.Sprintf("Archivo `%s` no existe.", fName))
	}
}
