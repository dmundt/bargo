package bargo

import (
	"strings"
	"testing"
)

func TestFormatPercentForWidth(t *testing.T) {
	tests := []struct {
		name        string
		value       float64
		decimals    int
		width       int
		expect      string
		expectWidth int
	}{
		{name: "fits default precision", value: 88.2, decimals: 1, width: 8, expect: "88.2%"},
		{name: "drops to integer", value: 93.456, decimals: 2, width: 4, expect: "93%"},
		{name: "truncates when needed", value: 100, decimals: 2, width: 3, expect: "100"},
		{name: "empty when width zero", value: 10, decimals: 1, width: 0, expect: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatPercentForWidth(tt.value, tt.decimals, tt.width)
			if got != tt.expect {
				t.Fatalf("expected %q, got %q", tt.expect, got)
			}
		})
	}
}

func TestNormalizeRatio(t *testing.T) {
	if got := normalizeRatio(50, 0, 100); got != 0.5 {
		t.Fatalf("expected 0.5, got %v", got)
	}
	if got := normalizeRatio(50, 10, 10); got != 0 {
		t.Fatalf("expected 0 for equal bounds, got %v", got)
	}
	if got := normalizeRatio(50, 100, 10); got != 0 {
		t.Fatalf("expected 0 for inverted bounds, got %v", got)
	}
}

func TestRenderInnerHeadRune(t *testing.T) {
	cfg := defaultConfig()
	cfg.showPercent = false
	cfg.headRune = '>'

	if got := renderInner(10, 50, 0.5, cfg); got != "====>     " {
		t.Fatalf("expected head rune at frontier, got %q", got)
	}
}

func TestRenderInnerNoHeadAtNearZeroRatio(t *testing.T) {
	// ratio 0.01 rounds to fillCount 0 at width 10; no head should appear.
	cfg := defaultConfig()
	cfg.showPercent = false
	cfg.headRune = '>'

	if got := renderInner(10, 1, 0.01, cfg); strings.ContainsRune(got, '>') {
		t.Fatalf("expected no head rune when fillCount rounds to zero, got %q", got)
	}
}

func TestRenderInnerNoHeadAtComplete(t *testing.T) {
	cfg := defaultConfig()
	cfg.showPercent = false
	cfg.headRune = '>'

	if got := renderInner(10, 100, 1, cfg); got != "==========" {
		t.Fatalf("expected full fill without head rune, got %q", got)
	}
}
