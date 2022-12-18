package cpu

import env "github.com/daanv2/optimal/environment"

type CacheKind int

const (
	CacheL1 CacheKind = 0
	CacheL2 CacheKind = 1
	CacheL3 CacheKind = 2
)

func (ck CacheKind) GetCacheSize() int64 {
	return GetCacheSize(ck)
}

func (ck CacheKind) String() string {
	switch ck {
	case CacheL1:
		return "L1"
	case CacheL2:
		return "L2"
	case CacheL3:
		return "L3"
	}

	return "unknown"
}

// GetCacheSize returns the size of the cache.
func GetCacheSize(kind CacheKind) int64 {
	cacheInfo := GetCPUInfo().Cache

	switch kind {
	case CacheL1:
		return cacheInfo.L1
	case CacheL2:
		return cacheInfo.L2
	case CacheL3:
		return cacheInfo.L3
	default:
		return cacheInfo.L3
	}
}

// GetDefaultCacheTarget returns the default cache target.
func GetDefaultCacheTarget() CacheKind {
	value := env.String.Lookup("CPU_CACHE_TARGET", "NONE")

	switch value {
	case "L1", "l1":
		return CacheL1
	case "L2", "l2":
		return CacheL2
	case "L3", "l3":
		return CacheL3
	}

	// Determine the target from the cache size
	result := CacheL1
	cpu := GetCPUInfo()

	if cpu.Cache.L1 <= 0 {
		result = CacheL2
	}
	if cpu.Cache.L2 <= 0 {
		result = CacheL3
	}

	return result
}
