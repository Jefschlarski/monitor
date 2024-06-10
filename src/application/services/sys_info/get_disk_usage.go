package sysinfo

import (
	serviceInterface "github.com/Jefschlarski/monitor/src/application/interfaces"
	repoInterface "github.com/Jefschlarski/monitor/src/repository/interfaces"
	"github.com/Jefschlarski/monitor/src/structs"
)

type GetDiskUsage struct {
	repo repoInterface.GetInfoInterface
}

func NewGetDiskUsage(repo repoInterface.GetInfoInterface) serviceInterface.GetDiskUsageInterface {
	return &GetDiskUsage{repo}
}

func (g *GetDiskUsage) GetDiskUsage() (diskInfo *structs.DiskUsage, err error) {
	return g.repo.GetDiskUsage()
}
