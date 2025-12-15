# Migration Plan: Vue.js (bioHub) to Go with YAML Config

## 1. Project Structure

Recommended Go layout:

```
go-biohub/
  config.yml
  template/
    index.html   # HTML template for rendering
    css/
      styles.css # CSS styles for the site
  static/        # For icons and other assets
  cmd/
    build-bio/
      build.go    # Reads config.yml and generates HTML from template/index.html
```

Notes:

- Keep template files under `template/` so HTML/CSS assets live together.
- `cmd/build-bio/build.go` should drive static site generation; add a `main.go` if you later need a dev server.
- Output the generated HTML to `dist/` (or another consistent folder) for publishing.

## 2. Analyze Existing Features

- Review the Vue components (`Header`, `Footer`, `LinkContainer`, `SocialList`) and document the data each one consumes.
- Identify all configurable parameters from the current app (avatar, name, headline, theme colors, socials, links).

## 3. Define the YAML Schema

Mirror the provided example so the Go config loader can parse:

```yaml
Params:
  Avatar: ...
  Name: ...
  Headline: ...
  Theme:
    Background: ...
    Text: ...
    Button: ...
    ButtonText: ...
    ButtonHover: ...
  Socials:
    - Icon: ...
      URL: ...
  Links:
    - Name: ...
      URL: ...
```

## 4. Go Project Setup

- Run `go mod init github.com/<you>/biohub`.
- Add `gopkg.in/yaml.v3` (and any other libraries you need).

## 5. Define Go Structs for Config

Create structs that map to the YAML structure so you can unmarshal cleanly.

## 6. Load and Parse the YAML

- Read `config.yml` from disk.
- Use `yaml.Unmarshal` to populate your structs.
- Return a `Config` struct that exposes `Params`.

## 7. Template Rendering

- Keep your HTML template in `template/index.html`.
- Use Go’s `html/template` package to apply `Config` data to the document.
- Serve the rendered page via `net/http` or write it directly to `dist/index.html`.

## 8. CSS Styling

- Store the provided CSS in `template/css/styles.css`.
- Reference it from the template with `<link rel="stylesheet" href="static/css/styles.css">`.
- Serve static files with something like `http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))` so CSS and assets are available.

CSS snippet for reference:

```css
*,
*::before,
*::after {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

ul {
  list-style: none;
}

a {
  text-decoration: none;
  color: inherit;
}

body {
  font-family: 'Open Sans', sans-serif;
}

#app {
  min-height: 100vh;
  width: min(90%, 600px);
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2.5rem;
}

.flex {
  display: flex;
}

.flex-col {
  flex-direction: column;
}

.flex-gap {
  gap: 1.5rem;
}

.radius-6 {
  border-radius: 6px;
}

.font-2 {
  font-size: 2rem;
}

.font-1-5 {
  font-size: 1.5rem;
}

header {
  padding: 3rem 0;
  margin: 0 auto;
  justify-content: center;
  align-items: center;
}

header img {
  border-radius: 50%;
  width: 150px;
  aspect-ratio: 1;
}

.social-links .btn {
  padding: 0.1rem;
  align-items: center;
  justify-content: center;
  width: auto;
  height: auto;
}

main {
  width: 100%;
}

.links .btn {
  display: block;
  padding: 0.75rem 1rem;
  text-align: center;
  font-weight: bold;
  letter-spacing: 0.2rem;
}

footer {
  margin-top: auto;
  padding-block: 1rem;
}
```

## 9. Static Assets

- Copy existing icons and any other media into `static/` so they can be referenced by the template.

## 10. Environment, Build & Makefile

- Document how to run `go build` and `go run`.
- Add a Makefile to simplify commands:
  - `make build` builds the binary from `cmd/build-bio/main.go`.
  - `make run` builds and runs the binary.
  - `make clean` removes the build artifact.

## 11. Testing & Validation

- Run the builder with various `config.yml` variations to ensure theme, socials, and links render correctly.
- Verify the generated HTML/CSS match the original styling requirements.

## 12. GitHub Actions: Go Format Check

Include a PR workflow that runs `gofmt -l` across all `.go` files and fails if formatting is needed. Keep this purely as a format check.

## 13. GitHub Actions: Deploy to Pages

Set up a workflow triggered on pushes to `main` that:

- Checks out the repo and sets up Go 1.22.
- Runs `go run main.go build` (or the equivalent build command) to populate `dist/`.
- Uploads the `dist/` folder via `actions/upload-pages-artifact@v3`.
- Deploys with `actions/deploy-pages@v4`.

## 14. Documentation

- Update `README.md` with usage instructions, how to edit `config.yml`, how to run the builder, and how deployment works.

---

## 15. Complete Migration: Mapping Vue.js Structure to Go

This section provides a step-by-step mapping and migration guide from the current Vue.js-based bioHub to the new Go-based static site generator, referencing the actual repo structure and Go implementation details.

### 15.1. Map Current Structure to Go Layout

| Vue.js Repo Path         | Go Static Site Path         | Migration Action/Notes                                  |
|-------------------------|-----------------------------|--------------------------------------------------------|
| `public/`               | `static/`                   | Copy all static assets (images, icons) here             |
| `public/icons/`         | `static/icons/`             | Copy all SVGs/icons here                                |
| `src/index.css`         | `template/css/styles.css`   | Move and adapt CSS for Go template                      |
| `src/App.vue`           | `template/index.html`       | Convert main Vue template to Go HTML template           |
| `src/components/`       | `template/index.html`       | Integrate component logic into Go template              |
| `src/constant/index.ts` | `config.yml`                | Convert constants/data to YAML config                   |
| `index.html` (root)     | `template/index.html`       | Use as reference for Go template                        |
| `README.md`             | `README.md`                 | Update with Go usage and migration notes                |
| `package.json`, `tsconfig.json`, etc. | *(remove)*     | No longer needed; replaced by Go tooling                |

#### Before and After Folder Structure

**Before (Vue.js Project):**

```text
bioHub/
  index.html
  package.json
  tsconfig.json
  tsconfig.node.json
  vite.config.ts
  README.md
  public/
    avatar.jpg
    favicon.svg
    icons/
      buymeacoffee.svg
      github.svg
      linkedin.svg
      tiktok.svg
      twitter.svg
      youtube.svg
  src/
    App.vue
    index.css
    main.ts
    vite-env.d.ts
    components/
      Footer.vue
      Header.vue
      LinkContainer.vue
      SocialList.vue
    constant/
      index.ts
```

**After (Go Static Site Generator):**

```text
go-biohub/
  config.yml
  README.md
  template/
    index.html
    css/
      styles.css
  static/
    avatar.jpg
    favicon.svg
    icons/
      buymeacoffee.svg
      github.svg
      linkedin.svg
      tiktok.svg
      twitter.svg
      youtube.svg
  cmd/
    build-bio/
      build.go
  dist/
    index.html
    css/
      styles.css
    avatar.jpg
    favicon.svg
    icons/
      buymeacoffee.svg
      github.svg
      linkedin.svg
      tiktok.svg
      twitter.svg
      youtube.svg
```

### 15.2. Migration Steps

1. **Copy Static Assets**

- Move all files from `public/` (including `icons/`) to `static/` in the Go project.

2. **Migrate CSS**

- Move `src/index.css` to `template/css/styles.css`.
- Update paths in the HTML template to reference `/static/css/styles.css`.

3. **Convert Vue Templates to Go HTML Template**

- Use `src/App.vue` and all components (`src/components/`) as reference to build `template/index.html`.
- Replace Vue bindings with Go template syntax (e.g., `{{ .Params.Name }}`).
- Integrate header, footer, social links, and link containers directly into the Go template.

4. **Extract Configurable Data to YAML**

- Move all static/configurable data (avatar, name, headline, theme, socials, links) from Vue/TS files to `config.yml`.
- Follow the YAML schema defined in section 3.

5. **Implement Go Static Site Generator**

- Set up Go project as described in section 4.
- Write Go code in `cmd/build-bio/build.go` to:
  - Load `config.yml`.
  - Parse YAML into Go structs.
  - Render `template/index.html` with data.
  - Output to `dist/index.html`.

6. **Update Documentation**

- Revise `README.md` to document Go-based workflow, config editing, and build steps.

7. **Remove Vue/Node Tooling**

- Delete `package.json`, `tsconfig.json`, and all Vue-specific files after migration is validated.

### 15.3. Go Usage Details for bioHub

- **Config Loading:** Use `gopkg.in/yaml.v3` to load and parse `config.yml`.
  Example:

  ```go
  import (
    "os"
    "gopkg.in/yaml.v3"
  )
  type Config struct { /* ... */ }
  f, _ := os.ReadFile("config.yml")
  var cfg Config
  yaml.Unmarshal(f, &cfg)
  ```

- **Template Rendering:** Use Go's `html/template` package to render `template/index.html` with config data.
  Example:

  ```go
  import "html/template"
  t := template.Must(template.ParseFiles("template/index.html"))
  t.Execute(os.Stdout, cfg)
  ```

- **Static Assets:** Serve or copy everything in `static/` to the output directory (`dist/`).

- **Build & Run:** Use the provided Makefile for building and running the generator.

- **Testing:** Validate the output in `dist/` matches the original Vue app visually and functionally.

---

This section ensures a clear, actionable path for migrating every part of the current Vue.js repo to a Go-based static site generator, with all Go usage details included for reference.
