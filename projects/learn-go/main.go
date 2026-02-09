// in go, code executed as an application must be in a main package
package main

import (
	packagee "learn-go/package"

	"example.com/anothermod"
)

func main() {
	packagee.Packagee()
	anothermod.AnotherMod()
}
