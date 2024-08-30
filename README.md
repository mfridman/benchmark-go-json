# benchmark-go-json

A playground for benchmarking the performance of Go's JSON encoding and decoding.

Specifically, using NPM package metadata as a test case.

https://github.com/npm/registry/blob/main/docs/responses/package-metadata.md#abbreviated-metadata-format

Using a descently-sized JSON file as an input.

```shell
go test -bench=. -benchmem -count=1 -v -benchtime=2s
goos: darwin
goarch: arm64
pkg: github.com/mfridman/benchmark-go-json
cpu: Apple M1 Max
BenchmarkStandardLibrary
BenchmarkStandardLibrary-10                   39          60079762 ns/op        18696526 B/op     554654 allocs/op
BenchmarkStandardLibraryV2
BenchmarkStandardLibraryV2-10                 36          58147962 ns/op         9785114 B/op     183463 allocs/op
BenchmarkGoccyJSON
BenchmarkGoccyJSON-10                         46          51471562 ns/op        32189440 B/op      30585 allocs/op
BenchmarkSegmentJSON
BenchmarkSegmentJSON-10                       66          36569542 ns/op         7373434 B/op     262043 allocs/op
PASS
ok      github.com/mfridman/benchmark-go-json   14.512s
```
