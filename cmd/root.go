package cmd

import (
	"github.com/spf13/cobra"
	scan "github.com/shield61/prepare-photo-import/internal/scan"
)

var rootCmd = &cobra.Command{Use: "preppi"}
var scanSummary scan.ScanSummary

func Execute() {
	scanSummary.FileCountMap = make(map[string]int)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
