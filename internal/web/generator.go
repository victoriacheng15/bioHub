package web

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	textTemplate "text/template"
)

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

// BuildSite orchestrates the site generation process
func BuildSite(configPath, templateDir, outputDir string) error {
	// Load configuration from config.yml
	config, err := LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	// Generate Index HTML
	if err := generateIndex(config, templateDir, outputDir); err != nil {
		return err
	}

	// Generate llms.txt
	if err := generateLLMS(config, templateDir, outputDir); err != nil {
		return err
	}

	// Generate robots.txt
	if err := generateRobots(config, templateDir, outputDir); err != nil {
		return err
	}

	// Copy static files
	if err := copyStatic(templateDir, outputDir); err != nil {
		return err
	}

	return nil
}

func generateIndex(config Config, templateDir, outputDir string) error {
	tmpl, err := template.ParseFiles(filepath.Join(templateDir, "index.html"))
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	out, err := os.Create(filepath.Join(outputDir, "index.html"))
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer out.Close()

	if err := tmpl.Execute(out, config); err != nil {
		return fmt.Errorf("error rendering template: %w", err)
	}
	return nil
}

func generateLLMS(config Config, templateDir, outputDir string) error {
	tmpl, err := textTemplate.ParseFiles(filepath.Join(templateDir, "llms.txt"))
	if err != nil {
		return fmt.Errorf("error parsing llms.txt template: %w", err)
	}

	out, err := os.Create(filepath.Join(outputDir, "llms.txt"))
	if err != nil {
		return fmt.Errorf("error creating llms.txt file: %w", err)
	}
	defer out.Close()

	if err := tmpl.Execute(out, config); err != nil {
		return fmt.Errorf("error rendering llms.txt: %w", err)
	}
	return nil
}

func generateRobots(config Config, templateDir, outputDir string) error {
	tmpl, err := textTemplate.ParseFiles(filepath.Join(templateDir, "robots.txt"))
	if err != nil {
		return fmt.Errorf("error parsing robots.txt template: %w", err)
	}

	out, err := os.Create(filepath.Join(outputDir, "robots.txt"))
	if err != nil {
		return fmt.Errorf("error creating robots.txt file: %w", err)
	}
	defer out.Close()

	if err := tmpl.Execute(out, config); err != nil {
		return fmt.Errorf("error rendering robots.txt: %w", err)
	}
	return nil
}

func copyStatic(templateDir, outputDir string) error {
	staticSrcDir := filepath.Join(templateDir, "static")
	staticDstDir := filepath.Join(outputDir, "static")

	if err := os.MkdirAll(staticDstDir, 0755); err != nil {
		return fmt.Errorf("error creating static destination directory: %w", err)
	}

	if err := CopyDir(staticSrcDir, staticDstDir); err != nil {
		return fmt.Errorf("error copying static files: %w", err)
	}
	return nil
}
