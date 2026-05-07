// Package bargo provides a single-line terminal progress bar with an optional
// centered percentage label.
package bargo

import "io"

// Bar renders a progress bar. Create one with [New].
type Bar struct {
	cfg config
}

// New returns a Bar configured by the given options.
// Defaults: progress range 0–100, fill '=', empty ' ', delimiters '[' and ']',
// head '=', centered percent visible, 1 decimal place,
// clamping enabled, carriage return disabled.
func New(opts ...Option) *Bar {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.max <= cfg.min {
		cfg.max = cfg.min + 1
	}
	return &Bar{cfg: cfg}
}

// Render returns the formatted progress bar string for the given progress
// value and inner width (excluding delimiters).
// If clamping is enabled, progress is silently clamped to the configured
// min/max bounds before rendering.
func (b *Bar) Render(progress float64, width int) string {
	return b.render(progress, width, "")
}

// RenderWithText returns the formatted progress bar string with optional
// trailing text appended after the right delimiter.
func (b *Bar) RenderWithText(progress float64, width int, suffix string) string {
	return b.render(progress, width, suffix)
}

func (b *Bar) render(progress float64, width int, suffix string) string {
	value := progress
	if b.cfg.clamp {
		value = clamp(value, b.cfg.min, b.cfg.max)
	}

	ratio := normalizeRatio(value, b.cfg.min, b.cfg.max)
	if ratio < 0 {
		ratio = 0
	}
	if ratio > 1 {
		ratio = 1
	}

	inner := renderInner(width, value, ratio, b.cfg)
	out := b.cfg.leftDelim + inner + b.cfg.rightDelim
	if suffix != "" {
		out += " " + suffix
	}
	return out
}

// WriteTo writes the rendered progress bar to w.
// If [WithCarriageReturn] is enabled a '\r' is prepended so the output
// overwrites the current terminal line.
func (b *Bar) WriteTo(w io.Writer, progress float64, width int) (int, error) {
	out := b.Render(progress, width)
	if b.cfg.carriageReturn {
		out = "\r" + out
	}
	return io.WriteString(w, out)
}

// WriteToWithText writes the rendered progress bar with trailing text to w.
// If [WithCarriageReturn] is enabled a '\r' is prepended so the output
// overwrites the current terminal line.
func (b *Bar) WriteToWithText(w io.Writer, progress float64, width int, suffix string) (int, error) {
	out := b.RenderWithText(progress, width, suffix)
	if b.cfg.carriageReturn {
		out = "\r" + out
	}
	return io.WriteString(w, out)
}
