package bargo

import (
	"fmt"
	"math"
)

func renderInner(width int, value float64, ratio float64, cfg config) string {
	if width <= 0 {
		return ""
	}

	bar := make([]rune, width)
	for i := range bar {
		bar[i] = cfg.emptyRune
	}

	fillCount := int(math.Round(ratio * float64(width)))
	if fillCount < 0 {
		fillCount = 0
	}
	if fillCount > width {
		fillCount = width
	}
	for i := 0; i < fillCount; i++ {
		bar[i] = cfg.fillRune
	}

	if ratio > 0 && ratio < 1 {
		headPos := fillCount - 1
		if headPos < 0 {
			headPos = 0
		}
		if headPos >= width {
			headPos = width - 1
		}
		bar[headPos] = cfg.headRune
	}

	if !cfg.showPercent {
		return string(bar)
	}

	percentText := formatPercentForWidth(value, cfg.decimals, width)

	label := []rune(percentText)
	if len(label) > width {
		label = label[:width]
	}
	start := (width - len(label)) / 2
	for i, r := range label {
		bar[start+i] = r
	}

	return string(bar)
}

func formatPercentForWidth(value float64, maxDecimals int, width int) string {
	if width <= 0 {
		return ""
	}
	if maxDecimals < 0 {
		maxDecimals = 0
	}
	for decimals := maxDecimals; decimals >= 0; decimals-- {
		t := fmt.Sprintf("%.*f%%", decimals, value)
		if len([]rune(t)) <= width {
			return t
		}
	}
	t := fmt.Sprintf("%.0f%%", value)
	r := []rune(t)
	if len(r) <= width {
		return t
	}
	return string(r[:width])
}

func clamp(v float64, min float64, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func normalizeRatio(v float64, min float64, max float64) float64 {
	d := max - min
	if d <= 0 {
		return 0
	}
	return (v - min) / d
}
