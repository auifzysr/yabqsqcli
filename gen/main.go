//go:generate go run .
//go:generate gofmt -w ../

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func head(s string) string {
	return s[0:1]
}

func main() {
	for _, cfg := range configs {
		var flagString string
		for _, opt := range cfg.Options {
			flagString += flagTemplates[opt] + "\n"
		}
		cfg.FlagTemplate = flagString
		fn := fmt.Sprintf("../cmd/%s_gen.go", cfg.Name)
		funcMap := template.FuncMap{
			"capitalize": capitalize,
			"head":       head,
		}
		t := template.Must(template.New("gen").Funcs(funcMap).Parse(cmdTemplate))
		buf := &bytes.Buffer{}
		if err := t.Execute(buf, cfg); err != nil {
			log.Fatalf("failed to execute template: %v", err)
		}

		if err := os.WriteFile(fn, buf.Bytes(), 0644); err != nil {
			log.Fatalf("failed to write %s: %v", fn, err)
		}
		log.Printf("generated %s\n", fn)
	}
}
