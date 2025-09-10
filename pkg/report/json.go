package report

import (
	"mcp-scan/pkg/scanner"
	"time"
)

type JSONReport struct {
	Metadata ReportMetadata     `json:"metadata"`
	Summary  ScanSummary        `json:"summary"`
	Results  []scanner.ScanResult `json:"results"`
}

type ReportMetadata struct {
	Tool      string    `json:"tool"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
	Scanner   string    `json:"scanner"`
}

type ScanSummary struct {
	TotalConfigs      int            `json:"totalConfigs"`
	TotalIssues       int            `json:"totalIssues"`
	IssuesBySeverity  map[string]int `json:"issuesBySeverity"`
	ConfigsWithIssues int            `json:"configsWithIssues"`
}

func GenerateJSONReport(results []scanner.ScanResult) *JSONReport {
	summary := calculateSummary(results)
	
	return &JSONReport{
		Metadata: ReportMetadata{
			Tool:      "mcp-scan",
			Version:   "1.0.0",
			Timestamp: time.Now(),
			Scanner:   "MCP Security Scanner",
		},
		Summary: summary,
		Results: results,
	}
}

func calculateSummary(results []scanner.ScanResult) ScanSummary {
	totalIssues := 0
	configsWithIssues := 0
	severityCounts := map[string]int{
		"critical": 0,
		"high":     0,
		"medium":   0,
		"low":      0,
	}

	for _, result := range results {
		if len(result.Issues) > 0 {
			configsWithIssues++
		}

		for _, issue := range result.Issues {
			totalIssues++
			severityCounts[issue.Severity]++
		}
	}

	return ScanSummary{
		TotalConfigs:      len(results),
		TotalIssues:       totalIssues,
		IssuesBySeverity:  severityCounts,
		ConfigsWithIssues: configsWithIssues,
	}
}