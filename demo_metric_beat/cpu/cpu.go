package cpu

import (
	"demo_metric-beat/util"
	"fmt"
	"strconv"
	"strings"
)

func PrintMetric() {
	m := &Monitor{}
	sample, err := m.getSample()
	if err != nil {
		panic(err)
	}
	fmt.Println(sample)
}

func (m *Monitor) getSample() (*Metrics, error) {
	cpuSample := &Cpu{}
	if err := cpuSample.Get(); err != nil {
		return nil, err
	}
	oldLastSample := m.lastSample
	m.lastSample = cpuSample
	return &Metrics{oldLastSample, cpuSample}, nil
}

// Monitor is used to monitor the overall CPU usage of the system.
type Monitor struct {
	lastSample *Cpu
}

// Metrics stores the current and the last sample collected by a Beat.
type Metrics struct {
	previousSample *Cpu
	currentSample  *Cpu
}

// Cpu contains CPU time stats
type Cpu struct {
	User uint64
	Nice uint64
	Sys  uint64
	Idle uint64
	Wait uint64
	// Irq interrupted
	Irq     uint64
	SoftIrq uint64
	Stolen  uint64
}

func (self *Cpu) Get() error {
	return util.ReadFile("/proc/stat", func(line string) bool {
		if len(line) > 4 && line[0:4] == "cpu " {
			parseCpuStat(self, line)
			return false
		}
		return true
	})
}

func parseCpuStat(self *Cpu, line string) error {
	fields := strings.Fields(line)

	self.User, _ = strtoull(fields[1])
	self.Nice, _ = strtoull(fields[2])
	self.Sys, _ = strtoull(fields[3])
	self.Idle, _ = strtoull(fields[4])
	self.Wait, _ = strtoull(fields[5])
	self.Irq, _ = strtoull(fields[6])
	self.SoftIrq, _ = strtoull(fields[7])
	self.Stolen, _ = strtoull(fields[8])

	return nil
}

func strtoull(val string) (uint64, error) {
	return strconv.ParseUint(val, 10, 64)
}
