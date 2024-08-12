package packagemanagement

/*
	Package management in Go, involves organizing,
	importing and managing dependencies for your Go programs.
	Go uses a module system introduced in Go 1.11, which is designed
	to simplify dependency management and versioning.

*/

/* Package management in Go.
1. Module: a module is a collection of go packages stored in a directory with a go.mod file
at its root. The go.mod file defines the module properties, including it paths, dependencies,
and versions.

 * go init: creates and initializes go.mod file
	go mod init example.com/module-name
	Example: go mod init github.com/mjthecoder65/wallet

 * Adding Dependencies
	- When you import a new package and run go build or go test. Go will automatically add the dependency to you go.mod
	file and download it to your module's cache.

 * go get: this command add or updated dependencies. It update to package to it new version.
	go get example.com/somapage

 * go mod downlaod: download all the module specified in the modules specified in the go.mod file to the module cache.
 It does not bulid or test the code., but it ensures that all required modules are available locally.

  * Tidying UP:
    The go mod tidy command remove any dependencies from go.mod file that are not longer
	used in the module and adds any that are missing.

2. Dependency management.
The go.mod file is used to specify the dependencies for a module.
The go.sum file contains checksums of the dependencies to ensure integrity.
go.sum ensure that the dependencies used in your project are exactly what they are supposed to.

*/

/* Import Commands in Go.
Importing a package in Go is straightforward. You use import keyword followed by the package path.
Package can be from standard library or third-party libraries.

	1. Importing a single package
		import "fmt"
	2. Importing multiple packages
		import (
			"fmt"
			"math"
		)

	3. Import with Alias: when you want to use different name for the package
		import m "math"

	4. Blank Import: Used to import package only for its side effects, such initializing or
	registration purposes.
	  	import _ "net/http/pprof"

*/

func LearnPackageManagement() {
	// Understand go Path and packages are orgnanized
	// Learn about the necessary commands.
}
