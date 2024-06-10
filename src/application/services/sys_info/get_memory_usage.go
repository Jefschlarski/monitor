package sysinfo

import (
	serviceInterface "github.com/Jefschlarski/monitor/src/application/interfaces"
	repoInterface "github.com/Jefschlarski/monitor/src/repository/interfaces"
	"github.com/Jefschlarski/monitor/src/structs"
)

type GetMemoryUsage struct {
	repo repoInterface.GetInfoInterface
}

func NewGetMemoryUsage(repo repoInterface.GetInfoInterface) serviceInterface.GetMemoryUsageInterface {
	return &GetMemoryUsage{repo}
}

func (g *GetMemoryUsage) GetMemoryUsage() (memUsage *structs.MemoryUsage, err error) {
	return g.repo.GetMemoryUsage()
}
