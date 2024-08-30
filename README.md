# benchmark-go-json

A playground for benchmarking the performance of Go's JSON encoding and decoding.

Specifically, using NPM package metadata as a test case.

https://github.com/npm/registry/blob/main/docs/responses/package-metadata.md#abbreviated-metadata-format

Using a descently-sized JSON file as an input.

```shell
benchstat old.txt new.txt
goos: darwin
goarch: arm64
pkg: github.com/mfridman/benchmark-go-json
cpu: Apple M1 Max
                     │   old.txt   │                new.txt                │
                     │   sec/op    │    sec/op     vs base                 │
StandardLibrary-10     58.94m ± 1%   128.71m ± 4%  +118.38% (p=0.000 n=10)
StandardLibraryV2-10   57.75m ± 3%    81.42m ± 3%   +40.99% (p=0.000 n=10)
GoccyJSON-10           50.83m ± 8%    68.55m ± 2%   +34.85% (p=0.000 n=10)
SegmentJSON-10         36.46m ± 0%    88.36m ± 0%  +142.32% (p=0.000 n=10)
geomean                50.12m         89.26m        +78.10%

                     │    old.txt    │                  new.txt                   │
                     │     B/op      │       B/op        vs base                  │
StandardLibrary-10     21.03Mi ± 15%     15.24Mi ± 157%          ~ (p=0.137 n=10)
StandardLibraryV2-10   9.332Mi ±  0%   238.565Mi ±   2%  +2456.42% (p=0.000 n=10)
GoccyJSON-10           24.72Mi ± 56%     28.70Mi ±  79%          ~ (p=0.894 n=10)
SegmentJSON-10         9.531Mi ±  2%     9.772Mi ±  80%          ~ (p=0.959 n=10)
geomean                14.66Mi           31.78Mi          +116.72%

                     │     old.txt     │                new.txt                 │
                     │    allocs/op    │   allocs/op    vs base                 │
StandardLibrary-10     554656.000 ± 0%    4.000 ± 250%  -100.00% (p=0.000 n=10)
StandardLibraryV2-10      183.46k ± 0%   41.19k ±   0%   -77.55% (p=0.000 n=10)
GoccyJSON-10            30583.000 ± 0%    8.000 ±  88%   -99.97% (p=0.000 n=10)
SegmentJSON-10         262044.000 ± 0%    2.000 ±  50%  -100.00% (p=0.000 n=10)
geomean                    169.0k         40.30          -99.98%
```
