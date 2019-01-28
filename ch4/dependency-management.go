package ch4

// If we `go get` within a project, it'll scan all the files, looking for `imports`
// to third-party libraries and will download them.
// In a way, our own source code becomes a `Gemfile` or `package.json`.

// If you call `go get -u` it'll update the packages (or you can update a specific
// package via `go get -u FULL_PACKAGE_NAME`).

// ## Shortcomings with `go get`

// Eventually, you might find `go get` inadequate.
// For one thing, there's no way to specify a revision, it always points to the
// master/head/trunk/default.
// This is an even larger problem if you have two projects needing different versions
// of the same library.

// To solve this, you can use a third-party dependency management tool.
// They are still young, but two promising ones are:
//
// 1. [goop](https://github.com/nitrous-io/goop)
// 2. [godep](https://github.com/tools/godep).
//
//  A more complete list is available at the [go-wiki](https://code.google.com/p/go-wiki/wiki/PackageManagementTools).
