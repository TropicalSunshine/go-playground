package main

import (
	"log"
)

const (
	ROUNTINES = 10
	TILL      = 1000000
)

func main() {
	log.SetFlags(log.Lmicroseconds | log.Ltime)
	for i := 0; i < ROUNTINES; i++ {
		go countTask(i)
	}
	for {
	}
}

// count to 1000
func countTask(id int) bool {
	for i := 0; i < TILL; i++ {
	}

	log.Printf("go routine %d has finished counting to 1000 \n", id)
	return true
}
