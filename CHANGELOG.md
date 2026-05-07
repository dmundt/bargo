# Changelog

All notable changes to this project are documented here.
The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [1.1.0] — 2026-05-07

### Added

- `WithHeadRune(r rune)` option: configures the moving head rune at the fill
  frontier. Default is `=`, matching the existing fill rune so prior output is
  unchanged.
- `WithPercentVisible(enabled bool)` option: hides the centered percentage
  label when set to `false`. Default is `true`.
- `RenderWithText(progress, width, suffix)` method: returns the bar string with
  optional trailing text appended after the right delimiter. Bar width is
  unaffected by the suffix.
- `WriteToWithText(w, progress, width, suffix)` method: writes the above to any
  `io.Writer`, respecting `WithCarriageReturn`.
- Runnable example updated to a 6-combination showcase covering defaults, head
  rune, hidden label, custom delimiters/runes, custom bounds, and per-update
  suffix text.
- GitHub Actions CI workflow (`.github/workflows/ci.yml`) running `go test ./...`
  on every push and pull request to `main`.
- CI badge and Go Reference badge added to README.

## [1.0.1] — 2026-05-06

### Changed

- README and GoDoc descriptions synced to reflect current configurable
  behavior and all new options and methods.

## [1.0.0] — initial release

### Added

- Single-line terminal progress bar with centered percentage label.
- Functional options: `WithBounds`, `WithPercentDecimals`, `WithFillRune`,
  `WithEmptyRune`, `WithLeftDelim`, `WithRightDelim`, `WithClamp`,
  `WithCarriageReturn`.
- `Render` and `WriteTo` public API.
