package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/kazweda/go-web-sandbox/templates"
)

func main() {
	outDir := "docs"
	outFile := filepath.Join(outDir, "index.html")

	// ★ ディレクトリを保証
	if err := os.MkdirAll(outDir, 0755); err != nil {
		panic(err)
	}

	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := templates.Page().Render(context.Background(), f); err != nil {
		panic(err)
	}
}
