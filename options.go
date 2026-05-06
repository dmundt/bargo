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

func WithBounds(min, max float64) Option {
	return func(c *config) {
		c.min = min
		c.max = max
	}
}

func WithPercentDecimals(n int) Option {
	return func(c *config) {
		if n < 0 {
			n = 0
		}
		c.decimals = n
	}
}

func WithFillRune(r rune) Option {
	return func(c *config) {
		c.fillRune = r
	}
}

func WithEmptyRune(r rune) Option {
	return func(c *config) {
		c.emptyRune = r
	}
}

func WithLeftDelim(s string) Option {
	return func(c *config) {
		c.leftDelim = s
	}
}

func WithRightDelim(s string) Option {
	return func(c *config) {
		c.rightDelim = s
	}
}

func WithClamp(enabled bool) Option {
	return func(c *config) {
		c.clamp = enabled
	}
}

func WithCarriageReturn(enabled bool) Option {
	return func(c *config) {
		c.carriageReturn = enabled
	}
}
