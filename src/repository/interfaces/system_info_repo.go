package interfaces

import (
	"github.com/Jefschlarski/monitor/src/application/interfaces"
)

type GetInfoInterface interface {
	interfaces.GetCpuUsageInterface
	interfaces.GetDiskUsageInterface
	interfaces.GetMemoryUsageInterface
	interfaces.GetMemorySwapUsageInterface
}
