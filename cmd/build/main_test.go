package main

import (
	"os"
	"path/filepath"
	"testing"
)

// TestLoadConfig groups all LoadConfig tests
func TestLoadConfig(t *testing.T) {
	t.Run("valid config", func(t *testing.T) {
		// Create a temporary config file
		tmpDir := t.TempDir()
		configPath := filepath.Join(tmpDir, "config.yml")

		configContent := `Params:
  Avatar: "static/avatar.jpg"
  Name: "Test User"
  Headline: "Test Headline"
  Theme:
    Background: "#1f2937"
    Text: "#ffffff"
    Button: "#60a5fa"
    ButtonText: "#f1f5f9"
    ButtonHover: "#1147bb"
    Link: "#1147bb"
    LinkText: "#f1f5f9"
    LinkHover: "#09265D"
  Socials:
    - Platform: "GitHub"
      Icon: "static/icons/github.svg"
      URL: "https://github.com/test"
  Links:
    - Name: "Website"
      URL: "https://example.com"
`

		err := os.WriteFile(configPath, []byte(configContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create test config file: %v", err)
		}

		config, err := LoadConfig(configPath)
		if err != nil {
			t.Fatalf("LoadConfig failed: %v", err)
		}

		// Verify loaded values
		if config.Params.Name != "Test User" {
			t.Errorf("Expected Name 'Test User', got '%s'", config.Params.Name)
		}

		if config.Params.Headline != "Test Headline" {
			t.Errorf("Expected Headline 'Test Headline', got '%s'", config.Params.Headline)
		}

		if config.Params.Theme.Background != "#1f2937" {
			t.Errorf("Expected Background '#1f2937', got '%s'", config.Params.Theme.Background)
		}

		if len(config.Params.Socials) != 1 {
			t.Errorf("Expected 1 Social, got %d", len(config.Params.Socials))
		}

		if config.Params.Socials[0].Platform != "GitHub" {
			t.Errorf("Expected Platform 'GitHub', got '%s'", config.Params.Socials[0].Platform)
		}

		if len(config.Params.Links) != 1 {
			t.Errorf("Expected 1 Link, got %d", len(config.Params.Links))
		}
	})

	t.Run("file not found", func(t *testing.T) {
		_, err := LoadConfig("nonexistent.yml")
		if err == nil {
			t.Error("Expected error for non-existent file, got nil")
		}
	})

	t.Run("invalid YAML", func(t *testing.T) {
		tmpDir := t.TempDir()
		configPath := filepath.Join(tmpDir, "invalid.yml")

		invalidContent := `This is not: valid: YAML: [
`

		err := os.WriteFile(configPath, []byte(invalidContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}

		_, err = LoadConfig(configPath)
		if err == nil {
			t.Error("Expected error for invalid YAML, got nil")
		}
	})

	t.Run("multiple socials and links", func(t *testing.T) {
		tmpDir := t.TempDir()
		configPath := filepath.Join(tmpDir, "config.yml")

		configContent := `Params:
  Avatar: "test.jpg"
  Name: "Victoria"
  Headline: "Developer"
  Theme:
    Background: "#000000"
    Text: "#ffffff"
    Button: "#0000ff"
    ButtonText: "#ffffff"
    ButtonHover: "#0000cc"
    Link: "#0000ff"
    LinkText: "#ffffff"
    LinkHover: "#0000cc"
  Socials:
    - Platform: "GitHub"
      Icon: "github.svg"
      URL: "https://github.com/test"
    - Platform: "LinkedIn"
      Icon: "linkedin.svg"
      URL: "https://linkedin.com/in/test"
  Links:
    - Name: "Portfolio"
      URL: "https://example.com"
    - Name: "Blog"
      URL: "https://blog.example.com"
`

		os.WriteFile(configPath, []byte(configContent), 0644)
		config, _ := LoadConfig(configPath)

		// Test multiple socials
		if len(config.Params.Socials) != 2 {
			t.Errorf("Expected 2 socials, got %d", len(config.Params.Socials))
		}

		if config.Params.Socials[1].Platform != "LinkedIn" {
			t.Errorf("Expected second social to be LinkedIn, got %s", config.Params.Socials[1].Platform)
		}

		// Test multiple links
		if len(config.Params.Links) != 2 {
			t.Errorf("Expected 2 links, got %d", len(config.Params.Links))
		}

		if config.Params.Links[1].Name != "Blog" {
			t.Errorf("Expected second link to be Blog, got %s", config.Params.Links[1].Name)
		}
	})
}

// TestCopyDir groups all CopyDir tests
func TestCopyDir(t *testing.T) {
	t.Run("successful directory copy", func(t *testing.T) {
		// Create source directory with files
		srcDir := t.TempDir()
		dstDir := t.TempDir()

		// Create test files
		testFile1 := filepath.Join(srcDir, "file1.txt")
		testFile2 := filepath.Join(srcDir, "file2.txt")
		subDir := filepath.Join(srcDir, "subdir")

		os.WriteFile(testFile1, []byte("content1"), 0644)
		os.WriteFile(testFile2, []byte("content2"), 0644)
		os.MkdirAll(subDir, 0755)
		os.WriteFile(filepath.Join(subDir, "file3.txt"), []byte("content3"), 0644)

		// Copy directory
		err := CopyDir(srcDir, dstDir)
		if err != nil {
			t.Fatalf("CopyDir failed: %v", err)
		}

		// Verify files were copied
		if _, err := os.Stat(filepath.Join(dstDir, "file1.txt")); err != nil {
			t.Errorf("file1.txt not copied: %v", err)
		}

		if _, err := os.Stat(filepath.Join(dstDir, "file2.txt")); err != nil {
			t.Errorf("file2.txt not copied: %v", err)
		}

		if _, err := os.Stat(filepath.Join(dstDir, "subdir", "file3.txt")); err != nil {
			t.Errorf("subdir/file3.txt not copied: %v", err)
		}

		// Verify content
		content, _ := os.ReadFile(filepath.Join(dstDir, "file1.txt"))
		if string(content) != "content1" {
			t.Errorf("file1.txt content mismatch: got '%s'", string(content))
		}
	})

	t.Run("skips .gitkeep files", func(t *testing.T) {
		srcDir := t.TempDir()
		dstDir := t.TempDir()

		// Create .gitkeep file
		os.WriteFile(filepath.Join(srcDir, ".gitkeep"), []byte(""), 0644)
		os.WriteFile(filepath.Join(srcDir, "real_file.txt"), []byte("content"), 0644)

		err := CopyDir(srcDir, dstDir)
		if err != nil {
			t.Fatalf("CopyDir failed: %v", err)
		}

		// Verify .gitkeep was not copied
		if _, err := os.Stat(filepath.Join(dstDir, ".gitkeep")); err == nil {
			t.Error(".gitkeep should not be copied")
		}

		// Verify real file was copied
		if _, err := os.Stat(filepath.Join(dstDir, "real_file.txt")); err != nil {
			t.Errorf("real_file.txt should be copied: %v", err)
		}
	})

	t.Run("source not found", func(t *testing.T) {
		dstDir := t.TempDir()

		err := CopyDir("/nonexistent/source", dstDir)
		if err == nil {
			t.Error("Expected error for non-existent source, got nil")
		}
	})
}
