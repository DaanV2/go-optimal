# Go Optimal

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/DaanV2/go-optimal)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/DaanV2/go-optimal)
[![🐹 Golang](https://github.com/daanv2/go-optimal/actions/workflows/go-checks.yml/badge.svg)](https://github.com/daanv2/go-optimal/actions/workflows/go-checks.yml)

A simple, fast, and easy-to-use library for optimally creating and handling data.

```bash
go get github.com/daanv2/go-optimal
```

## Documentation

The documentation can be found [go dev](https://pkg.go.dev/github.com/daanv2/go-optimal) or [here](https://github.com/daanv2/go-optimal/tree/main/doc).

# Environment Variables / Settings

| Variable                 | Optional | Documentation                                                                                   |
| ------------------------ | -------- | ----------------------------------------------------------------------------------------------- |
| `CPU_CACHE_L1`           | `true`   | Cache L1 size in bytes, When set overrides the scraped value                                    |
| `CPU_CACHE_L2`           | `true`   | Cache L2 size in bytes, When set overrides the scraped value                                    |
| `CPU_CACHE_L3`           | `true`   | Cache L3 size in bytes, When set overrides the scraped value                                    |
| `CPU_CACHE_OPTIMAL_SIZE` | `true`   | Optimal cache size in bytes, When set overrides the scraped value                               |
| `CPU_CACHE_TARGET`       | `true`   | What type of cache to target, can be L1, L2, L3, anything else and it will determine on its own |
