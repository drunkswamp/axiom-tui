## Currrently working on it.  

# axiom-tui

A clean, lightweight console dashboard that brings Docker containers, systemd services, and resource graphs into a single terminal view. Built purely in Go with the Charmbracelet Bubble Tea framework.

## Key Features

* **Docker Management:** Monitor local containers, view statuses, and track active workloads in real-time.
* **Systemd Integration:** Keep track of essential background services and system units directly from the TUI.
* **Live Metrics:** Lightweight resource monitors and graphs for quick system health checks.
* **Keyboard-Driven:** Fast, responsive, and completely optimized for terminal-only environments.

## Installation
Make sure you have Go installed on your machine, then run:

```bash

go install github.com/drunkswamp/axiom-tui/cmd/axiomtui@latest

```

Note: The executable will be installed to your $GOPATH/bin directory. Ensure this directory is added to your system's PATH to run the tool from anywhere.

## Quick Start
Simply run the compiled binary or call it directly if it's in your PATH:

```bash

axiomtui

```

## Hotkeys
Tab / Shift+Tab — Switch between component panels (Docker, Systemd, Metrics)

Arrows / Keyboard shortcuts — Navigate within lists

Ctrl+C / Q — Safely exit the application

## Requirements
Go 1.23.0 or higher

Docker daemon running locally

Windows / Linux environment (with systemd support for unit tracking)

## License
MIT License. Feel free to use, modify, and distribute.
