# MCP Scan - Scanner for MCP Configurations

🔒 **Experimental security scanner for Model Context Protocol (MCP) configuration files**

> ## ⚠️ **IMPORTANT: EXPERIMENTAL SOFTWARE** ⚠️
> 
> **THIS IS EDUCATIONAL SOFTWARE FOR DEMONSTRATION PURPOSES ONLY**
> 
> - **Not for production security decisions**
> - **May produce false positives or miss real vulnerabilities** 
> - **For learning about MCP security concepts only**
> - **Always verify findings manually**
> - **Not a replacement for professional security auditing**

---

## Features

- Detects exposed API keys, secrets, and cloud credentials
- Finds dangerous filesystem and command injection risks
- Checks for insecure configuration settings
- Colorful console output & structured JSON reports
- Sensitive data masking, report encryption, secure deletion

---

## 📊 SwampUp 2025 Talk  
**Slides and demo materials:**  
➡️ [SwampUp 2025 Slides & Demo](swampup/README.md)

---

## Installation

### Option 1: Download Pre-built Binary
```bash
# Linux/macOS
curl -L https://github.com/AndreaGriffiths11/mcp-config-scan/releases/latest/download/mcp-scan -o mcp-scan
chmod +x mcp-scan

# Windows
curl -L https://github.com/AndreaGriffiths11/mcp-config-scan/releases/latest/download/mcp-scan.exe -o mcp-scan.exe
```

### Option 2: Build from Source
```bash
git clone https://github.com/AndreaGriffiths11/mcp-config-scan.git
cd mcp-config-scan
go build -o mcp-scan
```

### Option 3: Go Install
```bash
go install github.com/AndreaGriffiths11/mcp-config-scan@latest
```

### Usage Examples

```bash
./mcp-scan                            # Scan current directory
./mcp-scan -c config/mcp.json         # Scan specific files
./mcp-scan -f json -o report.json     # Generate JSON report
./mcp-scan demo                       # Run demo scan
./mcp-scan -v                         # Verbose output
./mcp-scan -q                         # Compact output format
```

---

## Configuration Discovery

Scans config files in:
- `./mcp.json`
- `./config/mcp.json`
- `~/.mcp/config.json`
- `~/.config/mcp/config.json`
- Any JSON files in `./demos/`

---

## Security Checks

- **Critical:** Exposed API keys, cloud credentials, private keys, DB credentials
- **High:** Dangerous filesystem access, shell injection, insecure network settings
- **Medium:** Dangerous commands, debug mode, suspicious paths
- **Low:** Disabled configs, excessive timeouts

---

## Example Output

```
╔═══════════════════════════════════════════════════════════════╗
║                         MCP SCAN v1.0                         ║
║              Security Scanner for MCP Configurations          ║
║                     SwampUP 2025 Edition                      ║
╚═══════════════════════════════════════════════════════════════╝

✅ demos/secure-config.json - No security issues found
❌ demos/vulnerable-config.json - 12 issues found:
   [CRITICAL] Exposed OpenAI API Key detected
   [HIGH] Dangerous filesystem access in workingDir
   [HIGH] Potential shell injection vector
   [MEDIUM] Debug mode enabled

════════════════════════════════════════════════════════════════
📊 SCAN SUMMARY
   Configurations scanned: 4
   Security issues found: 18 ❌
⚠️  Please review and address the security findings above
════════════════════════════════════════════════════════════════
```

---

## Demo Configurations

Educational examples in the `demos/` directory:
- **secure-config.json** — Best practices
- **vulnerable-config.json** — Intentionally unsafe
- **mixed-config.json** — Secure & risky configs
- **development-config.json** — Common pitfalls

---

## Building from Source

```bash
go mod download
go build -o mcp-scan
```

---

## Contributing

We welcome contributions! Here's how you can help:

### 🐛 **Report Bugs**
- Use [GitHub Issues](https://github.com/AndreaGriffiths11/mcp-config-scan/issues)
- Include MCP config samples (redact secrets!)
- Describe expected vs actual behavior

### 💡 **Suggest Features**
- New secret patterns to detect
- Additional security checks
- Output format improvements
- Integration ideas

### 🔧 **Code Contributions**
1. **Fork** the repository
2. **Create** feature branch: `git checkout -b feature/amazing-detection`
3. **Add** tests for new security checks
4. **Ensure** all tests pass: `go test ./...`
5. **Submit** pull request with clear description

### 📋 **Development Setup**
```bash
git clone https://github.com/AndreaGriffiths11/mcp-config-scan.git
cd mcp-config-scan
go mod download
go build -o mcp-scan
./mcp-scan demo  # Test it works
```

### 🎯 **Areas Needing Help**
- Additional API key patterns (Anthropic, Mistral, etc.)
- Container/Docker security checks
- Performance optimizations
- Windows compatibility testing

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

---

## License

MIT License — see LICENSE file.

---

## ⚠️ Security Disclaimer

> ## **CRITICAL: EDUCATIONAL USE ONLY**
> 
> **MCP Scan is an experimental educational tool and should NEVER be used for:**
> - Production security decisions
> - Compliance reporting  
> - Security assessments of live systems
> - Any situation where accuracy is critical
> 
> **This tool may:**
> - Generate false positives (flag safe configurations as dangerous)
> - Miss real security vulnerabilities (false negatives)
> - Misinterpret configuration contexts
> - Fail to detect novel attack patterns
> 
> **Always:**
> - Verify all findings manually
> - Consult security professionals for production systems
> - Use established enterprise security tools for real audits
> - Understand this is for learning MCP security concepts only 
