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

func BuildSite(configPath, templatePath, outputDir, staticSrcDir, staticDstDir string) error {
	// Create output directories
	if err := os.MkdirAll(staticDstDir, 0755); err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	// Load configuration from config.yml
	config, err := LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Create output HTML file
	out, err := os.Create(filepath.Join(outputDir, "index.html"))
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer out.Close()

	// Render template with config data
	if err := tmpl.Execute(out, config); err != nil {
		return fmt.Errorf("error rendering template: %w", err)
	}

	// Copy static files
	if err := CopyDir(staticSrcDir, staticDstDir); err != nil {
		return fmt.Errorf("error copying static files: %w", err)
	}

	return nil
}

func run() int {
	// Load configuration to display debug info
	config, err := LoadConfig("config.yml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config.yml: %v\n", err)
		return 1
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

	// Build the site
	err = BuildSite("config.yml", "template/index.html", "dist", "template/static", "dist/static")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building site: %v\n", err)
		return 1
	}

	fmt.Println("Build complete. Files are in dist/")
	return 0
}

func main() {
	os.Exit(run())
}
