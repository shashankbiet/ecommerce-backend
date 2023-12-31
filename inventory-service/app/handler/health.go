package handler

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

// AppStats represents application statistics
type AppStats struct {
	Uptime     string `json:"uptime"`
	HeapAlloc  uint64 `json:"heap_alloc"`  //This is the total amount of heap memory that has been allocated. It may contain non-freed memory by the garbage collector.
	HeapInuse  uint64 `json:"heap_inuse"`  //This is the total amount of heap memory that is still in use.
	TotalAlloc uint64 `json:"total_alloc"` //This is the cumulative quantity of allocated heap over the process execution.
	Sys        uint64 `json:"sys"`         // This is the total amount of memory that has been allocated from the OS.
	NumGC      uint32 `json:"num_gc"`      //This is the number of garbage collector completed cycles.
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	stats := getAppStats()

	// Convert stats to JSON
	statsJSON, err := json.Marshal(stats)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(statsJSON)
}

func getAppStats() AppStats {
	// Uptime
	uptime := time.Since(startTime).Round(time.Second).String()

	// Memory usage
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return AppStats{
		Uptime:     uptime,
		HeapAlloc:  m.HeapAlloc,
		HeapInuse:  m.HeapInuse,
		TotalAlloc: m.TotalAlloc,
		Sys:        m.Sys,
		NumGC:      m.NumGC,
	}
}

var startTime = time.Now()
