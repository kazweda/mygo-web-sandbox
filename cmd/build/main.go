package main

import (
	"context"
	"os"

	"github.com/kazweda/go-web-sandbox/templates"
)

func main() {
	f, err := os.Create("docs/index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = templates.Page().Render(context.Background(), f)
	if err != nil {
		panic(err)
	}
}
