package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dmundt/bargo"
)

func main() {
	b := bargo.New(bargo.WithCarriageReturn(true))
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for v := 0.0; v <= 100.0; v++ {
		<-ticker.C
		_, _ = b.WriteTo(os.Stdout, v, 60)
	}
	fmt.Println()
}
