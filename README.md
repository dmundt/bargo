# bargo

[![CI](https://github.com/dmundt/bargo/actions/workflows/ci.yml/badge.svg)](https://github.com/dmundt/bargo/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/dmundt/bargo.svg)](https://pkg.go.dev/github.com/dmundt/bargo)

bargo is a small Go library for rendering single-line terminal progress bars with an optional centered percentage label.

## Installation

```bash
go get github.com/dmundt/bargo
```

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/dmundt/bargo"
)

func main() {
	b := bargo.New()
	fmt.Println(b.Render(88.2, 60))
}
```

Example output:

```text
[===========================88.2%===================       ]
```

## API

- `bargo.New(opts ...bargo.Option) *bargo.Bar`
- `(*Bar).Render(progress float64, width int) string`
- `(*Bar).RenderWithText(progress float64, width int, suffix string) string`
- `(*Bar).WriteTo(w io.Writer, progress float64, width int) (int, error)`
- `(*Bar).WriteToWithText(w io.Writer, progress float64, width int, suffix string) (int, error)`

Defaults:

- Progress input scale is `0` to `100`
- Width is the inner bar width (excluding delimiters)
- Percentage precision defaults to `1` decimal place
- Delimiters are `[` and `]`
- Fill and empty runes are `=` and space
- Moving head rune defaults to `=`
- Center percentage label is visible by default
- Out-of-range input is clamped by default

## Options

```go
b := bargo.New(
	bargo.WithPercentDecimals(1),
	bargo.WithFillRune('='),
	bargo.WithHeadRune('>'),
	bargo.WithEmptyRune(' '),
	bargo.WithPercentVisible(true),
	bargo.WithClamp(true),
)
```

Per-update text can be appended after the bar without changing its width:

```go
b := bargo.New()
fmt.Println(b.RenderWithText(5, 24, "20MB/400MB"))
fmt.Println(b.RenderWithText(40, 24, "File 4/10"))
```

## Runnable example

Run the included sample program:

```bash
go run ./example
```

The example is a showcase of six option combinations:

1. defaults
2. custom head rune
3. hidden percent label
4. custom delimiters and runes
5. custom bounds (0-400)
6. per-update suffix

## Showcase

The runnable example in `example/main.go` prints these six configurations:

```go
shows := []showcase{
	{
		title:    "1) defaults",
		bar:      bargo.New(),
		progress: 12,
		width:    36,
	},
	{
		title:    "2) custom head rune",
		bar:      bargo.New(bargo.WithHeadRune('>')),
		progress: 85,
		width:    36,
	},
	{
		title:    "3) hidden percent label",
		bar:      bargo.New(bargo.WithPercentVisible(false)),
		progress: 55,
		width:    36,
	},
	{
		title: "4) custom delimiters and runes",
		bar: bargo.New(
			bargo.WithLeftDelim("{"),
			bargo.WithRightDelim("}"),
			bargo.WithFillRune('#'),
			bargo.WithHeadRune('>'),
			bargo.WithEmptyRune('.'),
		),
		progress: 68,
		width:    36,
	},
	{
		title: "5) custom bounds (0-400)",
		bar: bargo.New(
			bargo.WithBounds(0, 400),
			bargo.WithPercentDecimals(0),
		),
		progress: 260,
		width:    36,
	},
	{
		title: "6) per-update suffix",
		bar: bargo.New(
			bargo.WithHeadRune('>'),
		),
		progress: 35,
		width:    36,
		suffix:   "File 4/10",
	},
}
```

Example output style:

```text
bargo showcase: 6 option combinations

1) defaults
[====           12.0%                ]

2) custom head rune
[===============85.0%==========>     ]

3) hidden percent label
[====================                ]

4) custom delimiters and runes
{###############68.0%###>............}

5) custom bounds (0-400)
[================260%===             ]

6) per-update suffix
[============>  35.0%                ] File 4/10
```

## Testing

```bash
go test ./...
```

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for the full release history.

## License

[MIT](LICENSE)
