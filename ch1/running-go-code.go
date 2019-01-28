package ch1

import (
	"fmt"
	"os"
)

func Running_go_code() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])
}
