# bioHub

Centralize your social links, personal site, and other key resources with this Go-powered static site generator. Easy to configure, theme, and deploy to GitHub Pages.

## Tech Stack

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![YAML](https://img.shields.io/badge/YAML-000000?style=for-the-badge&logo=yaml&logoColor=white)
![GitHub Pages](https://img.shields.io/badge/GitHub%20Pages-222222?style=for-the-badge&logo=github-pages&logoColor=white)

## Features

- Simple YAML configuration for your profile, links, and theme
- Generates a static HTML site using Go templates
- Fully responsive, accessible, and lightweight
- One-command build and automatic GitHub Pages deployment

## Usage

- **Edit your profile and links**:  
  Update `config.yml` with your name, headline, avatar, social links, icons, and custom theme colors. The `Links` section, for example, could look like this:

  ```yaml
  Links:
    - Name: "🐧 Personal Site"
      URL: "https://example.com/"
    - Name: "👋 About Me"
      URL: "https://example.com/about"
    - Name: "⚡ Life Lately"
      URL: "https://example.com/now"
  ```

- **Build the site**:  
  Ensure you have Go installed, then run:

  ```sh
  make build
  ```
  
  This generates your static site in the `dist/` folder.

- **Preview locally**:  
  Open the `dist/` folder in VS Code and use the [Live Server extension](https://marketplace.visualstudio.com/items?itemName=ritwickdey.LiveServer) to preview your site with live reload. You can:
  - Click **"Go Live"** in the status bar
  - Right-click an HTML file in the Explorer and select **"Open with Live Server"**
  - Use the keyboard shortcut:  
    - Windows/Linux: `Alt+L, Alt+O` to start; `Alt+L, Alt+C` to stop  
    - macOS: `Cmd+L, Cmd+O` to start; `Cmd+L, Cmd+C` to stop  
  - Open the Command Palette (`F1` or `Ctrl+Shift+P`) and run **"Live Server: Open With Live Server"**

- **Deploy**:  
  Push to the `main` branch. A GitHub Action automatically deploys your site to GitHub Pages.

## Folder Structure

```text
bioHub/
├── cmd/
│   └── web/
│       └── main.go         # Site entry point
├── internal/
│   └── web/                # Core domain logic
│       ├── config.go       # YAML config parsing
│       ├── generator.go    # Site generation logic
│       └── template/       # Site source templates
│           ├── index.html
│           └── static/
├── config.yml              # Profile and links configuration
├── dist/                   # Generated static site (do not edit)
├── Makefile                # Build and development commands
└── .github/workflows/      # CI/CD and automation workflows
```
