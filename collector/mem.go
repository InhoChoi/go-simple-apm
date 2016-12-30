package collector

import (
	rpc "github.com/inhochoi/go-simple-apm/go-simple-apm"
	"github.com/shirou/gopsutil/mem"
)

func GetMemoryStatus() rpc.MemoryStatus {
	v, _ := mem.VirtualMemory()
	return rpc.MemoryStatus{Total: uint64(v.Total), Used: uint64(v.Used), Free: uint64(v.Free), UsedPercent: v.UsedPercent}
}
