package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  `Display version and build information for mcp-scan`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("MCP Scan v1.0.0")
		color.White("SwampUP 2025 Security Edition")
		fmt.Println()
		color.Blue("Build: Go " + "1.25.0")
		color.Blue("Author: Security Team")
		color.Blue("License: MIT")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}