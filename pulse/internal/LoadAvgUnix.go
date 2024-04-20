package internal

import (
	"fmt"
	"os"
	"strings"
)

type LoadAvg struct {
	Avg1  float64 `json:"avg1"`  // The average processor workload of the last minute
	Avg5  float64 `json:"avg5"`  // The average processor workload of the last 5 minutes
	Avg15 float64 `json:"avg15"` // The average processor workload of the last 15 minutes
}

func ReadLoadAverage() (loadAvg LoadAvg, err error) {
	file, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return LoadAvg{}, err
	}

	content := string(file[:len(file)])
	loadAvg = LoadAvg{}
	fields := strings.Fields(content)
	fmt.Println(fields)

	return loadAvg, nil
}
