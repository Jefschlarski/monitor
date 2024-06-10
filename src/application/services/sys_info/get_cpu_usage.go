package sysinfo

import (
	serviceInterface "github.com/Jefschlarski/monitor/src/application/interfaces"
	repoInterface "github.com/Jefschlarski/monitor/src/repository/interfaces"
	"github.com/Jefschlarski/monitor/src/structs"
)

type GetCpuUsage struct {
	repo repoInterface.GetInfoInterface
}

func NewGetCpuUsage(repo repoInterface.GetInfoInterface) serviceInterface.GetCpuUsageInterface {
	return &GetCpuUsage{repo}
}

func (g *GetCpuUsage) GetCpuUsage() (cpuInfo *[]structs.CpuUsage, err error) {
	return g.repo.GetCpuUsage()
}
