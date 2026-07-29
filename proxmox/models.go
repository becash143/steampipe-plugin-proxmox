package proxmox

type Node struct {
	Node      string  `json:"node"`
	Status    string  `json:"status"`
	CPU       float64 `json:"cpu"`
	MaxCPU    int     `json:"max_cpu"`
	Memory    uint64  `json:"mem"`
	MaxMemory uint64  `json:"max_mem"`
	Disk      uint64  `json:"disk"`
	MaxDisk   uint64  `json:"max_disk"`
	Uptime    uint64  `json:"uptime"`
	Type      string  `json:"type"`
}

type VM struct {
	Node    string `json:"-"`
	VMID    int    `json:"vm_id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	CPUs    int    `json:"cpus"`
	Mem     uint64 `json:"mem"`
	MaxMem  uint64 `json:"max_mem"`
	MaxDisk uint64 `json:"max_disk"`
	Uptime  uint64 `json:"uptime"`
}

type Container struct {
	Node    string `json:"-"`
	VMID    int    `json:"vm_id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	CPUs    int    `json:"cpus"`
	Mem     uint64 `json:"mem"`
	MaxMem  uint64 `json:"max_mem"`
	MaxDisk uint64 `json:"max_disk"`
	Uptime  uint64 `json:"uptime"`
}

type ClusterResource struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Node    string  `json:"node"`
	VMID    int     `json:"vmid"`
	Name    string  `json:"name"`
	Status  string  `json:"status"`
	CPU     float64 `json:"cpu"`
	MaxCPU  int     `json:"max_cpu"`
	Mem     uint64  `json:"mem"`
	MaxMem  uint64  `json:"max_mem"`
	Disk    uint64  `json:"disk"`
	MaxDisk uint64  `json:"max_disk"`
	Uptime  uint64  `json:"uptime"`
	Pool    string  `json:"pool"`
}

type Storage struct {
	Node    string `json:"-"`
	Storage string `json:"storage"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Active  int    `json:"active"`
	Used    uint64 `json:"used"`
	Avail   uint64 `json:"avail"`
	Total   uint64 `json:"total"`
	Shared  int    `json:"shared"`
}

type NetworkInterface struct {
	Node      string `json:"-"`
	Iface     string `json:"iface"`
	Type      string `json:"type"`
	Active    int    `json:"active"`
	Autostart int    `json:"autostart"`
	Address   string `json:"address"`
	Netmask   string `json:"netmask"`
	Gateway   string `json:"gateway"`
	Method    string `json:"method"`
}

type Task struct {
	Node      string `json:"-"`
	UPID      string `json:"upid"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	User      string `json:"user"`
	PID       int    `json:"pid"`
	StartTime int64  `json:"starttime"`
	EndTime   int64  `json:"endtime"`
}

type User struct {
	UserID  string `json:"userid"`
	Enable  int    `json:"enable"`
	Expire  int64  `json:"expire"`
	Email   string `json:"email"`
	Comment string `json:"comment"`
}

type Pool struct {
	PoolID  string `json:"poolid"`
	Comment string `json:"comment"`
}
