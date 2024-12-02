# runtime/pprof - havlak2

This example shows how to create CPU and memory profiles. The code originated from the [benchgraffiti tests](https://github.com/rsc/benchgraffiti/blob/master/havlak/havlak2.go) mentioned [here](https://go.dev/blog/pprof). I added the memory profile based on Go Standard Library documentation.

## run the program and generate profile files

```bash
go install
gostdlib-runtime-pprof-havlak2 -cpuprofile=havlak2-cpu.prof
gostdlib-runtime-pprof-havlak2 -memprofile=havlak2-mem.prof
```

## run the profiler

For the CPU profile:

```bash
go tool pprof $GOPATH/bin/gostdlib-runtime-pprof-havlak2 havlak2-cpu.prof
```

For the memory profile with size stats:

```bash
go tool pprof $GOPATH/bin/gostdlib-runtime-pprof-havlak2 havlak2-mem.prof
```

For the memory profile with allocation-count stats:

```bash
go tool pprof --inuse_objects $GOPATH/bin/gostdlib-runtime-pprof-havlak2 havlak2-mem.prof
```

When starting the pprof tool, another useful flag for `go tool pprof` is `--nodefraction`. It culls out functions that don't have a certain percentage of samples. For example, to cull out functions that don't have at least 10% of the total samples, `--nodefraction=0.1`.

In the pprof tool, type `help` to see available commands.

## popular commands

* **topN** - lists the top N samples in the profile. `top` defaults to top 10 samples.

* **web** - creates a graph of the profile data in SVG format and opens a browser. When no arguments are passed, all functions are in the graph. To filter the graph, pass a function name as an argument.

* **list <function name>** - shows function code and the number of samples per line of code taken.

* **disasm <function name>** - shows the disassembly of a function and the number of samples per line of disassembly instruction taken.

* **weblist <function name>** - shows a combination of `list` and `disasm`. When clicking a line of code, the disassembly is shown for that line.
