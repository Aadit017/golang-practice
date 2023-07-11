// calling an external package
package main

import (
	"fmt"

	"rsc.io/quote/v4"
)

func calling_external() {
	fmt.Println(quote.Go())
}
