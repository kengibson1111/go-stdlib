package main

import (
	"cmp"
	"slices"
	"sync/atomic"
	"testing"
)

func main() {
	// This reports a custom benchmark metric relevant to a
	// specific algorithm (in this case, sorting) in parallel.
	testing.Benchmark(func(b *testing.B) {
		var compares atomic.Int64
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				s := []int{5, 4, 3, 2, 1}
				slices.SortFunc(s, func(a, b int) int {
					// Because RunParallel runs the function many
					// times in parallel, we must increment the
					// counter atomically to avoid racing writes.
					compares.Add(1)
					return cmp.Compare(a, b)
				})
			}
		})

		// NOTE: Report each metric once, after all of the parallel
		// calls have completed.

		// This metric is per-operation, so divide by b.N and
		// report it as a "/op" unit.
		b.ReportMetric(float64(compares.Load())/float64(b.N), "compares/op")

		// This metric is per-time, so divide by b.Elapsed and
		// report it as a "/ns" unit.
		b.ReportMetric(float64(compares.Load())/float64(b.Elapsed().Nanoseconds()), "compares/ns")
	})
}
