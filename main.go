package main

import (
	"fmt"
	"os"
)

func main() {
	if err := App().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(100)
	}
}
