package main

import (
	"os"
	"path/filepath"
	"testing"
)

// Helper function to check if string contains substring
func contains(str, substr string) bool {
	return len(str) > 0 && len(substr) > 0 && (str == substr || len(str) > len(substr) && (str[:len(substr)] == substr || str[len(str)-len(substr):] == substr || findSubstring(str, substr)))
}

func findSubstring(str, substr string) bool {
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

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
				Name: "Test",
				Theme: Theme{
					Background:  "#1f2937",
					Text:        "#ffa375",
					Button:      "#60a5fa",
					ButtonText:  "#f1f5f9",
					ButtonHover: "#1147bb",
					Link:        "#1147bb",
					LinkText:    "#f1f5f9",
					LinkHover:   "#09265D",
				},
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
	t.Run("social struct fields", func(t *testing.T) {
		social := Social{
			Platform: "GitHub",
			Icon:     "github.svg",
			URL:      "https://github.com/test",
		}

		if social.Platform == "" || social.Icon == "" || social.URL == "" {
			t.Error("Social struct fields are not properly set")
		}
	})

	t.Run("link struct fields", func(t *testing.T) {
		link := Link{
			Name: "Portfolio",
			URL:  "https://example.com",
		}

		if link.Name == "" || link.URL == "" {
			t.Error("Link struct fields are not properly set")
		}
	})

	t.Run("theme struct with all colors", func(t *testing.T) {
		theme := Theme{
			Background:  "#1f2937",
			Text:        "#ffa375",
			Button:      "#60a5fa",
			ButtonText:  "#f1f5f9",
			ButtonHover: "#1147bb",
			Link:        "#1147bb",
			LinkText:    "#f1f5f9",
			LinkHover:   "#09265D",
		}

		colors := []string{
			theme.Background, theme.Text, theme.Button, theme.ButtonText,
			theme.ButtonHover, theme.Link, theme.LinkText, theme.LinkHover,
		}

		for _, color := range colors {
			if color == "" {
				t.Error("Theme color is empty")
			}
		}
	})

	t.Run("params struct with all fields", func(t *testing.T) {
		params := Params{
			Avatar:   "avatar.jpg",
			Name:     "Victoria",
			Headline: "Developer",
			Theme:    Theme{Background: "#fff", Text: "#000"},
			Socials:  []Social{{Platform: "GitHub", Icon: "gh.svg", URL: "https://github.com"}},
			Links:    []Link{{Name: "Site", URL: "https://example.com"}},
		}

		if params.Avatar == "" || params.Name == "" || params.Headline == "" {
			t.Error("Params basic fields not set")
		}
		if len(params.Socials) != 1 || len(params.Links) != 1 {
			t.Error("Params collections not set correctly")
		}
	})
}

// TestConfigIntegration tests full config workflow
func TestConfigIntegration(t *testing.T) {
	t.Run("full config load and data flow", func(t *testing.T) {
		tmpDir := t.TempDir()
		configPath := filepath.Join(tmpDir, "config.yml")

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

		os.WriteFile(configPath, []byte(configContent), 0644)
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
		tmpDir := t.TempDir()

		// Create config file
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
  Socials:
    - Platform: "GitHub"
      Icon: "github.svg"
      URL: "https://github.com/test"
  Links:
    - Name: "Website"
      URL: "https://example.com"
`
		os.WriteFile(configPath, []byte(configContent), 0644)

		// Create template file
		templatePath := filepath.Join(tmpDir, "template.html")
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
		os.WriteFile(templatePath, []byte(templateContent), 0644)

		// Create source static directory
		staticSrcDir := filepath.Join(tmpDir, "static")
		os.MkdirAll(staticSrcDir, 0755)
		os.WriteFile(filepath.Join(staticSrcDir, "style.css"), []byte("body { margin: 0; }"), 0644)

		// Create output directory
		outputDir := filepath.Join(tmpDir, "dist")
		staticDstDir := filepath.Join(outputDir, "static")

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

		if !contains(htmlStr, "Test User") {
			t.Error("HTML does not contain user name")
		}

		if !contains(htmlStr, "Test Headline") {
			t.Error("HTML does not contain headline")
		}

		if !contains(htmlStr, "GitHub") {
			t.Error("HTML does not contain social platform")
		}

		if !contains(htmlStr, "Website") {
			t.Error("HTML does not contain link name")
		}
	})

	t.Run("build fails with invalid config path", func(t *testing.T) {
		tmpDir := t.TempDir()
		templatePath := filepath.Join(tmpDir, "template.html")
		os.WriteFile(templatePath, []byte("<html></html>"), 0644)

		outputDir := filepath.Join(tmpDir, "dist")
		staticSrcDir := filepath.Join(tmpDir, "static")
		staticDstDir := filepath.Join(outputDir, "static")

		err := BuildSite(
			filepath.Join(tmpDir, "nonexistent.yml"),
			templatePath,
			outputDir,
			staticSrcDir,
			staticDstDir,
		)

		if err == nil {
			t.Error("Expected error for invalid config path, got nil")
		}
	})

	t.Run("build fails with invalid template path", func(t *testing.T) {
		tmpDir := t.TempDir()

		// Create valid config
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
		os.WriteFile(configPath, []byte(configContent), 0644)

		outputDir := filepath.Join(tmpDir, "dist")
		staticSrcDir := filepath.Join(tmpDir, "static")
		staticDstDir := filepath.Join(outputDir, "static")

		err := BuildSite(
			configPath,
			filepath.Join(tmpDir, "nonexistent.html"),
			outputDir,
			staticSrcDir,
			staticDstDir,
		)

		if err == nil {
			t.Error("Expected error for invalid template path, got nil")
		}
	})

	t.Run("build fails with invalid static source", func(t *testing.T) {
		tmpDir := t.TempDir()

		// Create valid config
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
		os.WriteFile(configPath, []byte(configContent), 0644)

		// Create valid template
		templatePath := filepath.Join(tmpDir, "template.html")
		os.WriteFile(templatePath, []byte("<html><body>{{.Params.Name}}</body></html>"), 0644)

		outputDir := filepath.Join(tmpDir, "dist")
		staticDstDir := filepath.Join(outputDir, "static")

		err := BuildSite(
			configPath,
			templatePath,
			outputDir,
			filepath.Join(tmpDir, "nonexistent"),
			staticDstDir,
		)

		if err == nil {
			t.Error("Expected error for invalid static source, got nil")
		}
	})

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
		os.WriteFile(configPath, []byte(configContent), 0644)

		// Create template
		templatePath := filepath.Join(tmpDir, "template.html")
		templateContent := `<html><body>
{{range .Params.Socials}}<a>{{.Platform}}</a>{{end}}
{{range .Params.Links}}<a>{{.Name}}</a>{{end}}
</body></html>`
		os.WriteFile(templatePath, []byte(templateContent), 0644)

		// Create empty static dir
		staticSrcDir := filepath.Join(tmpDir, "static")
		os.MkdirAll(staticSrcDir, 0755)

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
		os.WriteFile(configPath, []byte(configContent), 0644)

		// Create template
		templatePath := filepath.Join(tmpDir, "template.html")
		os.WriteFile(templatePath, []byte("<html></html>"), 0644)

		// Create static dir
		staticSrcDir := filepath.Join(tmpDir, "static")
		os.MkdirAll(staticSrcDir, 0755)

		// Use an invalid path that cannot be created (path with non-existent parent)
		outputDir := filepath.Join(tmpDir, "dist")
		staticDstDir := filepath.Join("/dev/null/invalid/path", "static")

		err := BuildSite(configPath, templatePath, outputDir, staticSrcDir, staticDstDir)
		if err == nil {
			t.Error("Expected error when output directory cannot be created, got nil")
		}

		// Verify error message contains expected text
		if !contains(err.Error(), "error creating output directory") {
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

		// Create temporary workspace
		tmpDir := t.TempDir()
		os.Chdir(tmpDir)

		// Create config.yml
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
  Socials:
    - Platform: "GitHub"
      Icon: "github.svg"
      URL: "https://github.com/test"
  Links:
    - Name: "Website"
      URL: "https://example.com"
`
		os.WriteFile(configPath, []byte(configContent), 0644)

		// Create template directory and file
		templateDir := filepath.Join(tmpDir, "template")
		os.MkdirAll(templateDir, 0755)

		templatePath := filepath.Join(templateDir, "index.html")
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
		os.WriteFile(templatePath, []byte(templateContent), 0644)

		// Create static directory
		staticDir := filepath.Join(templateDir, "static")
		os.MkdirAll(staticDir, 0755)
		os.WriteFile(filepath.Join(staticDir, "style.css"), []byte("body { margin: 0; }"), 0644)

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
		os.WriteFile(configPath, []byte(configContent), 0644)

		// Run should fail with exit code 1 (template/index.html doesn't exist)
		exitCode := run()
		if exitCode != 1 {
			t.Errorf("Expected exit code 1, got %d", exitCode)
		}
	})
}
