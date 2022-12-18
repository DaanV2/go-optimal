package cpu

import (
	env "github.com/daanv2/optimal/environment"
	"github.com/klauspost/cpuid/v2"
)

func init() {
	CPU := cpuid.CPU

	cpuinfo = &CPUData{
		BrandName: CPU.BrandName,
		Cache: CacheInfo{
			L1: env.Int64.Lookup("CPU_CACHE_L1", int64(CPU.Cache.L1I)),
			L2: env.Int64.Lookup("CPU_CACHE_L2", int64(CPU.Cache.L2)),
			L3: env.Int64.Lookup("CPU_CACHE_L3", int64(CPU.Cache.L3)),
		},
	}

	cpuinfo.Cache.CheckAndEstimate()
}

// cpuinfo is the CPU information
var cpuinfo *CPUData

// GetCPUInfo returns the CPU information
func GetCPUInfo() *CPUData {
	return cpuinfo
}

// CPUData contains the CPU information
type CPUData struct {
	// BrandName is the name of the CPU
	BrandName string
	// Cache contains the cache information
	Cache CacheInfo
}

// CacheInfo contains the cache information
type CacheInfo struct {
	// L1 is the size of the L1 cache in bytes
	L1 int64
	// L2 is the size of the L2 cache in bytes
	L2 int64
	// L3 is the size of the L3 cache in bytes
	L3 int64
}

// CheckAndEstimate checks if the cache sizes are valid and estimates them if they are not
func (c *CacheInfo) CheckAndEstimate() {
	if c.L1 <= 0 {
		c.L1 = 128_000 // 128KB is currently the default value for L1 cache size of a modern CPU
	}
	if c.L2 <= 0 {
		c.L2 = 8_388_608 // 8MB is currently the default value for L2 cache size of a modern CPU
	}
	if c.L3 <= 0 {
		c.L3 = 33_554_432 // 32MB is currently the default value for L3 cache size of a modern CPU
	}
}

// Println prints the CPU information
func (c *CPUData) Print() {
	println("CPU:", c.BrandName)
	println("Cache L1:", c.Cache.L1)
	println("Cache L2:", c.Cache.L2)
	println("Cache L3:", c.Cache.L3)
}
