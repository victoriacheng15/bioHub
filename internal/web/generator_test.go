package web

import (
	"os"
	"path/filepath"
	"testing"
)

// getValidConfigYAML returns standard YAML config content for testing
func getValidConfigYAML() string {
	return `Params:
  Avatar: "static/avatar.jpg"
  Name: "Test User"
  Headline: "Test Headline"
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
func createFullTestSetup(t *testing.T, configYAML, templateContent string) (tmpDir, configPath, templateDir, outputDir string) {
	tmpDir = t.TempDir()

	// Create config file
	configPath = createTempConfigFile(t, tmpDir, "config.yml", configYAML)

	// Create template directory and file
	templateDir = filepath.Join(tmpDir, "template")
	if err := os.MkdirAll(templateDir, 0755); err != nil {
		t.Fatalf("Failed to create template directory: %v", err)
	}

	templatePath := filepath.Join(templateDir, "index.html")
	if err := os.WriteFile(templatePath, []byte(templateContent), 0644); err != nil {
		t.Fatalf("Failed to create template file: %v", err)
	}

	if err := os.WriteFile(filepath.Join(templateDir, "llms.txt"), []byte("llms"), 0644); err != nil {
		t.Fatalf("Failed to create llms.txt template file: %v", err)
	}

	if err := os.WriteFile(filepath.Join(templateDir, "robots.txt"), []byte("robots"), 0644); err != nil {
		t.Fatalf("Failed to create robots.txt template file: %v", err)
	}

	// Create static source directory
	staticSrcDir := filepath.Join(templateDir, "static")
	if err := os.MkdirAll(staticSrcDir, 0755); err != nil {
		t.Fatalf("Failed to create static directory: %v", err)
	}

	// Create output directory
	outputDir = filepath.Join(tmpDir, "dist")

	return tmpDir, configPath, templateDir, outputDir
}

func TestLoadConfig(t *testing.T) {
	tmpDir := t.TempDir()
	tests := []struct {
		name       string
		configPath string
		content    string
		wantErr    bool
	}{
		{
			name:       "valid config",
			configPath: filepath.Join(tmpDir, "valid.yml"),
			content:    getValidConfigYAML(),
			wantErr:    false,
		},
		{
			name:       "file not found",
			configPath: "nonexistent.yml",
			content:    "",
			wantErr:    true,
		},
		{
			name:       "invalid YAML",
			configPath: filepath.Join(tmpDir, "invalid.yml"),
			content:    "invalid: yaml: {",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.content != "" {
				err := os.WriteFile(tt.configPath, []byte(tt.content), 0644)
				if err != nil {
					t.Fatalf("Failed to write test config: %v", err)
				}
			}

			config, err := LoadConfig(tt.configPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && config.Params.Name != "Test User" {
				t.Errorf("Expected Name 'Test User', got '%s'", config.Params.Name)
			}
		})
	}
}

func TestCopyDir(t *testing.T) {
	tests := []struct {
		name      string
		setupFunc func(t *testing.T) (src, dst string)
		wantErr   bool
	}{
		{
			name: "successful directory copy",
			setupFunc: func(t *testing.T) (string, string) {
				src := t.TempDir()
				dst := t.TempDir()
				testFile := filepath.Join(src, "file1.txt")
				if err := os.WriteFile(testFile, []byte("content1"), 0644); err != nil {
					t.Fatalf("Failed to write test file: %v", err)
				}
				return src, dst
			},
			wantErr: false,
		},
		{
			name: "source not found",
			setupFunc: func(t *testing.T) (string, string) {
				return "nonexistent_src", t.TempDir()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src, dst := tt.setupFunc(t)
			err := CopyDir(src, dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("CopyDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuildSite(t *testing.T) {
	tests := []struct {
		name      string
		setupFunc func(t *testing.T) (configPath, templateDir, outputDir string)
		wantErr   bool
	}{
		{
			name: "successful build",
			setupFunc: func(t *testing.T) (string, string, string) {
				templateContent := `<html><body>{{.Params.Name}}</body></html>`
				_, configPath, templateDir, outputDir := createFullTestSetup(t, getValidConfigYAML(), templateContent)
				staticSrcDir := filepath.Join(templateDir, "static")
				if err := os.WriteFile(filepath.Join(staticSrcDir, "style.css"), []byte("body {}"), 0644); err != nil {
					t.Fatalf("Failed to write style sheet: %v", err)
				}
				return configPath, templateDir, outputDir
			},
			wantErr: false,
		},
		{
			name: "invalid config",
			setupFunc: func(t *testing.T) (string, string, string) {
				tmpDir := t.TempDir()
				return filepath.Join(tmpDir, "nonexistent.yml"), tmpDir, filepath.Join(tmpDir, "dist")
			},
			wantErr: true,
		},
		{
			name: "invalid template path",
			setupFunc: func(t *testing.T) (string, string, string) {
				tmpDir := t.TempDir()
				configPath := createTempConfigFile(t, tmpDir, "config.yml", getValidConfigYAML())
				return configPath, tmpDir, filepath.Join(tmpDir, "dist")
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configPath, templateDir, outputDir := tt.setupFunc(t)
			err := BuildSite(configPath, templateDir, outputDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildSite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerators(t *testing.T) {
	config := Config{Params: Params{Name: "Test"}}

	tests := []struct {
		name    string
		genFunc func(t *testing.T) error
		wantErr bool
	}{
		{
			name: "generateIndex fail creating dir",
			genFunc: func(t *testing.T) error {
				tmpDir := t.TempDir()
				filePath := filepath.Join(tmpDir, "file")
				os.WriteFile(filePath, []byte(""), 0644)
				return generateIndex(config, tmpDir, filepath.Join(filePath, "subdir"))
			},
			wantErr: true,
		},
		{
			name: "generateLLMS fail parsing",
			genFunc: func(t *testing.T) error {
				tmpDir := t.TempDir()
				return generateLLMS(config, tmpDir, filepath.Join(tmpDir, "dist"))
			},
			wantErr: true,
		},
		{
			name: "generateRobots fail parsing",
			genFunc: func(t *testing.T) error {
				tmpDir := t.TempDir()
				return generateRobots(config, tmpDir, filepath.Join(tmpDir, "dist"))
			},
			wantErr: true,
		},
		{
			name: "copyStatic fail",
			genFunc: func(t *testing.T) error {
				tmpDir := t.TempDir()
				return copyStatic(tmpDir, filepath.Join(tmpDir, "dist"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.genFunc(t)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}
