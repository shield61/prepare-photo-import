package scan

import (
	"fmt"
	"os"
	"path/filepath"
)


func ScanDirectories(path string, summary *ScanSummary) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			summary.Directories++
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error scanning directory: %s\n", err)
	}
}