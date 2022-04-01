package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	typesFlag := flag.String("types", "", "")
	pkg := flag.String("package", "main", "")
	flag.Parse()
	types := strings.Split(*typesFlag, ",")

	data := map[string]interface{}{
		"package": *pkg,
		"types":   types,
	}

	t := template.New("")
	t.Funcs(map[string]interface{}{
		"title": strings.Title,
	})
	template.Must(t.Parse(templateContent))
	f, err := os.Create("typedchannels.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = t.Execute(f, data)
	if err != nil {
		log.Fatal(err)
	}
}

const templateContent = `
// Code generated DO NOT EDIT

package {{.package}}

{{range .types}}
{{ $type := printf "%vSmartChannel" (title .)}}
type {{$type}} struct {
	smartChannel
}
func (sc {{$type}}) Send(msg {{.}}) error {
	return sc.smartChannel.Send(msg)
}
func (sc {{$type}}) Receive() ({{.}}, bool) {
	if msg, ok := sc.smartChannel.Receive(); ok {
		result, ok := msg.({{.}})
		return result, ok
	} else {
		var result {{.}}
		return result, ok
	}
}
{{end}}

`
