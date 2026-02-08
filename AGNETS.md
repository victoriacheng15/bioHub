# Agent Guide for bioHub

This document provides context and instructions for AI agents working on the **bioHub** project, a custom Static Site Generator (SSG) written in Go.

## 1. Project Overview

**bioHub** is a personal website and links platform built with a custom Go-based SSG. It is designed for simplicity, high performance, and zero external runtime dependencies.

- **Core Tech**: Go (Golang) 1.25+
- **Styling**: Minimal CSS / HTML templates
- **Content**: Configuration-driven (via `config.yml`) and HTML templates
- **Goal**: Single-binary simplicity and fast deployment.

## 2. Build and Test Commands

The project uses a **Nix wrapper logic** within the `Makefile` to ensure a reproducible environment. You can run standard `make` commands; they will automatically use `nix-shell` if available.

| Command | Description |
| :--- | :--- |
| `make build` | **Primary Build Command**. Builds the `biohub` binary and generates the site in `dist/`. |
| `make test` | Runs all Go unit tests. |
| `make vet` | Runs `go vet` for static analysis. |
| `make format` | Automatically formats all Go code. |
| `make cov-log` | Generates and displays a test coverage report in the terminal. |

**Important:** The build system handles the Nix environment automatically. If you are already inside a `nix-shell`, the commands will run directly.

## 3. Code Style Guidelines

### Go

- **Strict Adherence**: Code **must** pass `go fmt` and `go vet`.
- **Idiomatic Go**: Prefer standard library solutions. Keep functions small and focused.
- **Error Handling**: Handle errors explicitly. Use descriptive error messages.
- **Imports**: Group standard library imports separately from third-party imports.

### Configuration & Templates

- **YAML**: `config.yml` manages site metadata and links. Maintain clear structure.
- **Templates**: HTML templates are located in `template/`. Maintain clean, semantic HTML.

## 4. Testing Instructions

- **Unit Tests**: Run `make test` to execute the Go test suite.
- **Coverage**: Run `make cov-log` to see a detailed coverage report in the terminal.
- **New Features**: Any new logic in the SSG **must** include accompanying unit tests in `main_test.go`.

## 5. Security & Automation

- **CI/CD**: GitHub Actions (`lint.yml`, `deploy.yml`) handle linting, testing, and deployment to GitHub Pages.
- **File System**: The SSG reads from `template/` and `config.yml` and writes to `dist/`.
- **Automation**: CI workflows are aligned with `Makefile` targets to ensure consistency between local and remote environments.