# go-stdlib ![go-stdlib](/images/google-go.png)

Code examples using [Go's standard library](https://pkg.go.dev/std).[^1] It is good to know what Go offers out of the box before building addition modules. The directories in this repo match the directory structure in the standard library.

This is a collection of various public examples. Each example in the standard library docs has a code window, so definitely use that first. Some of the examples in this repo expand or modify an example, and there is a README for each example that, at times, provides additional insight.

[Types](https://github.com/kengibson1111/go-stdlib/blob/master/types.md) is a summary of primitive types in Go. There are times when it helps with the examples.

Before diving into this repo, it may help to start with the [Tour of Go](https://go.dev/tour/welcome/1). A companion repo for the tour is [here](https://github.com/kengibson1111/tour-of-go).

## .netrc

All leaf directories are `package main` binaries. To download and install a binary, `go install` requires a $HOME/.netrc with an entry for github.com authentication:

`machine github.com login <github username> password <github personal access token>`

To install a binary in $GOPATH/bin, navigate to a leaf directory in the repo. Each leaf directory has a brief README. Use `go install`.

For example, if you wanted to install `log/gostdlib-log-logger`:

```bash
$ go install github.com/kengibson1111/go-stdlib/log/gostdlib-log-logger@latest
```

## work in progress

* There are still a few directories to convert to module format.

* With Go upgrades, I will do my best to keep examples current, enhance examples with command-line arguments, and add examples when relevant.

***

[^1]: Icon provided by [seekvectors.com](https://seekvectors.com/).
