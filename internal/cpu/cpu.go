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

var cpuinfo *CPUData

func GetCPUInfo() *CPUData {
	return cpuinfo
}

type CPUData struct {
	BrandName string
	Cache     CacheInfo
}

type CacheInfo struct {
	L1 int64
	L2 int64
	L3 int64
}

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
