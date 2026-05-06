package bargo

type config struct {
	min            float64
	max            float64
	decimals       int
	fillRune       rune
	emptyRune      rune
	leftDelim      string
	rightDelim     string
	clamp          bool
	carriageReturn bool
}

// Option is a functional option for configuring a [Bar].
type Option func(*config)

func defaultConfig() config {
	return config{
		min:            0,
		max:            100,
		decimals:       1,
		fillRune:       '=',
		emptyRune:      ' ',
		leftDelim:      "[",
		rightDelim:     "]",
		clamp:          true,
		carriageReturn: false,
	}
}

// WithBounds sets the minimum and maximum values of the progress range.
// The default range is 0 to 100.
func WithBounds(min, max float64) Option {
	return func(c *config) {
		c.min = min
		c.max = max
	}
}

// WithPercentDecimals sets the number of decimal places shown in the
// percentage label. Negative values are treated as zero. Default is 1.
func WithPercentDecimals(n int) Option {
	return func(c *config) {
		if n < 0 {
			n = 0
		}
		c.decimals = n
	}
}

// WithFillRune sets the rune used for the filled portion of the bar.
// Default is '='.
func WithFillRune(r rune) Option {
	return func(c *config) {
		c.fillRune = r
	}
}

// WithEmptyRune sets the rune used for the unfilled portion of the bar.
// Default is ' ' (space).
func WithEmptyRune(r rune) Option {
	return func(c *config) {
		c.emptyRune = r
	}
}

// WithLeftDelim sets the opening delimiter of the bar. Default is "[".
func WithLeftDelim(s string) Option {
	return func(c *config) {
		c.leftDelim = s
	}
}

// WithRightDelim sets the closing delimiter of the bar. Default is "]".
func WithRightDelim(s string) Option {
	return func(c *config) {
		c.rightDelim = s
	}
}

// WithClamp controls whether out-of-range progress values are silently
// clamped to the configured bounds. Default is true.
func WithClamp(enabled bool) Option {
	return func(c *config) {
		c.clamp = enabled
	}
}

// WithCarriageReturn controls whether [Bar.WriteTo] prepends '\r' to the
// output, causing the bar to overwrite the current terminal line. Default is false.
func WithCarriageReturn(enabled bool) Option {
	return func(c *config) {
		c.carriageReturn = enabled
	}
}
