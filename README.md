# bargo

bargo is a small Go library for rendering single-line terminal progress bars with an embedded percentage label.

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
- `(*Bar).WriteTo(w io.Writer, progress float64, width int) (int, error)`

Defaults:

- Progress input scale is `0` to `100`
- Width is the inner bar width (excluding delimiters)
- Percentage precision defaults to `1` decimal place
- Delimiters are `[` and `]`
- Fill and empty runes are `=` and space
- Out-of-range input is clamped by default

## Options

```go
b := bargo.New(
	bargo.WithPercentDecimals(1),
	bargo.WithFillRune('='),
	bargo.WithEmptyRune(' '),
	bargo.WithClamp(true),
)
```

## Runnable example

Run the included sample program:

```bash
go run ./example
```

The example uses `WithCarriageReturn(true)` and `WriteTo` so progress updates are rendered on the same terminal line.

It runs on a timer and increments from `0` to `100` percent.

Expected style:

```text
[============================0.0%                            ]
...
[===========================100.0%===========================]
```

## Testing

```bash
go test ./...
```

## License

[MIT](LICENSE)
