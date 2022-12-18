package optimal

import (
	env "github.com/daanv2/optimal/environment"
	"github.com/daanv2/optimal/internal/cpu"
)

func init() {
	tempTarget, ok := env.Int64.LookupIf("CPU_CACHE_OPTIMAL_SIZE")
	if !ok {
		switch cpu.GetDefaultCacheTarget() {
		case cpu.CacheL1:
			// 60% of the L1 cache
			tempTarget = (cpu.GetCacheSize(cpu.CacheL1) * 100) / 60

		default:
			fallthrough
		case cpu.CacheL2, cpu.CacheL3:
			// 90% of the L2, L3 cache
			tempTarget = (cpu.GetCacheSize(cpu.CacheL2) * 100) / 90
		}
	}

	targetSize = int64(tempTarget)
}

// The amount of bytes data can maximally be
var targetSize int64

// The target size in bytes of what can be used in a cache
func GetTargetSize() int64 {
	return targetSize
}

// SetTargetSize sets the target size in bytes of what can be used in a cache
func SetTargetSize(size int64) {
	targetSize = size
}
