package bargo

import "io"

type Bar struct {
	cfg config
}

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

func (b *Bar) WriteTo(w io.Writer, progress float64, width int) (int, error) {
	out := b.Render(progress, width)
	if b.cfg.carriageReturn {
		out = "\r" + out
	}
	return io.WriteString(w, out)
}
