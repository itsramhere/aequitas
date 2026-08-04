package telemetry

import (
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
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

type GCPauseSample struct {
	NumGC      uint32        `json:"num_gc"`
	PauseTotal time.Duration `json:"pause_total"`
	LastPause  time.Duration `json:"last_pause"`
}

type GCMonitor struct{}

func NewGCMonitor() *GCMonitor {
	return &GCMonitor{}
}

func (m *GCMonitor) Sample() GCPauseSample {
	var stats debug.GCStats
	debug.ReadGCStats(&stats)
	var lastPause time.Duration
	if len(stats.Pause) > 0 {
		lastPause = stats.Pause[0]
	}
	return GCPauseSample{
		NumGC:      uint32(stats.NumGC),
		PauseTotal: stats.PauseTotal,
		LastPause:  lastPause,
	}
}

func ParseEnvInt(key string, fallback int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return fallback
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return fallback
	}
	return val
}
