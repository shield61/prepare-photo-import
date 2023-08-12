package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var extFlags []string

var filesCmd = &cobra.Command{
	Use:   "files [path]",
	Short: "List files with specified extensions in the specified path",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		fmt.Println("Scanning files in:", path)
		scanFiles(path, &scanSummary, extFlags)
		fmt.Printf("\nSummary: Scanned %d files.\n", scanSummary.Files)
		printFileExtensionSummary(extFlags, &scanSummary)
	},
}

func init() {
	filesCmd.Flags().StringArrayVarP(&extFlags, "ext", "e", nil, "Include files with specified extension")
	rootCmd.AddCommand(filesCmd)
}

func scanFiles(path string, summary *ScanSummary, extensions []string) {
	totalSize := int64(0)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			extension := filepath.Ext(info.Name())
			lowercaseExtension := strings.ToLower(extension)
			if extIncluded(lowercaseExtension, extensions) {
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

func printFileExtensionSummary(extensions []string, summary *ScanSummary) {
	if len(extensions) > 0 {
		fmt.Println("File Extension Summary:")
		for _, ext := range extensions {
			count := summary.FileCountMap[ext]
			fmt.Printf("Summary for %s files: %d\n", ext, count)
		}
	}
}
