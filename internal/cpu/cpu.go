package cpu

import (
	"fmt"
	"runtime"

	"github.com/klauspost/cpuid/v2"
	env "github.com/optimal/environment"
)

func init() {
	CPU := cpuid.CPU

	cpuinfo = CPUData{
		BrandName: CPU.BrandName,
		Cache: CacheInfo{
			L1: env.Int64.Lookup("CPU_CACHE_L1", int64(CPU.Cache.L1I)),
			L2: env.Int64.Lookup("CPU_CACHE_L2", int64(CPU.Cache.L2)),
			L3: env.Int64.Lookup("CPU_CACHE_L3", int64(CPU.Cache.L3)),
		},
	}

	cpuinfo.Cache.CheckAndEstimate()

	fmt.Printf("cpu: %s\n", cpuinfo.BrandName)
	fmt.Printf("cache L1: %v\n", cpuinfo.Cache.L1)
	fmt.Printf("cache L2: %v\n", cpuinfo.Cache.L2)
	fmt.Printf("cache L3: %v\n", cpuinfo.Cache.L3)
	fmt.Printf("concurrency: %v\n", runtime.GOMAXPROCS(0))
}

var cpuinfo CPUData

func GetCPUInfo() CPUData {
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
