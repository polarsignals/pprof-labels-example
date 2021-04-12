package main

import (
	"context"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	// Setup CPU profiling.
	f, err := os.Create("cpuprofile.pb.gz")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	ctx := context.Background()
	// Run workload for tenant1
	iteratePerTenant(ctx, "tenant1")
	// Run workload for tenant2
	iteratePerTenant(ctx, "tenant2")
}

var (
	iterationsPerTenant = map[string]int{
		"tenant1": 10_000_000_000,
		"tenant2": 1_000_000_000,
	}
)

func iteratePerTenant(ctx context.Context, tenant string) {
	// pprof.Do instruments the CPU profiler to differentiate function call
	// stacks by unique labels.
	pprof.Do(ctx, pprof.Labels("tenant", tenant), func(ctx context.Context) {
		iterate(iterationsPerTenant[tenant])
	})
}

func iterate(iterations int) {
	for i := 0; i < iterations; i++ {
	}
}
