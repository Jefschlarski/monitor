package interfaces

import (
	"github.com/Jefschlarski/monitor/src/structs"
)

type GetMemoryUsageInterface interface {
	GetMemoryUsage() (memUsage *structs.MemoryUsage, err error)
}

type GetMemorySwapUsageInterface interface {
	GetMemorySwapUsage() (swapUsage *structs.SwapUsage, err error)
}

type GetCpuUsageInterface interface {
	GetCpuUsage() (cpuInfo *[]structs.CpuUsage, err error)
}

type GetDiskUsageInterface interface {
	GetDiskUsage() (diskInfo *structs.DiskUsage, err error)
}
