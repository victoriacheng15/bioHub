package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Theme struct {
	Background  string `yaml:"Background"`
	Text        string `yaml:"Text"`
	Button      string `yaml:"Button"`
	ButtonText  string `yaml:"ButtonText"`
	ButtonHover string `yaml:"ButtonHover"`
	Link        string `yaml:"Link"`
	LinkText    string `yaml:"LinkText"`
	LinkHover   string `yaml:"LinkHover"`
}

type Social struct {
	Platform string `yaml:"Platform"`
	Icon     string `yaml:"Icon"`
	URL      string `yaml:"URL"`
}

type Link struct {
	Name string `yaml:"Name"`
	URL  string `yaml:"URL"`
}

type Params struct {
	Avatar   string   `yaml:"Avatar"`
	Name     string   `yaml:"Name"`
	Headline string   `yaml:"Headline"`
	Theme    Theme    `yaml:"Theme"`
	Socials  []Social `yaml:"Socials"`
	Links    []Link   `yaml:"Links"`
}

type Config struct {
	Params Params `yaml:"Params"`
}

// LoadConfig reads a YAML configuration file and unmarshals it into a Config struct
func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error parsing config file: %w", err)
	}

	return config, nil
}

// CopyDir recursively copies the contents of the directory at src to dst
func CopyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		// Skip .gitkeep files
		if info.Name() == ".gitkeep" {
			return nil
		}

		destPath := filepath.Join(dst, relPath)
		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()

		out, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, in)
		return err
	})
}

func main() {
	// Create output directories
	if err := os.MkdirAll("dist/static", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating dist/static: %v\n", err)
		os.Exit(1)
	}

	// Load configuration from config.yml
	config, err := LoadConfig("config.yml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config.yml: %v\n", err)
		os.Exit(1)
	}

	// Debug: Print loaded config
	fmt.Printf("Loaded config:\n")
	fmt.Printf("  Name: %s\n", config.Params.Name)
	fmt.Printf("  Headline: %s\n", config.Params.Headline)
	fmt.Printf("  Avatar: %s\n", config.Params.Avatar)
	fmt.Printf("  Theme Background: %s\n", config.Params.Theme.Background)
	fmt.Printf("  Socials: %d\n", len(config.Params.Socials))
	fmt.Printf("  Links: %d\n", len(config.Params.Links))
	fmt.Println()

	// Parse the HTML template from template/index.html
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing template: %v\n", err)
		os.Exit(1)
	}

	// Create output HTML file in dist/
	out, err := os.Create("dist/index.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating dist/index.html: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	// Render template with config data and write to dist/index.html
	if err := tmpl.Execute(out, config); err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering template: %v\n", err)
		os.Exit(1)
	}

	// Copy entire template folder (including static assets) to dist/
	err = CopyDir("template/static", "dist/static")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error copying template files: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Build complete. Files are in dist/")
}
