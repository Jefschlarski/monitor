package dto

import (
	"github.com/Jefschlarski/monitor/src/structs"
)

type sysInfoDto struct {
	MemoryUsage *structs.MemoryUsage `json:"memory_usage"`
	CpuUsage    *[]structs.CpuUsage  `json:"cpu_usage"`
	DiskUsage   *structs.DiskUsage   `json:"disk_usage"`
	SwapUsage   *structs.SwapUsage   `json:"swap_usage"`
}

func NewInfoDto(memoryUsage *structs.MemoryUsage, cpuUsage *[]structs.CpuUsage, diskUsage *structs.DiskUsage, swapUsage *structs.SwapUsage) *sysInfoDto {
	swap := structs.NewSwapUsage(swapUsage.Total, swapUsage.Used, swapUsage.Free, swapUsage.UsedPercent)

	return &sysInfoDto{
		MemoryUsage: memoryUsage,
		CpuUsage:    cpuUsage,
		DiskUsage:   diskUsage,
		SwapUsage:   swap,
	}
}
