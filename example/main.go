package main

import (
	"fmt"

	"github.com/dmundt/bargo"
)

type showcase struct {
	title    string
	bar      *bargo.Bar
	progress float64
	width    int
	suffix   string
}

func main() {
	shows := []showcase{
		{
			title:    "1) defaults",
			bar:      bargo.New(),
			progress: 12,
			width:    36,
		},
		{
			title:    "2) custom head rune",
			bar:      bargo.New(bargo.WithHeadRune('>')),
			progress: 85,
			width:    36,
		},
		{
			title:    "3) hidden percent label",
			bar:      bargo.New(bargo.WithPercentVisible(false)),
			progress: 55,
			width:    36,
		},
		{
			title: "4) custom delimiters and runes",
			bar: bargo.New(
				bargo.WithLeftDelim("{"),
				bargo.WithRightDelim("}"),
				bargo.WithFillRune('#'),
				bargo.WithHeadRune('>'),
				bargo.WithEmptyRune('.'),
			),
			progress: 68,
			width:    36,
		},
		{
			title: "5) custom bounds (0-400)",
			bar: bargo.New(
				bargo.WithBounds(0, 400),
				bargo.WithPercentDecimals(0),
			),
			progress: 260,
			width:    36,
		},
		{
			title: "6) per-update suffix",
			bar: bargo.New(
				bargo.WithHeadRune('>'),
			),
			progress: 35,
			width:    36,
			suffix:   "File 4/10",
		},
	}

	fmt.Println("bargo showcase: 6 option combinations")
	fmt.Println()
	for _, s := range shows {
		line := s.bar.Render(s.progress, s.width)
		if s.suffix != "" {
			line = s.bar.RenderWithText(s.progress, s.width, s.suffix)
		}
		fmt.Printf("%s\n%s\n\n", s.title, line)
	}
}
