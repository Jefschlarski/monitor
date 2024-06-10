package structs

type SwapUsage struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

func NewSwapUsage(total uint64, used uint64, free uint64, usedPercent float64) *SwapUsage {
	return &SwapUsage{
		Total:       total,
		Used:        used,
		Free:        free,
		UsedPercent: usedPercent,
	}
}

type CpuUsage struct {
	Cpu         uint    `json:"cpu"`
	UsedPercent float64 `json:"used_percent"`
}

func NewCpuUsage(index uint, usage float64) *CpuUsage {
	return &CpuUsage{
		Cpu:         index,
		UsedPercent: usage,
	}
}

type DiskUsage struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

func NewDiskUsage(total uint64, used uint64, free uint64, usedPercent float64) *DiskUsage {
	return &DiskUsage{
		Total:       total,
		Used:        used,
		Free:        free,
		UsedPercent: usedPercent,
	}
}

type MemoryUsage struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Available   uint64  `json:"available"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
	Cached      uint64  `json:"cached"`
}

func NewMemoryUsage(total uint64, used uint64, available uint64, free uint64, usedPercent float64, cached uint64) *MemoryUsage {
	return &MemoryUsage{
		Total:       total,
		Used:        used,
		Available:   available,
		Free:        free,
		UsedPercent: usedPercent,
		Cached:      cached,
	}
}
