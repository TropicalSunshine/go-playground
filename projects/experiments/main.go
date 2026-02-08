package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		go countTask(i)
	}
}

// count to 1000
func countTask(id int) bool {
	for i := 0; i < 1000; i++ {
	}

	fmt.Printf("go routine %d", id)
	return true
}
