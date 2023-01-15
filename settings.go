package optimal

import (
	"fmt"

	env "github.com/daanv2/go-optimal/environment"
	"github.com/daanv2/go-optimal/internal/cpu"
)

func init() {
	tempTarget, found := env.Int64.LookupIf("CPU_CACHE_OPTIMAL_SIZE")
	if !found {
		switch cpu.GetDefaultCacheTarget() {
		case cpu.CacheL1:
			// 60% of the L1 cache
			tempTarget = percentage(cpu.GetCacheSize(cpu.CacheL1), 60)

		case cpu.CacheL2:
			// 90% of the L2, L3 cache
			tempTarget = percentage(cpu.GetCacheSize(cpu.CacheL1), 90)

		default:
			fallthrough
		case cpu.CacheL3:
			tempTarget = percentage(cpu.GetCacheSize(cpu.CacheL1), 95)
		}
	}

	targetSize = int64(tempTarget)

	fmt.Println("==== // Go Optimal // ====")
	cpu.GetCPUInfo().Print()
	fmt.Println("Target cpu cache: ", cpu.GetDefaultCacheTarget())
	fmt.Println("Target size:      ", targetSize)
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
