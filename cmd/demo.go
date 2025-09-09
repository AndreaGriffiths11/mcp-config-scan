package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Run a demo scan on sample configurations",
	Long: `Run mcp-scan against the included demo configurations to see 
the tool in action. This is perfect for presentations and testing.`,
	Run: runDemo,
}

func init() {
	rootCmd.AddCommand(demoCmd)
}

func runDemo(cmd *cobra.Command, args []string) {
	color.Cyan("🎭 MCP SCAN DEMO MODE")
	color.White("Running security scan on demo configurations...")
	fmt.Println()

	// Check if demos directory exists
	demoDir := "./demos"
	if _, err := os.Stat(demoDir); os.IsNotExist(err) {
		color.Red("❌ Demo directory not found: %s", demoDir)
		color.Yellow("💡 Make sure you're running from the project root directory")
		return
	}

	// Count demo files
	files, err := filepath.Glob(filepath.Join(demoDir, "*.json"))
	if err != nil {
		color.Red("❌ Error scanning demo directory: %v", err)
		return
	}

	if len(files) == 0 {
		color.Yellow("⚠️  No demo configuration files found in %s", demoDir)
		return
	}

	color.Green("📁 Found %d demo configurations", len(files))
	fmt.Println()

	// Set demo-specific flags and run scan
	configPaths = files
	verbose = true
	
	// Call the main scan function
	runScan(cmd, args)

	// Demo conclusion
	fmt.Println()
	color.Cyan("🎯 DEMO COMPLETE")
	color.White("This demo showed various MCP security issues including:")
	color.Red("  • Exposed API keys and secrets")
	color.Red("  • Dangerous filesystem access patterns") 
	color.Red("  • Command injection vulnerabilities")
	color.Red("  • Insecure configuration settings")
	fmt.Println()
}