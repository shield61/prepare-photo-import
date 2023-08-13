package scan

import (
	"fmt"
	"strings"
)

type ExtensionCount struct {
	Extension string
	Count     int
}



type ScanSummary struct {
	Files        int
	Directories  int
	FileCountMap map[string]int
}

func humanReadableSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

func extIncluded(ext string, extensions []string) bool {
	for _, e := range extensions {
		if strings.TrimPrefix(e, ".") == strings.TrimPrefix(ext, ".") {
			return true
		}
	}
	return false
}
