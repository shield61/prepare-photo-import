package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "preppi"}
var scanSummary ScanSummary

func Execute() {
	scanSummary.FileCountMap = make(map[string]int)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
