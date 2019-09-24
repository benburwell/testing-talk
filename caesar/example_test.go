package caesar

import (
	"fmt"
)

func ExampleEncode() {
	secret := Encode("attack at dawn")
	fmt.Println(secret)
	// Output: dwwdfn dw gdzq
}

func ExampleCoder() {
	c := Coder{
		Key: 5,
		Ranges: []RuneRange{
			{Start: 'A', End: 'Z'},
		},
	}
	fmt.Println(c.Encode("aAbBcCdD"))
	// Output: aFbGcHdI
}
