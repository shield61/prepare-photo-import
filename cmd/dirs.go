package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var dirsCmd = &cobra.Command{
	Use:   "dirs [path]",
	Short: "List all directories in the specified path",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		fmt.Println("Scanning directories in:", path)
		scanDirectories(path, &scanSummary)
		fmt.Printf("\nSummary: Scanned %d directories.\n", scanSummary.Directories)
	},
}

func init() {
	rootCmd.AddCommand(dirsCmd)
}

func scanDirectories(path string, summary *ScanSummary) {
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
