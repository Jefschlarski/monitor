package sysinfo

import (
	serviceInterface "github.com/Jefschlarski/monitor/src/application/interfaces"
	repoInterface "github.com/Jefschlarski/monitor/src/repository/interfaces"
	"github.com/Jefschlarski/monitor/src/structs"
)

type GetMemorySwapUsage struct {
	repo repoInterface.GetInfoInterface
}

func NewGetMemorySwapUsage(repo repoInterface.GetInfoInterface) serviceInterface.GetMemorySwapUsageInterface {
	return &GetMemorySwapUsage{repo}
}

func (g *GetMemorySwapUsage) GetMemorySwapUsage() (swapUsage *structs.SwapUsage, err error) {
	return g.repo.GetMemorySwapUsage()
}
