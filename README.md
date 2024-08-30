# benchmark-go-json

A playground for benchmarking the performance of Go's JSON encoding and decoding.

Specifically, using NPM package metadata as a test case.

https://github.com/npm/registry/blob/main/docs/responses/package-metadata.md#abbreviated-metadata-format

Using a descently-sized JSON file as an input.

```shell
$ benchstat old.txt new.txt
goos: darwin
goarch: arm64
pkg: github.com/mfridman/benchmark-go-json
cpu: Apple M1 Max
                     │   old.txt   │                new.txt                │
                     │   sec/op    │    sec/op     vs base                 │
StandardLibrary-10     58.90m ± 2%   129.05m ± 1%  +119.10% (p=0.000 n=10)
StandardLibraryV2-10   56.83m ± 2%    81.52m ± 4%   +43.44% (p=0.000 n=10)
GoccyJSON-10           53.95m ± 6%    72.04m ± 5%   +33.54% (p=0.000 n=10)
SegmentJSON-10         36.33m ± 1%    89.16m ± 1%  +145.42% (p=0.000 n=10)
geomean                50.61m         90.67m        +79.15%

                     │    old.txt    │                  new.txt                  │
                     │     B/op      │      B/op        vs base                  │
StandardLibrary-10     17.83Mi ± 18%     15.24Mi ± 21%    -14.51% (p=0.000 n=10)
StandardLibraryV2-10   9.332Mi ±  0%   239.136Mi ±  3%  +2462.54% (p=0.000 n=10)
GoccyJSON-10           31.66Mi ± 20%     30.69Mi ± 54%          ~ (p=0.971 n=10)
SegmentJSON-10         9.500Mi ± 51%     7.692Mi ± 74%          ~ (p=0.670 n=10)
geomean                14.96Mi           30.46Mi         +103.63%

                     │     old.txt     │                new.txt                 │
                     │    allocs/op    │   allocs/op    vs base                 │
StandardLibrary-10     554654.000 ± 0%    4.000 ±  50%  -100.00% (p=0.000 n=10)
StandardLibraryV2-10      183.46k ± 0%   41.20k ±   0%   -77.55% (p=0.000 n=10)
GoccyJSON-10             30585.00 ± 0%    10.00 ±  90%   -99.97% (p=0.000 n=10)
SegmentJSON-10         262043.000 ± 0%    1.000 ± 100%  -100.00% (p=0.000 n=10)
geomean                    169.0k         35.83          -99.98%
```
