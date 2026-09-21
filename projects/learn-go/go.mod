// in a module you collect one or more related packages for a discrete and useful set of functions

// for example you might create a module with packages that have functions for doing financial analysis so that others
// writing financial applications can use your work

// go code is grouped into packages
// packages are grouped into modules

// your module specifies dependencies needed to run your code
// including go version and the set of other modules it requires

// As you add or improve functionality in your module,
// you publish new versions of the module.
// Developers writing code that calls functions in your module can import the module's updated packages
// and test with the new version before putting it into production use.
module learn-go

go 1.25

replace example.com/anothermod => ./anothermod

// The command found the local code in the greetings directory,
// then added a require directive to specify that example.com/hello requires example.com/greetings.
// You created this dependency when you imported the greetings package in hello.go.
// The number following the module path is a pseudo-version number -- a generated number used in place of a semantic version number (which the module doesn't have yet).

// To reference a published module,
// a go.mod file would typically omit the replace directive and use a require directive with a tagged version number at the end.
// require example.com/greetings v1.1.0
require example.com/anothermod v0.0.0-00010101000000-000000000000
