package proxmox

type Node struct {
	Node      string  `json:"node"`
	Status    string  `json:"status"`
	CPU       float64 `json:"cpu"`
	MaxCPU    int     `json:"maxcpu"`
	Memory    uint64  `json:"mem"`
	MaxMemory uint64  `json:"maxmem"`
	Disk      uint64  `json:"disk"`
	MaxDisk   uint64  `json:"maxdisk"`
	Uptime    uint64  `json:"uptime"`
	Type      string  `json:"type"`
}
