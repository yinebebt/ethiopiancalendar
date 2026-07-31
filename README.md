<p align="center">
<img src="logo.svg" alt="logo" width="110" height="110">
</p>
<h1 align="center"><a href="https://pkg.go.dev/github.com/yinebebt/ethiocal">Ethiocal — Ethiopian Calendar (ባሕረ-ሐሳብ)</a></h1>

[![Go Reference](https://pkg.go.dev/badge/github.com/yinebebt/ethiocal.svg)](https://pkg.go.dev/github.com/yinebebt/ethiocal)
[![ci-badge](https://github.com/yinebebt/ethiocal/actions/workflows/ci.yml/badge.svg)](https://github.com/yinebebt/ethiocal/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/yinebebt/ethiocal)](https://goreportcard.com/report/github.com/yinebebt/ethiocal)

## Description

Ethiocal is an Ethiopian calendar (ባሕረ-ሐሳብ) tool for retrieving fasting and holiday dates based on the
[EOTC](https://www.ethiopianorthodox.org/) calendar, and converting between Ethiopian and
Gregorian dates. Ethiopia follows its own calendar with 13 months (twelve 30-day months
plus a 5- or 6-day 13th month).

### Features

* **GUI app** — graphical interface for desktop (macOS/Linux/Windows) and Android.
* **CLI** — `ethiocal-cli` subcommands for scripting and terminal use.
* Get Ethiopian fasting and religious festival dates for a specific year.
* Convert Ethiopian dates to Gregorian dates and vice versa.

## Installation

### Download a binary (no Go required)

Pre-built binaries are available on the [Releases](https://github.com/yinebebt/ethiocal/releases) page. Pick your platform:

| Platform | Download |
| --- | --- |
| macOS (Apple Silicon) | `curl -Lo ethiocal.zip https://github.com/yinebebt/ethiocal/releases/latest/download/ethiocal-macos-arm64.app.zip && unzip ethiocal.zip` |
| macOS (Intel) | `curl -Lo ethiocal.zip https://github.com/yinebebt/ethiocal/releases/latest/download/ethiocal-macos-amd64.app.zip && unzip ethiocal.zip` |
| Linux (x86_64) | `curl -Lo ethiocal.tar.xz https://github.com/yinebebt/ethiocal/releases/latest/download/ethiocal-linux-amd64.tar.xz && tar xJf ethiocal.tar.xz` |
| Windows (x86_64) | `curl -Lo ethiocal.zip https://github.com/yinebebt/ethiocal/releases/latest/download/ethiocal-windows-amd64.exe.zip && unzip ethiocal.zip` |
| Android (arm64) | [`ethiocal-android.apk`](https://github.com/yinebebt/ethiocal/releases/latest) — sideload on a 64-bit device |

The CLI tool ships per platform too, as `ethiocal-cli-<os>-<arch>` on the same Releases page.

### Install with Go

```bash
go install github.com/yinebebt/ethiocal@latest
```

### Build from source

```bash
git clone https://github.com/yinebebt/ethiocal.git
cd ethiocal
make build    # builds both: ./ethiocal (GUI) and ./ethiocal-cli (CLI)
```

> **Note:** Building the GUI app requires a C compiler and OpenGL headers because
> Fyne uses CGO. On Ubuntu/Debian: `sudo apt-get install libgl1-mesa-dev xorg-dev`.
> macOS and Windows have these out of the box. The CLI builds with pure Go (no CGO).

## Usage

### GUI

Run `ethiocal` (or launch the packaged app) to open the GUI:

```bash
ethiocal
```

The GUI provides three tabs:

* **Home** — today’s Gregorian and Ethiopian dates, plus a few nearby feasts.
* **Date Converter** — convert between Gregorian and Ethiopian calendars.
* **Bahire-Hasab** — fasting and festival dates for an Ethiopian year.

### CLI

The terminal tool is a separate binary (`ethiocal-cli`), so the GUI app
carries no CLI dependencies. Install with
`go install github.com/yinebebt/ethiocal/cmd/ethiocal-cli@latest`.

```bash
# Get religious dates for Ethiopian year 2017
ethiocal-cli bahir 2017

# Convert Gregorian date to Ethiopian (YYYY-MM-DD)
ethiocal-cli convert gtoe 2025-2-2

# Convert Ethiopian date to Gregorian
ethiocal-cli convert etog 2017-5-25
```

### As a Go library

```go
import (
    "github.com/yinebebt/ethiocal/bahirehasab"
    "github.com/yinebebt/ethiocal/dateconverter"
)

// Get festivals for Ethiopian year 2017
festival, err := bahirehasab.NewFestival(2017)

// Gregorian → Ethiopian
etDate, err := dateconverter.Ethiopian(2025, 2, 2)

// Ethiopian → Gregorian
gregDate, err := dateconverter.Gregorian(2017, 5, 25)
```

## License

See [LICENSE](LICENSE).


