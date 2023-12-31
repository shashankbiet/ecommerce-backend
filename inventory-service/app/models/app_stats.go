package model

// AppStats represents application statistics
type AppStats struct {
	Uptime        string `json:"uptime"`
	MemoryUsage   uint64 `json:"memoryUsage"`
	TotalMemory   uint64 `json:"TotalMemory"`
	NumGoroutines int    `json:"numGoroutines"`
}
