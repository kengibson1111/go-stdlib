# net/http/pprof

No code examples of this package, but here are tool examples assuming an HTTP server running on port 6060 that has the following import statement:

```
import _ "net/http/pprof"
```

Use the pprof tool to look at the heap profile:

```bash
go tool pprof http://localhost:6060/debug/pprof/heap
```

Or to look at a 30-second CPU profile:

```bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

Or to look at the goroutine blocking profile, after calling runtime.SetBlockProfileRate in your program:

```bash
go tool pprof http://localhost:6060/debug/pprof/block
```

Or to look at the holders of contended mutexes, after calling runtime.SetMutexProfileFraction in your program:

```bash
go tool pprof http://localhost:6060/debug/pprof/mutex
```

The package also exports a handler that serves execution trace data for the "go tool trace" command. To collect a 5-second execution trace:

```bash
curl -o trace.out http://localhost:6060/debug/pprof/trace?seconds=5
go tool trace trace.out
```

To view all available profiles, open [http://localhost:6060/debug/pprof/](http://localhost:6060/debug/pprof/) in your browser.

For a study of the facility in action, visit [https://blog.golang.org/2011/06/profiling-go-programs.html](https://blog.golang.org/2011/06/profiling-go-programs.html).
