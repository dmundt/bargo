// Package bargo provides a single-line terminal progress bar with an
// embedded percentage label. Progress values are on a 0–100 scale.
package bargo

import "io"

// Bar renders a progress bar. Create one with [New].
type Bar struct {
	cfg config
}

// New returns a Bar configured by the given options.
// Defaults: progress range 0–100, fill '=', empty ' ', delimiters '[' and ']',
// 1 decimal place, clamping enabled, carriage return disabled.
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
// value (0–100) and inner width (excluding delimiters).
// If clamping is enabled, progress is silently clamped to the configured
// min/max bounds before rendering.
func (b *Bar) Render(progress float64, width int) string {
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
	return b.cfg.leftDelim + inner + b.cfg.rightDelim
}

// WriteTo writes the rendered progress bar to w. If [WithCarriageReturn] is
// enabled a '\r' is prepended so the bar overwrites the current terminal line.
func (b *Bar) WriteTo(w io.Writer, progress float64, width int) (int, error) {
	out := b.Render(progress, width)
	if b.cfg.carriageReturn {
		out = "\r" + out
	}
	return io.WriteString(w, out)
}
