### Coverage Report
linked_list.out
`$go test -coverprofile=linked_list_coverage`
`$go tool cover -html="linked_list_coverage"`

### Benchmark
if you just want to benchmark the specific series, you can pass the regexp
`go test -bench=ToArray`