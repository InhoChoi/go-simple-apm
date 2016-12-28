package collector

import "github.com/shirou/gopsutil/cpu"
import rpc "github.com/inhochoi/go-simple-apm/go-simple-apm"

func GetCpuStatus() rpc.CpuStatus {
	info, _ := cpu.Info()
	cores, _ := cpu.Counts(true)
	percent, _ := cpu.Percent(0, true)
	return rpc.CpuStatus{Modelname: info[0].ModelName, Cores: int32(cores), Status: percent}
}
