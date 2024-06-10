package repo

import (
	"fmt"
	"time"

	"github.com/Jefschlarski/monitor/src/repository/interfaces"
	"github.com/Jefschlarski/monitor/src/structs"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemInfoRepo struct{}

func NewSystemInfoRepo() interfaces.GetInfoInterface {
	return &SystemInfoRepo{}
}

func (s *SystemInfoRepo) GetMemoryUsage() (memUsage *structs.MemoryUsage, err error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Erro ao obter uso da memória:", err)
		return
	}

	memUsage = structs.NewMemoryUsage(memInfo.Total, memInfo.Used, memInfo.Available, memInfo.Free, memInfo.UsedPercent, memInfo.Cached)

	return
}

func (s *SystemInfoRepo) GetMemorySwapUsage() (swapUsage *structs.SwapUsage, err error) {
	swapInfo, err := mem.SwapMemory()
	if err != nil {
		fmt.Println("Erro ao obter uso da swap:", err)
		return
	}

	swapUsage = structs.NewSwapUsage(swapInfo.Total, swapInfo.Used, swapInfo.Free, swapInfo.UsedPercent)

	return
}

func (s *SystemInfoRepo) GetCpuUsage() (cpuPercent *[]structs.CpuUsage, err error) {
	cpuInfos, err := cpu.Percent(time.Second, true)

	if err != nil {
		fmt.Println("Erro ao obter uso da CPU:", err)
		return
	}

	cpuPercent = &[]structs.CpuUsage{}
	for _, cpuInfo := range cpuInfos {
		*cpuPercent = append(*cpuPercent, *structs.NewCpuUsage(uint(len(*cpuPercent)), cpuInfo))
	}

	return
}

func (s *SystemInfoRepo) GetDiskUsage() (diskUsage *structs.DiskUsage, err error) {
	diskInfo, err := disk.Usage(".")
	if err != nil {
		fmt.Println("Erro ao obter uso do disco:", err)
		return
	}

	diskUsage = structs.NewDiskUsage(diskInfo.Total, diskInfo.Used, diskInfo.Free, diskInfo.UsedPercent)

	return
}
