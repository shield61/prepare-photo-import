package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	scan "github.com/shield61/prepare-photo-import/internal/scan"
)

var dirsCmd = &cobra.Command{
	Use:   "dirs [path]",
	Short: "List all directories in the specified path",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		fmt.Println("Scanning directories in:", path)
		scan.ScanDirectories(path, &scanSummary)
		fmt.Printf("\nSummary: Scanned %d directories.\n", scanSummary.Directories)
	},
}

func init() {
	rootCmd.AddCommand(dirsCmd)
}
