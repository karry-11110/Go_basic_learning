package main

import (
	"fmt"

	mypackage "test_gomod/package"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	mypackage.Ne()

}
