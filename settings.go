package optimal

import (
	"fmt"

	env "github.com/daanv2/go-optimal/environment"
	"github.com/daanv2/go-optimal/internal/cpu"
)

func init() {
	tempTarget, found := env.Int64.LookupIf("CPU_CACHE_OPTIMAL_SIZE")
	if !found {
		tempTarget = OptimalBytesForCache(cpu.GetDefaultCacheTarget())
	}

	targetSize = int64(tempTarget)

	fmt.Println("==== // Go Optimal // ====")
	cpu.GetCPUInfo().Print()
	fmt.Println("Target cpu cache: ", cpu.GetDefaultCacheTarget())
	fmt.Println("Target size:      ", targetSize)
}

func OptimalBytesForCache(cache cpu.CacheKind) int64 {
	switch cache {
	case cpu.CacheL1:
		// 60% of the L1 cache
		size := cpu.GetCacheSize(cpu.CacheL1)
		return percentage(size, 60)

	case cpu.CacheL2:
		// 90% of the L2
		size := cpu.GetCacheSize(cpu.CacheL2)
		return percentage(size, 90)

	default:
		fallthrough
	case cpu.CacheL3:
		// 95% of the L3
		size := cpu.GetCacheSize(cpu.CacheL3)
		return percentage(size, 95)
	}
}

func percentage(value int64, percentage int64) int64 {
	return (value * percentage) / 100
}

// The amount of bytes data can maximally be
var targetSize int64

// GetTargetSize The target size in bytes of what can be used in a cache
func GetTargetSize() int64 {
	return targetSize
}

// SetTargetSize sets the target size in bytes of what can be used in a cache
func SetTargetSize(size int64) {
	targetSize = size
}
