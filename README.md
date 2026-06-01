# axiom-tui

A lightweight, keyboard-driven terminal control plane for local infrastructure. It brings Docker containers, systemd services, and basic system metrics into a clean, unified TUI dashboard.

![axiom-tui-screenshot](https://github.com/drunkswamp/axiom-tui/assets/12345/your-screenshot-filename.png) 
*(Note: Add a real screenshot here after committing.)*

**Project Status:** Early work-in-progress. The application is buildable and functional for basic monitoring, but under active development.

## Key Features

*   **Docker Dashboard:** See running containers, check their status, and monitor basic resource usage from the Docker tab. The view correctly distinguishes between loading, empty, error, and populated list states.
*   **Systemd Services:** The systemd tab currently shows mock-backed services, providing a functional preview of service monitoring.
*   **Live Metrics:** Get a quick system health check with a CPU and RAM usage snapshot on the Metrics tab.
*   **Bilingual UI:** Defaults to English. Switch to Russian by setting the `AXIOMTUI_LANG=ru` environment variable.
*   **Cross-Platform:** Builds and runs on both Linux and Windows.
*   **Pure Go TUI:** Built with the excellent Charmbracelet Bubble Tea framework for a fast, responsive terminal experience.

## Requirements

*   **Go:** Version 1.26.3 or a compatible version that supports the module's dependencies.
*   **Docker:** A running Docker daemon is optional for the application to start but required for the Docker tab to show container data.
*   **Git:** Required for cloning the repository.

## Build and Run From Source

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/drunkswamp/axiom-tui.git
    cd axiom-tui
    ```

2.  **Navigate to the Go module:**
    The active Go module is in the `axiom-tui/` subdirectory.
    ```bash
    cd axiom-tui
    ```

3.  **Download dependencies:**
    ```bash
    go mod download
    ```

4.  **(Recommended) Run tests:**
    ```bash
    go test -v ./...
    ```

5.  **Build the binary:**
    ```bash
    go build -o axiomtui ./cmd/axiomtui
    ```
    This creates an executable named `axiomtui` in the current directory.

6.  **Run the application:**
    ```bash
    ./axiomtui
    ```
    The application will start in English by default.

## Language Selection

*   **Run in English (default):**
    ```bash
    ./axiomtui
    ```
*   **Run in Russian:**
    ```bash
    AXIOMTUI_LANG=ru ./axiomtui
    ```

## Development

To verify your changes during development, run these commands from the `axiom-tui/` directory:

```bash
# Run the full test suite
go test -v ./...

# Build for your native OS
go build -v ./...

# Cross-compile for Linux to ensure CI will pass
GOOS=linux GOARCH=amd64 go build -v ./...
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
