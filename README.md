# go-stdlib ![go-stdlib](/images/google-go.png)

Code examples using [Go's standard library](https://pkg.go.dev/std).[^1] It is good to know what Go offers out of the box before building addition modules. The directories in this repo match the directory structure in the standard library.

This is a collection of various public examples. Each example in the standard library docs has a code window, so definitely use that first if you prefer. Some of the examples in this repo expand or modify an example, and there is a README for each example that, at times, provides additional insight.

[Types](https://github.com/kengibson1111/go-stdlib/blob/master/types.md) is a summary of primitive types in Go. There are times when it helps with the examples.

Before diving into this repo, it may help to start with the [Tour of Go](https://go.dev/tour/welcome/1). My [companion repo](https://github.com/kengibson1111/tour-of-go) may help.

## .netrc

All leaf directories are `package main` binaries. To download and install a binary, `go install` requires a $HOME/.netrc with an entry for github.com authentication:

`machine github.com login <github username> password <github personal access token>`

## binary installation

To install a binary in $GOPATH/bin, navigate to a leaf directory in the repo. Each leaf directory has a brief README. Use `go install`. Currently, there are one or more "default example" modules for each package. A "default example" module usually shows basic use of one or more functions without processing command-line arguments.

For example, if you wanted to install the "default example" `log/gostdlib-log-logger`:

```bash
go install github.com/kengibson1111/go-stdlib/log/gostdlib-log-logger@latest
```

Then, run the binary `gostdlib-log-logger` in $GOPATH/bin.

Over time, new example modules will be added that are more complex - each intended to demonstrate a capability with a combination of functions. Even though a more complex example module may use functions from other packages, it is meant to highlight functions in a specific package.

## work in progress

* There are still a few directories to convert to module format.

* With Go upgrades, I will do my best to keep examples current, enhance examples with command-line arguments, and add examples when relevant.

***

[^1]: Icon provided by [seekvectors.com](https://seekvectors.com/).
