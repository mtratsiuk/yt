package main

import (
	"embed"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed *.go.template
var templateFs embed.FS

var assetsPath = flag.String("in", ".", "path to assets folder")
var assetsGlob = flag.String("g", "**/*.png", "assets glob pattern")
var outPath = flag.String("out", "assets.go", "out file name")

type TemplateData struct {
	PackageName string
	Assets      []TemplateAsset
}

type TemplateAsset struct {
	Name string
	Path string
	Ext  string
}

func main() {
	flag.Parse()

	tpl, err := template.ParseFS(templateFs, "assets.go.template")
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}

	matches, err := filepath.Glob(filepath.Join(*assetsPath, *assetsGlob))
	if err != nil {
		log.Fatalf("failed to parse glob %v: %v", assetsGlob, err)
	}

	td := TemplateData{}
	td.PackageName = filepath.Base(*assetsPath)
	td.Assets = make([]TemplateAsset, 0)

	for _, m := range matches {
		base := filepath.Base(m)
		td.Assets = append(td.Assets, TemplateAsset{toUpperCamelCase(base), base, filepath.Ext(base)})
	}

	out, err := os.Create(*outPath)
	if err != nil {
		log.Fatalf("failed to open out file: %v", err)
	}
	defer out.Close()

	if err := tpl.ExecuteTemplate(out, "assets.go.template", td); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}
}

func toUpperCamelCase(value string) string {
	name, _ := strings.CutSuffix(value, filepath.Ext(value))
	parts := strings.Split(name, "_")

	var sb strings.Builder

	for _, part := range parts {
		sb.WriteString(strings.ToUpper(string(part[0])))

		if len(part) > 1 {
			sb.WriteString(part[1:])
		}
	}

	return sb.String()
}
