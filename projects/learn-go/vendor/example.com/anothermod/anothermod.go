package anothermod

import "fmt"

// For production use, you’d publish the example.com/anotherMod module from its repository
// (with a module path that reflected its published location), where Go tools could find it to download it.

// For now, because you haven't published the module yet,
// you need to adapt the learn-go module so it can find the example.com/anotherMod code on your local file system.

// run `go mod edit -replace example.com/another-mod=./another-mod`
func AnotherMod() {
	fmt.Println("Hello from another mod")
}
