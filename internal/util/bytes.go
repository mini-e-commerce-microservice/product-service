package util

func BytesToMB(bytes int64) float64 {
	return float64(bytes) / (1024 * 1024)
}
