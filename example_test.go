package bargo

import "fmt"

func ExampleBar_Render() {
	b := New()
	fmt.Println(b.Render(50, 20))
	// Output: [=======50.0%        ]
}
