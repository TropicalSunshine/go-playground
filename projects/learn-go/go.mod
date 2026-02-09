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

go 1.24.2

replace example.com/anothermod => ./anothermod

require example.com/anothermod v0.0.0-00010101000000-000000000000
