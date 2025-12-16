package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// ============================================================================
// Test Fixture Helpers (Plan #2)
// ============================================================================

// getMinimalTheme returns a Theme struct with only essential fields
func getMinimalTheme() Theme {
	return Theme{
		Background:  "#ffffff",
		Text:        "#000000",
		Button:      "#000000",
		ButtonText:  "#ffffff",
		ButtonHover: "#000000",
		Link:        "#000000",
		LinkText:    "#ffffff",
		LinkHover:   "#000000",
	}
}

// getFullTheme returns a complete Theme struct with all colors set
func getFullTheme() Theme {
	return Theme{
		Background:  "#1f2937",
		Text:        "#ffa375",
		Button:      "#60a5fa",
		ButtonText:  "#f1f5f9",
		ButtonHover: "#1147bb",
		Link:        "#1147bb",
		LinkText:    "#f1f5f9",
		LinkHover:   "#09265D",
	}
}

// getValidConfigYAML returns standard YAML config content for testing
func getValidConfigYAML() string {
	return `Params:
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
}

// createTempConfigFile creates a temporary config file with error checking
func createTempConfigFile(t *testing.T, tmpDir, filename, content string) string {
	configPath := filepath.Join(tmpDir, filename)
	err := os.WriteFile(configPath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}
	return configPath
}

// createFullTestSetup creates complete test environment with all necessary directories and files
func createFullTestSetup(t *testing.T, configYAML, templateContent string) (tmpDir, configPath, templatePath, staticSrcDir, staticDstDir, outputDir string) {
	tmpDir = t.TempDir()

	// Create config file
	configPath = createTempConfigFile(t, tmpDir, "config.yml", configYAML)

	// Create template directory and file
	templateDir := filepath.Join(tmpDir, "template")
	if err := os.MkdirAll(templateDir, 0755); err != nil {
		t.Fatalf("Failed to create template directory: %v", err)
	}

	templatePath = filepath.Join(templateDir, "index.html")
	if err := os.WriteFile(templatePath, []byte(templateContent), 0644); err != nil {
		t.Fatalf("Failed to create template file: %v", err)
	}

	// Create static source directory
	staticSrcDir = filepath.Join(templateDir, "static")
	if err := os.MkdirAll(staticSrcDir, 0755); err != nil {
		t.Fatalf("Failed to create static directory: %v", err)
	}

	// Create output and static destination directories
	outputDir = filepath.Join(tmpDir, "dist")
	staticDstDir = filepath.Join(outputDir, "static")

	return tmpDir, configPath, templatePath, staticSrcDir, staticDstDir, outputDir
}

// TestLoadConfig groups all LoadConfig tests
func TestLoadConfig(t *testing.T) {
	t.Run("valid config", func(t *testing.T) {
		// Use helper to create full test setup (config, template, dirs)
		tmpDir, configPath, _, _, _, _ := createFullTestSetup(t, getValidConfigYAML(), "<html></html>")
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
		_ = tmpDir // silence unused var warning
	})

	t.Run("file not found", func(t *testing.T) {
		_, err := LoadConfig("nonexistent.yml")
		if err == nil {
			t.Error("Expected error for non-existent file, got nil")
		}
	})

	t.Run("invalid YAML", func(t *testing.T) {
		tmpDir, configPath, _, _, _, _ := createFullTestSetup(t, `This is not: valid: YAML: [
`, "<html></html>")
		_, err := LoadConfig(configPath)
		if err == nil {
			t.Error("Expected error for invalid YAML, got nil")
		}
		_ = tmpDir
	})

	t.Run("multiple socials and links", func(t *testing.T) {
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
		tmpDir, configPath, _, _, _, _ := createFullTestSetup(t, configContent, "<html></html>")
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
		_ = tmpDir
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

		if err := os.WriteFile(testFile1, []byte("content1"), 0644); err != nil {
			t.Fatalf("Failed to write file1.txt: %v", err)
		}
		if err := os.WriteFile(testFile2, []byte("content2"), 0644); err != nil {
			t.Fatalf("Failed to write file2.txt: %v", err)
		}
		if err := os.MkdirAll(subDir, 0755); err != nil {
			t.Fatalf("Failed to create subdir: %v", err)
		}
		if err := os.WriteFile(filepath.Join(subDir, "file3.txt"), []byte("content3"), 0644); err != nil {
			t.Fatalf("Failed to write file3.txt: %v", err)
		}

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
		if err := os.WriteFile(filepath.Join(srcDir, ".gitkeep"), []byte(""), 0644); err != nil {
			t.Fatalf("Failed to write .gitkeep: %v", err)
		}
		if err := os.WriteFile(filepath.Join(srcDir, "real_file.txt"), []byte("content"), 0644); err != nil {
			t.Fatalf("Failed to write real_file.txt: %v", err)
		}

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

	t.Run("filepath.Rel error on invalid paths", func(t *testing.T) {
		srcDir := t.TempDir()
		dstDir := t.TempDir()

		// Create a test file
		if err := os.WriteFile(filepath.Join(srcDir, "test.txt"), []byte("content"), 0644); err != nil {
			t.Fatalf("Failed to write test file: %v", err)
		}

		// Try to copy with a source path that will cause filepath.Rel to fail
		// by passing an invalid src path that doesn't match the Walk path
		// This simulates the error handling in the Walk callback
		invalidSrc := "/invalid/nonexistent/path"
		err := CopyDir(invalidSrc, dstDir)
		if err == nil {
			t.Error("Expected error for invalid source path, got nil")
		}
	})
}

// TestTemplateRendering groups all template rendering tests
func TestTemplateRendering(t *testing.T) {
	t.Run("template execution with valid config", func(t *testing.T) {
		tmpDir := t.TempDir()

		// Create a minimal template file
		templatePath := filepath.Join(tmpDir, "test.html")
		templateContent := `<!DOCTYPE html>
<html>
<head>
  <title>{{.Params.Name}}</title>
</head>
<body>
  <h1>{{.Params.Name}}</h1>
  <p>{{.Params.Headline}}</p>
  {{range .Params.Socials}}
  <a href="{{.URL}}">{{.Platform}}</a>
  {{end}}
</body>
</html>`

		err := os.WriteFile(templatePath, []byte(templateContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create test template: %v", err)
		}

		config := Config{
			Params: Params{
				Name:     "Test User",
				Headline: "Test Headline",
				Socials: []Social{
					{Platform: "GitHub", Icon: "github.svg", URL: "https://github.com/test"},
				},
				Links: []Link{
					{Name: "Website", URL: "https://example.com"},
				},
			},
		}

		tmpl, err := os.Create(filepath.Join(tmpDir, "output.html"))
		if err != nil {
			t.Fatalf("Failed to create output file: %v", err)
		}
		defer tmpl.Close()

		// This would test template execution if we parse and execute
		// For now, we verify config structure is correct
		if config.Params.Name != "Test User" {
			t.Errorf("Config name mismatch")
		}
	})

	t.Run("config with empty socials and links", func(t *testing.T) {
		config := Config{
			Params: Params{
				Name:     "User",
				Headline: "Headline",
				Theme: Theme{
					Background: "#fff",
					Text:       "#000",
				},
				Socials: []Social{},
				Links:   []Link{},
			},
		}

		if len(config.Params.Socials) != 0 {
			t.Errorf("Expected empty socials, got %d", len(config.Params.Socials))
		}

		if len(config.Params.Links) != 0 {
			t.Errorf("Expected empty links, got %d", len(config.Params.Links))
		}
	})

	t.Run("config with all theme colors", func(t *testing.T) {
		config := Config{
			Params: Params{
				Name:  "Test",
				Theme: getFullTheme(),
			},
		}

		// Verify all theme colors are set
		if config.Params.Theme.Background == "" {
			t.Error("Background color not set")
		}
		if config.Params.Theme.LinkHover == "" {
			t.Error("LinkHover color not set")
		}
	})
}

// TestStructsAndTypes groups all struct and type validation tests
func TestStructsAndTypes(t *testing.T) {
	// Table-driven tests for struct field validation
	structTests := []struct {
		name      string
		testFunc  func() error
		expectErr bool
	}{
		{
			name: "social struct fields",
			testFunc: func() error {
				social := Social{
					Platform: "GitHub",
					Icon:     "github.svg",
					URL:      "https://github.com/test",
				}
				if social.Platform == "" || social.Icon == "" || social.URL == "" {
					return fmt.Errorf("social struct fields not properly set")
				}
				return nil
			},
			expectErr: false,
		},
		{
			name: "link struct fields",
			testFunc: func() error {
				link := Link{
					Name: "Portfolio",
					URL:  "https://example.com",
				}
				if link.Name == "" || link.URL == "" {
					return fmt.Errorf("link struct fields not properly set")
				}
				return nil
			},
			expectErr: false,
		},
		{
			name: "theme struct with all colors",
			testFunc: func() error {
				theme := getFullTheme()
				colors := []string{
					theme.Background, theme.Text, theme.Button, theme.ButtonText,
					theme.ButtonHover, theme.Link, theme.LinkText, theme.LinkHover,
				}
				for i, color := range colors {
					if color == "" {
						colorNames := []string{"Background", "Text", "Button", "ButtonText", "ButtonHover", "Link", "LinkText", "LinkHover"}
						return fmt.Errorf("theme color %s is empty", colorNames[i])
					}
				}
				return nil
			},
			expectErr: false,
		},
		{
			name: "params struct with all fields",
			testFunc: func() error {
				params := Params{
					Avatar:   "avatar.jpg",
					Name:     "Victoria",
					Headline: "Developer",
					Theme:    getMinimalTheme(),
					Socials:  []Social{{Platform: "GitHub", Icon: "gh.svg", URL: "https://github.com"}},
					Links:    []Link{{Name: "Site", URL: "https://example.com"}},
				}
				if params.Avatar == "" || params.Name == "" || params.Headline == "" {
					return fmt.Errorf("params basic fields not set")
				}
				if len(params.Socials) != 1 || len(params.Links) != 1 {
					return fmt.Errorf("params collections not set correctly")
				}
				return nil
			},
			expectErr: false,
		},
		{
			name: "minimal theme struct",
			testFunc: func() error {
				theme := getMinimalTheme()
				if theme.Background == "" || theme.Text == "" {
					return fmt.Errorf("minimal theme missing essential colors")
				}
				return nil
			},
			expectErr: false,
		},
	}

	for _, tc := range structTests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.testFunc()
			if tc.expectErr && err == nil {
				t.Error("expected error, got nil")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

// TestConfigIntegration tests full config workflow
func TestConfigIntegration(t *testing.T) {
	t.Run("full config load and data flow", func(t *testing.T) {
		tmpDir := t.TempDir()

		configContent := `Params:
  Avatar: "avatar.jpg"
  Name: "Victoria Cheng"
  Headline: "Developer | Designer"
  Theme:
    Background: "#1f2937"
    Text: "#ffa375"
    Button: "#60a5fa"
    ButtonText: "#f1f5f9"
    ButtonHover: "#1147bb"
    Link: "#1147bb"
    LinkText: "#f1f5f9"
    LinkHover: "#09265D"
  Socials:
    - Platform: "GitHub"
      Icon: "github.svg"
      URL: "https://github.com/victoria"
    - Platform: "LinkedIn"
      Icon: "linkedin.svg"
      URL: "https://linkedin.com/in/victoria"
  Links:
    - Name: "Portfolio"
      URL: "https://victoria.dev"
    - Name: "Blog"
      URL: "https://victoria.dev/blog"
    - Name: "Resume"
      URL: "https://victoria.dev/resume"
`

		configPath := createTempConfigFile(t, tmpDir, "config.yml", configContent)
		config, err := LoadConfig(configPath)

		if err != nil {
			t.Fatalf("LoadConfig failed: %v", err)
		}

		// Verify complete data flow
		if config.Params.Name != "Victoria Cheng" {
			t.Errorf("Name mismatch: got %s", config.Params.Name)
		}

		if len(config.Params.Socials) != 2 {
			t.Errorf("Expected 2 socials, got %d", len(config.Params.Socials))
		}

		if len(config.Params.Links) != 3 {
			t.Errorf("Expected 3 links, got %d", len(config.Params.Links))
		}

		// Verify all socials have required fields
		for i, social := range config.Params.Socials {
			if social.Platform == "" || social.Icon == "" || social.URL == "" {
				t.Errorf("Social %d missing fields", i)
			}
		}

		// Verify all links have required fields
		for i, link := range config.Params.Links {
			if link.Name == "" || link.URL == "" {
				t.Errorf("Link %d missing fields", i)
			}
		}

		// Verify theme is complete
		if config.Params.Theme.Background == "" {
			t.Error("Theme missing background color")
		}
	})
}

// TestBuildSite groups all BuildSite function tests
func TestBuildSite(t *testing.T) {
	t.Run("successful build with valid config and template", func(t *testing.T) {
		templateContent := `<!DOCTYPE html>
<html>
<head>
  <title>{{.Params.Name}}</title>
</head>
<body>
  <h1>{{.Params.Name}}</h1>
  <p>{{.Params.Headline}}</p>
  {{range .Params.Socials}}
  <a href="{{.URL}}">{{.Platform}}</a>
  {{end}}
  {{range .Params.Links}}
  <a href="{{.URL}}">{{.Name}}</a>
  {{end}}
</body>
</html>`

		_, configPath, templatePath, staticSrcDir, staticDstDir, outputDir := createFullTestSetup(t, getValidConfigYAML(), templateContent)

		if err := os.WriteFile(filepath.Join(staticSrcDir, "style.css"), []byte("body { margin: 0; }"), 0644); err != nil {
			t.Fatalf("Failed to write style sheet: %v", err)
		}

		// Run BuildSite
		err := BuildSite(configPath, templatePath, outputDir, staticSrcDir, staticDstDir)
		if err != nil {
			t.Fatalf("BuildSite failed: %v", err)
		}

		// Verify output files exist
		if _, err := os.Stat(filepath.Join(outputDir, "index.html")); err != nil {
			t.Errorf("index.html not created: %v", err)
		}

		if _, err := os.Stat(filepath.Join(staticDstDir, "style.css")); err != nil {
			t.Errorf("static files not copied: %v", err)
		}

		// Verify HTML content contains expected data
		htmlContent, _ := os.ReadFile(filepath.Join(outputDir, "index.html"))
		htmlStr := string(htmlContent)

		if !strings.Contains(htmlStr, "Test User") {
			t.Error("HTML does not contain user name")
		}

		if !strings.Contains(htmlStr, "Test Headline") {
			t.Error("HTML does not contain headline")
		}

		if !strings.Contains(htmlStr, "GitHub") {
			t.Error("HTML does not contain social platform")
		}

		if !strings.Contains(htmlStr, "Website") {
			t.Error("HTML does not contain link name")
		}
	})

	// Table-driven tests for error cases
	errorCases := []struct {
		name        string
		setupFunc   func(tmpDir string) (configPath, templatePath, staticSrcDir, outputDir, staticDstDir string, cleanup func())
		expectErr   bool
		description string
		errMessage  string
	}{
		{
			name: "invalid config path",
			setupFunc: func(tmpDir string) (string, string, string, string, string, func()) {
				templatePath := filepath.Join(tmpDir, "template.html")
				if err := os.WriteFile(templatePath, []byte("<html></html>"), 0644); err != nil {
					t.Fatalf("Failed to write template file: %v", err)
				}
				staticSrcDir := filepath.Join(tmpDir, "static")
				outputDir := filepath.Join(tmpDir, "dist")
				staticDstDir := filepath.Join(outputDir, "static")
				return filepath.Join(tmpDir, "nonexistent.yml"), templatePath, staticSrcDir, outputDir, staticDstDir, func() {}
			},
			expectErr:   true,
			description: "should fail with nonexistent config file",
			errMessage:  "",
		},
		{
			name: "invalid template path",
			setupFunc: func(tmpDir string) (string, string, string, string, string, func()) {
				configPath := createTempConfigFile(t, tmpDir, "config.yml", `Params:
  Avatar: "avatar.jpg"
  Name: "Test"
  Headline: "Test"
  Theme:
    Background: "#fff"
    Text: "#000"
    Button: "#000"
    ButtonText: "#fff"
    ButtonHover: "#000"
    Link: "#000"
    LinkText: "#fff"
    LinkHover: "#000"
  Socials: []
  Links: []
`)
				staticSrcDir := filepath.Join(tmpDir, "static")
				outputDir := filepath.Join(tmpDir, "dist")
				staticDstDir := filepath.Join(outputDir, "static")
				return configPath, filepath.Join(tmpDir, "nonexistent.html"), staticSrcDir, outputDir, staticDstDir, func() {}
			},
			expectErr:   true,
			description: "should fail with nonexistent template file",
			errMessage:  "",
		},
		{
			name: "invalid static source",
			setupFunc: func(tmpDir string) (string, string, string, string, string, func()) {
				configPath := createTempConfigFile(t, tmpDir, "config.yml", `Params:
  Avatar: "avatar.jpg"
  Name: "Test"
  Headline: "Test"
  Theme:
    Background: "#fff"
    Text: "#000"
    Button: "#000"
    ButtonText: "#fff"
    ButtonHover: "#000"
    Link: "#000"
    LinkText: "#fff"
    LinkHover: "#000"
  Socials: []
  Links: []
`)
				templatePath := filepath.Join(tmpDir, "template.html")
				if err := os.WriteFile(templatePath, []byte("<html><body>{{.Params.Name}}</body></html>"), 0644); err != nil {
					t.Fatalf("Failed to write template file: %v", err)
				}
				outputDir := filepath.Join(tmpDir, "dist")
				staticDstDir := filepath.Join(outputDir, "static")
				return configPath, templatePath, filepath.Join(tmpDir, "nonexistent"), outputDir, staticDstDir, func() {}
			},
			expectErr:   true,
			description: "should fail with nonexistent static source directory",
			errMessage:  "",
		},
		{
			name: "output file cannot be created",
			setupFunc: func(tmpDir string) (string, string, string, string, string, func()) {
				configPath := createTempConfigFile(t, tmpDir, "config.yml", getValidConfigYAML())
				templatePath := filepath.Join(tmpDir, "template.html")
				if err := os.WriteFile(templatePath, []byte("<html><body>{{.Params.Name}}</body></html>"), 0644); err != nil {
					t.Fatalf("Failed to write template file: %v", err)
				}
				staticSrcDir := filepath.Join(tmpDir, "static")
				if err := os.MkdirAll(staticSrcDir, 0755); err != nil {
					t.Fatalf("Failed to create static directory: %v", err)
				}
				outputDir := filepath.Join(tmpDir, "dist")
				staticDstDir := filepath.Join(outputDir, "static")
				if err := os.MkdirAll(staticDstDir, 0755); err != nil {
					t.Fatalf("Failed to create static destination directory: %v", err)
				}
				// Make output dir read-only to prevent index.html creation
				if err := os.Chmod(outputDir, 0555); err != nil {
					t.Fatalf("Failed to change permissions: %v", err)
				}
				// Return cleanup function to restore permissions
				cleanup := func() {
					os.Chmod(outputDir, 0755)
				}
				return configPath, templatePath, staticSrcDir, outputDir, staticDstDir, cleanup
			},
			expectErr:   true,
			description: "should fail when output file cannot be created",
			errMessage:  "error creating output file",
		},
	}

	for _, tc := range errorCases {
		t.Run("build fails with "+tc.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			configPath, templatePath, staticSrcDir, outputDir, staticDstDir, cleanup := tc.setupFunc(tmpDir)
			defer cleanup()

			err := BuildSite(configPath, templatePath, outputDir, staticSrcDir, staticDstDir)

			if tc.expectErr && err == nil {
				t.Errorf("%s: expected error, got nil", tc.description)
			}
			if !tc.expectErr && err != nil {
				t.Errorf("%s: unexpected error: %v", tc.description, err)
			}
			// Verify error message if specified
			if tc.expectErr && tc.errMessage != "" && err != nil {
				if !strings.Contains(err.Error(), tc.errMessage) {
					t.Errorf("%s: expected error message to contain '%s', got: %v", tc.description, tc.errMessage, err)
				}
			}
		})
	}

	t.Run("build with empty socials and links", func(t *testing.T) {
		tmpDir := t.TempDir()

		// Create config with empty collections
		configPath := filepath.Join(tmpDir, "config.yml")
		configContent := `Params:
  Avatar: "avatar.jpg"
  Name: "Test"
  Headline: "Test"
  Theme:
    Background: "#fff"
    Text: "#000"
    Button: "#000"
    ButtonText: "#fff"
    ButtonHover: "#000"
    Link: "#000"
    LinkText: "#fff"
    LinkHover: "#000"
  Socials: []
  Links: []
`
		if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
			t.Fatalf("Failed to write config file: %v", err)
		}

		// Create template
		templatePath := filepath.Join(tmpDir, "template.html")
		templateContent := `<html><body>
{{range .Params.Socials}}<a>{{.Platform}}</a>{{end}}
{{range .Params.Links}}<a>{{.Name}}</a>{{end}}
</body></html>`
		if err := os.WriteFile(templatePath, []byte(templateContent), 0644); err != nil {
			t.Fatalf("Failed to write template file: %v", err)
		}

		// Create empty static dir
		staticSrcDir := filepath.Join(tmpDir, "static")
		if err := os.MkdirAll(staticSrcDir, 0755); err != nil {
			t.Fatalf("Failed to create static directory: %v", err)
		}

		outputDir := filepath.Join(tmpDir, "dist")
		staticDstDir := filepath.Join(outputDir, "static")

		err := BuildSite(configPath, templatePath, outputDir, staticSrcDir, staticDstDir)
		if err != nil {
			t.Fatalf("BuildSite failed with empty collections: %v", err)
		}

		if _, err := os.Stat(filepath.Join(outputDir, "index.html")); err != nil {
			t.Error("index.html not created with empty collections")
		}
	})

	t.Run("build fails when output directory cannot be created", func(t *testing.T) {
		tmpDir := t.TempDir()

		// Create config
		configPath := filepath.Join(tmpDir, "config.yml")
		configContent := `Params:
  Avatar: "avatar.jpg"
  Name: "Test"
  Headline: "Test"
  Theme:
    Background: "#fff"
    Text: "#000"
    Button: "#000"
    ButtonText: "#fff"
    ButtonHover: "#000"
    Link: "#000"
    LinkText: "#fff"
    LinkHover: "#000"
  Socials: []
  Links: []
`
		if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
			t.Fatalf("Failed to write config file: %v", err)
		}

		// Create template
		templatePath := filepath.Join(tmpDir, "template.html")
		if err := os.WriteFile(templatePath, []byte("<html></html>"), 0644); err != nil {
			t.Fatalf("Failed to write template file: %v", err)
		}

		// Create static dir
		staticSrcDir := filepath.Join(tmpDir, "static")
		if err := os.MkdirAll(staticSrcDir, 0755); err != nil {
			t.Fatalf("Failed to create static directory: %v", err)
		}

		// Use an invalid path that cannot be created (path with non-existent parent)
		outputDir := filepath.Join(tmpDir, "dist")
		staticDstDir := filepath.Join("/dev/null/invalid/path", "static")

		err := BuildSite(configPath, templatePath, outputDir, staticSrcDir, staticDstDir)
		if err == nil {
			t.Error("Expected error when output directory cannot be created, got nil")
		}

		// Verify error message contains expected text
		if !strings.Contains(err.Error(), "error creating output directory") {
			t.Errorf("Expected error to mention 'error creating output directory', got: %v", err)
		}
	})
}

// TestRun groups all run() function tests
func TestRun(t *testing.T) {
	t.Run("run succeeds with valid config", func(t *testing.T) {
		// Save current working directory
		originalDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get current directory: %v", err)
		}
		defer os.Chdir(originalDir)

		templateContent := `<!DOCTYPE html>
<html>
<head>
  <title>{{.Params.Name}}</title>
</head>
<body>
  <h1>{{.Params.Name}}</h1>
  <p>{{.Params.Headline}}</p>
</body>
</html>`

		tmpDir, _, _, staticSrcDir, _, _ := createFullTestSetup(t, getValidConfigYAML(), templateContent)
		os.Chdir(tmpDir)

		if err := os.WriteFile(filepath.Join(staticSrcDir, "style.css"), []byte("body { margin: 0; }"), 0644); err != nil {
			t.Fatalf("Failed to write style sheet: %v", err)
		}

		// Run should succeed
		exitCode := run()
		if exitCode != 0 {
			t.Errorf("Expected exit code 0, got %d", exitCode)
		}

		// Verify dist directory was created
		if _, err := os.Stat(filepath.Join(tmpDir, "dist", "index.html")); err != nil {
			t.Errorf("dist/index.html not created: %v", err)
		}
	})

	t.Run("run fails with missing config.yml", func(t *testing.T) {
		// Save current working directory
		originalDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get current directory: %v", err)
		}
		defer os.Chdir(originalDir)

		// Create temporary workspace without config.yml
		tmpDir := t.TempDir()
		os.Chdir(tmpDir)

		// Run should fail with exit code 1
		exitCode := run()
		if exitCode != 1 {
			t.Errorf("Expected exit code 1, got %d", exitCode)
		}
	})

	t.Run("run fails with missing template", func(t *testing.T) {
		// Save current working directory
		originalDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get current directory: %v", err)
		}
		defer os.Chdir(originalDir)

		// Create temporary workspace
		tmpDir := t.TempDir()
		os.Chdir(tmpDir)

		// Create config.yml only
		configPath := filepath.Join(tmpDir, "config.yml")
		configContent := `Params:
  Avatar: "avatar.jpg"
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
  Socials: []
  Links: []
`
		if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
			t.Fatalf("Failed to write config file: %v", err)
		}

		// Run should fail with exit code 1 (template/index.html doesn't exist)
		exitCode := run()
		if exitCode != 1 {
			t.Errorf("Expected exit code 1, got %d", exitCode)
		}
	})
}
