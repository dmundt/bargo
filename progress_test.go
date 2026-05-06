package bargo

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestRenderBasicShape(t *testing.T) {
	b := New()
	got := b.Render(88.2, 60)

	if !strings.Contains(got, "88.2%") {
		t.Fatalf("expected percent label in output, got %q", got)
	}
	if len([]rune(got)) != 62 {
		t.Fatalf("expected total width 62, got %d (%q)", len([]rune(got)), got)
	}
}

func TestRenderZeroWidth(t *testing.T) {
	b := New()
	if got := b.Render(50, 0); got != "[]" {
		t.Fatalf("expected [] for zero width, got %q", got)
	}
}

func TestRenderTinyWidthDropsDecimals(t *testing.T) {
	b := New(WithPercentDecimals(2))
	got := b.Render(93.456, 4)

	if strings.Contains(got, ".") {
		t.Fatalf("expected decimal point to be dropped for tiny width, got %q", got)
	}
	if len([]rune(got)) != 6 {
		t.Fatalf("expected width 6 including delimiters, got %d (%q)", len([]rune(got)), got)
	}
}

func TestRenderClampBehavior(t *testing.T) {
	clamped := New()
	if got := clamped.Render(130, 20); !strings.Contains(got, "100.0%") {
		t.Fatalf("expected clamped percent text, got %q", got)
	}

	unclamped := New(WithClamp(false))
	if got := unclamped.Render(130, 20); !strings.Contains(got, "130.0%") {
		t.Fatalf("expected unclamped percent text, got %q", got)
	}
}

func TestCustomDelimitersAndRunes(t *testing.T) {
	b := New(
		WithLeftDelim("{"),
		WithRightDelim("}"),
		WithFillRune('#'),
		WithEmptyRune('-'),
	)
	got := b.Render(50, 14)

	if !strings.HasPrefix(got, "{") || !strings.HasSuffix(got, "}") {
		t.Fatalf("expected custom delimiters, got %q", got)
	}
	if !strings.ContainsRune(got, '#') {
		t.Fatalf("expected custom fill rune in output, got %q", got)
	}
}

func TestWriteToCarriageReturn(t *testing.T) {
	b := New(WithCarriageReturn(true))
	var buf bytes.Buffer

	_, err := b.WriteTo(&buf, 25, 12)
	if err != nil {
		t.Fatalf("unexpected write error: %v", err)
	}
	if !strings.HasPrefix(buf.String(), "\r[") {
		t.Fatalf("expected carriage-return prefixed output, got %q", buf.String())
	}
}

func TestWriteToPropagatesError(t *testing.T) {
	b := New()
	_, err := b.WriteTo(failingWriter{}, 10, 10)
	if err == nil {
		t.Fatal("expected write error")
	}
	if !errors.Is(err, errWriteFailed) {
		t.Fatalf("expected sentinel write error, got %v", err)
	}
}

var errWriteFailed = errors.New("write failed")

type failingWriter struct{}

func (f failingWriter) Write(_ []byte) (int, error) {
	return 0, errWriteFailed
}
