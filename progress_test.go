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

func TestRenderHeadRuneAtPartialProgress(t *testing.T) {
	b := New(WithHeadRune('>'), WithPercentVisible(false))
	got := b.Render(50, 10)

	if got != "[====>     ]" {
		t.Fatalf("expected head rune at fill frontier, got %q", got)
	}
}

func TestRenderNoHeadAtFullProgress(t *testing.T) {
	b := New(WithHeadRune('>'), WithPercentVisible(false))
	got := b.Render(100, 10)

	if got != "[==========]" {
		t.Fatalf("expected fully filled bar without head rune, got %q", got)
	}
}

func TestRenderPercentHidden(t *testing.T) {
	b := New(WithPercentVisible(false))
	got := b.Render(50, 20)

	if strings.Contains(got, "%") {
		t.Fatalf("expected no percent label, got %q", got)
	}
	if len([]rune(got)) != 22 {
		t.Fatalf("expected total width 22, got %d (%q)", len([]rune(got)), got)
	}
}

func TestRenderWithTextAppendsSuffix(t *testing.T) {
	b := New()
	got := b.RenderWithText(50, 20, "20MB/400MB")

	if !strings.HasSuffix(got, "] 20MB/400MB") {
		t.Fatalf("expected suffix after bar, got %q", got)
	}
}

func TestRenderWithTextEmptySuffixMatchesRender(t *testing.T) {
	b := New()
	if got, want := b.RenderWithText(42, 16, ""), b.Render(42, 16); got != want {
		t.Fatalf("expected empty suffix output to match Render: got %q want %q", got, want)
	}
}

func TestWriteToWithTextCarriageReturn(t *testing.T) {
	b := New(WithCarriageReturn(true))
	var buf bytes.Buffer

	_, err := b.WriteToWithText(&buf, 25, 12, "File 4/10")
	if err != nil {
		t.Fatalf("unexpected write error: %v", err)
	}
	if !strings.HasPrefix(buf.String(), "\r[") {
		t.Fatalf("expected carriage-return prefixed output, got %q", buf.String())
	}
	if !strings.HasSuffix(buf.String(), "] File 4/10") {
		t.Fatalf("expected suffix after bar, got %q", buf.String())
	}
}

func TestWriteToWithTextPropagatesError(t *testing.T) {
	b := New()
	_, err := b.WriteToWithText(failingWriter{}, 10, 10, "x")
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
