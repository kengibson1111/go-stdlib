# go-stdlib ![go-stdlib](/images/google-go.png)

Code examples using [Go's standard library](https://pkg.go.dev/std).[^1] It is good to know what Go offers out of the box before building addition modules. The directories in this repo match the directory structure in the standard library.

This is a collection of various public examples. Each example in the standard library docs has a code window, so definitely use that first. Some of the examples in this repo expand or modify an example, and there is a README for each example that, at times, provides additional insight.

[Types](https://github.com/kengibson1111/go-stdlib/blob/master/types.md) is a summary of primitive types in Go. There are times when it helps with the examples.

Before diving into this repo, it may help to start with the [Tour of Go](https://go.dev/tour/welcome/1). A companion repo for the tour is [here](https://github.com/kengibson1111/tour-of-go).

## reflect

* structtag - from golang's docs. It looks like golang provides an easy way to parse JSON-like tags and
  values within a string. Very nice. You can reflect an internal type that is a struct with one string.
  The string value is intended to be like JSON tags and values. Get the one field value through
  reflection. Then get the tags. Maybe a way to extend this example is to look up tag names
  dynamically and grab values.

* typeof - from golang's docs. This shows the use of TypeOf in order to check whether or not an
  interface implementation really implements an interface. Pretty simple.

## testing

* benchparallel - this doesn't output anything at this point, but I thought it was interesting
  because of the text/template work I was doing at the time. Benchmark functions are meant to be run
  as part of a test suite usually with the "go test -cpu" command. golang's baseline text/template
  functionality is very powerful, and I saw this as a value extension to any templates created and
  executed. I will most likely revisit this.

## text/scanner

This shows how to use scanner to parse text tokens.

## text/tabwriter

This shows how to use tabwriter to format text in neat columns.

## text/template

pre-cursor to html/template examples. Everything in html/template builds upon text/template.

* template - shows basic template functionality.

* block - this is showing how to customize a template block. block is shorthand for defining a
  template and executing it in place. That is what allows the template block to be redefined.
  In this sample, overlay is redefining the master template block. Notice how the master template
  instance assign the join function pointer and not the overlay. But the overlay template
  redefinition of the master template block uses join. Starting to get interesting when
  considering dynamic execution functionality.

* func - like the previous sample, this assigns a function pointer. I like this sample because
  it also shows how to use a pipe ("|") feeding the return value of one function into another
  like a Unix pipe.

* glob - this sample shows how to load all templates into a template instance based on a global
  pattern. Then if templates call other templates, they're all available in the same template
  instance. golang has some serious dynamic execution muscle.

* helpers - this builds on the previous sample using ParseGlob to load helper templates. Then
  it dynamically loads two more templates to the helpers and executes each of the new templates.
  Although this sample doesn't show it, imagine if your helpers are calling functions like
  in the block and func samples. Hmmmmm.

* share - this shows how to define a set of base helpers and share them. In this sample, T0
  and T1 are defined and we know that T1 will be calling T2, but T2 is not defined yet. Then
  the helpers are cloned twice - each with its own definition of T2. Kind of sounds like a
  dynamic way to define 2 implementations of the same interface. Powerful.

***

[^1]: Icon provided by [seekvectors.com](https://seekvectors.com/).
