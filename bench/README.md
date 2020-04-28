# run test/benchmark for mem alloc

`go test -bench=. -memprofile=mem.out`

with -bench you can filter the concrete benchmark.
e.g.: if the name of the benchmark is: *BenchmarkTree* than can you *-bench=Tree*

*mem.out* is the name of the file, where the result is saved

# run pprof

`go tool pprof -alloc_objects -top -cum mem.out`

or for more details

`go tool pprof -alloc_objects -list Bench mem`

*--list Bench* means, filter for the function with the (part) name *Bench* (or other filter)

# is using the heap or the stack

`go run --gcflags "-m -m" main.go`

with the comment *escapes to heap* you can see, what is going to the heap