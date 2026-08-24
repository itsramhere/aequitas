package telemetry

import (
	"os"
	"runtime"
)

type GCRuntimeSettings struct {
	GOGC       string `json:"gogc"`
	GOMAXPROCS int    `json:"gomaxprocs"`
	NumCPU     int    `json:"num_cpu"`
}

func GetGCRuntimeSettings() GCRuntimeSettings {
	gogc := os.Getenv("GOGC")
	if gogc == "" {
		gogc = "100 (default)"
	}
	return GCRuntimeSettings{
		GOGC:       gogc,
		GOMAXPROCS: runtime.GOMAXPROCS(0),
		NumCPU:     runtime.NumCPU(),
	}
}
