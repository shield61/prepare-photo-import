package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/shield61/prepare-photo-import/internal/scan"
	
)

var extFlags []string

var filesCmd = &cobra.Command{
	Use:   "files [path]",
	Short: "List files with specified extensions in the specified path",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		fmt.Println("Scanning files in:", path)
		ScanFiles(path, &scanSummary, extFlags)
		fmt.Printf("\nSummary: Scanned %d files.\n", scanSummary.Files)
		PrintFileExtensionSummary(extFlags, &scanSummary)
	},
}

func init() {
	filesCmd.Flags().StringArrayVarP(&extFlags, "ext", "e", nil, "Include files with specified extension")
	rootCmd.AddCommand(filesCmd)
}