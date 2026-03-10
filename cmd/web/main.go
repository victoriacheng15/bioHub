package main

import (
	"fmt"
	"os"

	"github.com/victoriacheng15/biohub/internal/web"
)

func main() {
	// Build the site
	err := web.BuildSite("config.yml", "internal/web/template", "dist")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building site: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Build complete. Files are in dist/")
}
