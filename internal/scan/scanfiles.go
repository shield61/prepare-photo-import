package scan

import (
	"fmt"
	"os"
	"path/filepath"
)

func ScanFiles(path string, summary *ScanSummary, extensions []string) {
	totalSize := int64(0)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			extension := filepath.Ext(info.Name())
			if extIncluded(extension, extensions) {
				totalSize += info.Size()
				summary.Files++
			}
			summary.FileCountMap[extension]++
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error scanning directory: %s\n", err)
	}

	fmt.Printf("\nTotal Size of Matching Files: %s\n", humanReadableSize(totalSize))
}

func PrintFileExtensionSummary(extensions []string, summary *ScanSummary) {
	if len(extensions) > 0 {
		fmt.Println("File Extension Summary:")
		for _, ext := range extensions {
			count := summary.FileCountMap[ext]
			fmt.Printf("Summary for %s files: %d\n", ext, count)
		}
	}
}
