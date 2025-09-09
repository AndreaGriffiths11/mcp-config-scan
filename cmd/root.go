package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"mcp-scan/pkg/config"
	"mcp-scan/pkg/report"
	"mcp-scan/pkg/scanner"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	outputFormat string
	outputFile   string
	configPaths  []string
	verbose      bool
	compact      bool
	maskSecrets  bool
)

var rootCmd = &cobra.Command{
	Use:   "mcp-scan",
	Short: "A security scanner for MCP (Model Context Protocol) configurations",
	Long: `MCP Scan is an EXPERIMENTAL security tool for educational purposes that analyzes 
MCP configuration files for potential security vulnerabilities including exposed secrets, 
unsafe filesystem access patterns, and insecure server configurations.

⚠️  WARNING: EDUCATIONAL USE ONLY - NOT FOR PRODUCTION SECURITY DECISIONS
⚠️  May produce false positives or miss real vulnerabilities
⚠️  Always verify findings manually and consult security professionals

Perfect for learning about MCP security concepts and demonstration purposes.`,
	Run: runScan,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&outputFormat, "format", "f", "console", "Output format: console, json")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file (default: stdout)")
	rootCmd.Flags().StringSliceVarP(&configPaths, "config", "c", []string{}, "Additional config paths to scan")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output with detailed explanations")
	rootCmd.Flags().BoolVarP(&compact, "compact", "q", false, "Enable compact output format (less visual, more text-focused)")
	rootCmd.Flags().BoolVarP(&maskSecrets, "mask-secrets", "m", false, "Mask/redact sensitive values in output")
}

func runScan(cmd *cobra.Command, args []string) {
	// Print banner
	printBanner()

	// Discover config files
	configs, err := discoverConfigs()
	if err != nil {
		color.Red("Error discovering configs: %v", err)
		os.Exit(1)
	}

	if len(configs) == 0 {
		color.Yellow("⚠️  No MCP configuration files found")
		color.White("Try specifying paths with --config or create some demo configs")
		return
	}

	// Print color legend if not in compact mode
	if !compact && verbose {
		printColorLegend()
	}

	// Scan configurations
	var allResults []scanner.ScanResult
	totalIssues := 0
	issuesBySeverity := map[string]int{
		"critical": 0,
		"high":     0,
		"medium":   0,
		"low":      0,
	}

	// Show progress during scanning
	fmt.Println()
	color.Cyan("🔍 Scanning %d configuration files...", len(configs))

	// Don't show progress bar in compact mode
	showProgress := !compact

	for i, configPath := range configs {
		// Show progress percentage
		if showProgress {
			progress := float64(i) / float64(len(configs)) * 100
			fmt.Printf("\r   Progress: [")
			progressWidth := 30
			completedWidth := int(float64(progressWidth) * progress / 100)

			for j := 0; j < progressWidth; j++ {
				if j < completedWidth {
					fmt.Print("█")
				} else {
					fmt.Print("░")
				}
			}
			fmt.Printf("] %.1f%%", progress)
		}

		mcpConfig, err := config.LoadMCPConfig(configPath)
		if err != nil {
			if showProgress {
				fmt.Println()
			}
			color.Red("❌ Error loading %s: %v", configPath, err)
			continue
		}

		result := scanner.ScanConfig(configPath, mcpConfig, maskSecrets)
		allResults = append(allResults, result)
		totalIssues += len(result.Issues)

		// Count issues by severity
		for _, issue := range result.Issues {
			severity := strings.ToLower(issue.Severity)
			issuesBySeverity[severity]++
		}
	}

	// Complete progress bar
	if showProgress {
		fmt.Printf("\r   Progress: [")
		for j := 0; j < 30; j++ {
			fmt.Print("█")
		}
		fmt.Printf("] 100.0%%\n\n")
	}

	// Print detailed results
	if outputFormat == "console" {
		for _, result := range allResults {
			printConsoleResult(result)
		}

		// Generate summary with severity counts
		printSummary(len(configs), totalIssues, issuesBySeverity)

		// Show help text for additional options
		if !compact && totalIssues > 0 {
			printHelpText()
		}
	}

	// Generate JSON report if requested
	if outputFormat == "json" {
		jsonReport := report.GenerateJSONReport(allResults)
		output, err := json.MarshalIndent(jsonReport, "", "  ")
		if err != nil {
			color.Red("Error generating JSON report: %v", err)
			os.Exit(1)
		}

		if outputFile != "" {
			err = os.WriteFile(outputFile, output, 0644)
			if err != nil {
				color.Red("Error writing to file: %v", err)
				os.Exit(1)
			}
			color.Green("📄 Report saved to: %s", outputFile)
		} else {
			fmt.Println(string(output))
		}
	}

	// Exit with error code if issues found
	if totalIssues > 0 {
		os.Exit(1)
	}
}

// Print color legend explaining what each color means
func printColorLegend() {
	fmt.Println("🎨 COLOR LEGEND:")
	color.New(color.FgMagenta, color.Bold).Print("   🔴 CRITICAL ")
	fmt.Println("- Immediate security risk, requires urgent attention")

	color.New(color.FgRed, color.Bold).Print("   🟠 HIGH     ")
	fmt.Println("- Serious security vulnerability")

	color.New(color.FgYellow).Print("   🟡 MEDIUM   ")
	fmt.Println("- Moderate security issue")

	color.New(color.FgBlue).Print("   🔵 LOW      ")
	fmt.Println("- Minor security concern")
	fmt.Println()
}

// Print help text for additional options
func printHelpText() {
	fmt.Println()
	color.White("💡 ADDITIONAL OPTIONS:")
	fmt.Println("   --verbose, -v      Show detailed descriptions and recommendations")
	fmt.Println("   --compact, -q      Display results in a compact format")
	fmt.Println("   --mask-secrets, -m Mask/redact sensitive values in output")
	fmt.Println("   --format, -f       Output format (console, json)")
	fmt.Println("   --output, -o       Save results to a file")
	fmt.Println()
}

func printBanner() {
	// Get version information
	version := "v1.0"

	if compact {
		// Simplified banner for compact mode
		fmt.Printf("MCP SCAN %s - Security Scanner for MCP Configurations\n", version)
		return
	}

	// Full visual banner
	banner := color.New(color.FgCyan, color.Bold)
	banner.Println("╔═══════════════════════════════════════════════════════════════╗")
	banner.Println("║                         MCP SCAN " + version + "                        ║")
	banner.Println("║              Security Scanner for MCP Configurations         ║")
	banner.Println("║                     SwampUP 2025 Edition                     ║")
	banner.Println("╚═══════════════════════════════════════════════════════════════╝")

	// Add scan date and hostname
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "unknown"
	}

	scanDate := time.Now().Format("January 2, 2006")

	info := color.New(color.FgWhite)
	info.Printf("🔒 Secure MCP Scanner | Date: %s | Host: %s\n", scanDate, hostname)
	info.Println("🔍 Scanning for credentials, command injection, and filesystem risks")
	fmt.Println()
}

func discoverConfigs() ([]string, error) {
	var configs []string

	// Add user-specified paths
	configs = append(configs, configPaths...)

	// If no specific configs provided, search common locations
	var commonPaths []string
	if len(configPaths) == 0 {
		commonPaths = []string{
			"./mcp.json",
			"./config/mcp.json",
			"./.mcp/config.json",
			filepath.Join(os.Getenv("HOME"), ".mcp", "config.json"),
			filepath.Join(os.Getenv("HOME"), ".config", "mcp", "config.json"),
			filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Claude", "claude_desktop_config.json"),
			"./demos", // For demo configs
		}
	}

	for _, path := range commonPaths {
		if path == "./demos" {
			// Scan directory for JSON files
			err := filepath.Walk(path, func(walkPath string, info os.FileInfo, err error) error {
				if err != nil {
					return nil // Ignore errors, directory might not exist
				}
				if strings.HasSuffix(walkPath, ".json") && !info.IsDir() {
					configs = append(configs, walkPath)
				}
				return nil
			})
			if err != nil && verbose {
				color.Yellow("Warning: Could not scan directory %s: %v", path, err)
			}
		} else {
			if stat, err := os.Stat(path); err == nil && !stat.IsDir() {
				configs = append(configs, path)
			}
		}
	}

	// Remove duplicates and filter existing files
	var validConfigs []string
	seen := make(map[string]bool)

	for _, configPath := range configs {
		if seen[configPath] {
			continue
		}
		seen[configPath] = true

		if _, err := os.Stat(configPath); err == nil {
			validConfigs = append(validConfigs, configPath)
		}
	}

	return validConfigs, nil
}

func printConsoleResult(result scanner.ScanResult) {
	if len(result.Issues) == 0 {
		color.Green("✅ %s - No security issues found", result.FilePath)
		return
	}

	// Group issues by severity
	issuesBySeverity := map[string][]scanner.Issue{
		"critical": {},
		"high":     {},
		"medium":   {},
		"low":      {},
	}

	for _, issue := range result.Issues {
		severity := strings.ToLower(issue.Severity)
		issuesBySeverity[severity] = append(issuesBySeverity[severity], issue)
	}

	// Count issues by severity
	criticalCount := len(issuesBySeverity["critical"])
	highCount := len(issuesBySeverity["high"])
	mediumCount := len(issuesBySeverity["medium"])
	lowCount := len(issuesBySeverity["low"])

	// Print file header with severity counts
	fileHeader := color.New(color.FgRed, color.Bold)

	if compact {
		// Compact display format
		fileHeader.Printf("❌ %s: %d issues ", result.FilePath, len(result.Issues))
		fmt.Printf("[Critical: %d, High: %d, Medium: %d, Low: %d]\n",
			criticalCount, highCount, mediumCount, lowCount)

		// Print most critical issues first (top 3 of each severity)
		printCompactIssues("CRITICAL", issuesBySeverity["critical"])
		printCompactIssues("HIGH", issuesBySeverity["high"])

		if verbose {
			printCompactIssues("MEDIUM", issuesBySeverity["medium"])
			printCompactIssues("LOW", issuesBySeverity["low"])
		} else if mediumCount+lowCount > 0 {
			fmt.Printf("   ... and %d more medium/low issues (use --verbose to see all)\n",
				mediumCount+lowCount)
		}

		fmt.Println()
		return
	}

	// Full visual display format
	fileHeader.Printf("❌ %s - %d issues found ", result.FilePath, len(result.Issues))

	if criticalCount > 0 {
		color.New(color.FgMagenta, color.Bold).Printf("(Critical: %d ", criticalCount)
	}
	if highCount > 0 {
		color.New(color.FgRed).Printf("High: %d ", highCount)
	}
	if mediumCount > 0 {
		color.New(color.FgYellow).Printf("Medium: %d ", mediumCount)
	}
	if lowCount > 0 {
		color.New(color.FgBlue).Printf("Low: %d", lowCount)
	}
	fmt.Println(")")

	// Print divider
	color.White("   ─────────────────────────────────────────────────────────")

	// Display issues by severity (highest to lowest)
	severities := []string{"critical", "high", "medium", "low"}
	severitySymbols := map[string]string{
		"critical": "🔴",
		"high":     "🟠",
		"medium":   "🟡",
		"low":      "🔵",
	}

	for _, severity := range severities {
		issues := issuesBySeverity[severity]
		if len(issues) == 0 {
			continue
		}

		severityColor := getSeverityColor(severity)
		severityColor("   %s %s (%d issues)", severitySymbols[severity], strings.ToUpper(severity), len(issues))

		for i, issue := range issues {
			prefix := "   ├── "
			if i == len(issues)-1 {
				prefix = "   └── "
			}

			severityColor("%s%s", prefix, issue.Title)

			if verbose {
				indent := "   │   "
				if i == len(issues)-1 {
					indent = "       "
				}

				color.White("%s%s", indent, issue.Description)
				if issue.Recommendation != "" {
					color.Blue("%s💡 %s", indent, issue.Recommendation)
				}
				if issue.Location != "" {
					color.Cyan("%s🔍 %s", indent, issue.Location)
				}
			}
		}

		fmt.Println()
	}
}

// Helper function to print issues in compact format
func printCompactIssues(severity string, issues []scanner.Issue) {
	if len(issues) == 0 {
		return
	}

	// Get color for this severity
	severityColor := getSeverityColor(strings.ToLower(severity))

	// Print at most 3 issues per severity in compact mode
	maxToShow := len(issues)
	if maxToShow > 3 && !verbose {
		maxToShow = 3
	}

	for i := 0; i < maxToShow; i++ {
		issue := issues[i]
		severityColor("   [%s] %s", severity, issue.Title)
		if verbose && issue.Location != "" {
			color.Cyan("      Location: %s", issue.Location)
		}
	}

	// Show a count of remaining issues if not showing all
	if maxToShow < len(issues) {
		severityColor("   ... and %d more %s issues\n", len(issues)-maxToShow, strings.ToLower(severity))
	}
}

func getSeverityColor(severity string) func(format string, a ...interface{}) {
	switch strings.ToLower(severity) {
	case "critical":
		return func(format string, a ...interface{}) {
			color.New(color.FgMagenta, color.Bold).Printf(format, a...)
		}
	case "high":
		return func(format string, a ...interface{}) {
			color.New(color.FgRed, color.Bold).Printf(format, a...)
		}
	case "medium":
		return func(format string, a ...interface{}) {
			color.New(color.FgYellow).Printf(format, a...)
		}
	case "low":
		return func(format string, a ...interface{}) {
			color.New(color.FgBlue).Printf(format, a...)
		}
	default:
		return func(format string, a ...interface{}) {
			color.White(format, a...)
		}
	}
}

func printSummary(configCount, issueCount int, issuesBySeverity map[string]int) {
	divider := "════════════════════════════════════════════════════════════════"

	if compact {
		// Compact summary display
		fmt.Println(divider)
		color.Cyan("SUMMARY: Scanned %d configs, found %d issues", configCount, issueCount)

		if issueCount > 0 {
			fmt.Printf("Critical: %d, High: %d, Medium: %d, Low: %d\n",
				issuesBySeverity["critical"], issuesBySeverity["high"],
				issuesBySeverity["medium"], issuesBySeverity["low"])

			if issuesBySeverity["critical"] > 0 {
				color.New(color.FgMagenta, color.Bold).Println("CRITICAL SECURITY ISSUES DETECTED - ACTION REQUIRED")
			}
		} else {
			color.Green("No security issues found - All configurations are secure!")
		}
		fmt.Println(divider)
		return
	}

	// Full visual summary
	color.White(divider)

	title := color.New(color.FgCyan, color.Bold)
	title.Println("📊 SCAN SUMMARY")

	color.White("   Configurations scanned: %d", configCount)

	if issueCount == 0 {
		color.Green("   Security issues found: %d ✅", issueCount)
		color.Green("   🎉 All configurations are secure!")
	} else {
		// Show issue breakdown by severity
		color.Red("   Security issues found: %d ❌", issueCount)

		// Draw severity bar chart
		if issuesBySeverity["critical"] > 0 {
			criticalBar := strings.Repeat("█", min(issuesBySeverity["critical"], 20))
			color.New(color.FgMagenta, color.Bold).Printf("   🔴 Critical: %d %s\n", issuesBySeverity["critical"], criticalBar)
		}

		if issuesBySeverity["high"] > 0 {
			highBar := strings.Repeat("█", min(issuesBySeverity["high"], 20))
			color.New(color.FgRed, color.Bold).Printf("   🟠 High:     %d %s\n", issuesBySeverity["high"], highBar)
		}

		if issuesBySeverity["medium"] > 0 {
			mediumBar := strings.Repeat("█", min(issuesBySeverity["medium"], 20))
			color.New(color.FgYellow).Printf("   🟡 Medium:   %d %s\n", issuesBySeverity["medium"], mediumBar)
		}

		if issuesBySeverity["low"] > 0 {
			lowBar := strings.Repeat("█", min(issuesBySeverity["low"], 20))
			color.New(color.FgBlue).Printf("   🔵 Low:      %d %s\n", issuesBySeverity["low"], lowBar)
		}

		// Add action recommendations based on severity
		fmt.Println()
		if issuesBySeverity["critical"] > 0 {
			color.New(color.FgMagenta, color.Bold).Println("   ⚠️  CRITICAL SECURITY ISSUES DETECTED - IMMEDIATE ACTION REQUIRED!")
		} else if issuesBySeverity["high"] > 0 {
			color.New(color.FgRed, color.Bold).Println("   ⚠️  High severity issues detected - prompt action recommended")
		} else {
			color.Yellow("   ⚠️  Please review and address the security findings above")
		}

		// Add timestamp
		now := time.Now().Format("2006-01-02 15:04:05")
		color.White("   Scan completed: %s", now)
	}

	color.White(divider)

	// Add command suggestion for JSON report if there are issues
	if issueCount > 0 {
		fmt.Println()
		color.White("💡 To generate a detailed JSON report:")
		color.Cyan("   ./mcp-scan -f json -o security-report.json")
		fmt.Println()
	}
}

// Helper function to get the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
