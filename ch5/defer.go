package ch5

import (
	"fmt"
	"os"
)

// some resources require that we explicitly release them
// in case we forget to release them, a common pattern is to `defer` their release

func Defer_the_closing_of_a_file() {
	file, err := os.Open("a_file_to_read")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// read the file...
}
