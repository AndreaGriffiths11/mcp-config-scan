# Contributing to MCP Scan

Thank you for your interest in contributing to MCP Scan! This guide will help you get started.

## 🚀 Quick Start

1. **Fork** the repository on GitHub
2. **Clone** your fork locally:
   ```bash
   git clone https://github.com/YOUR-USERNAME/mcp-config-scan.git
   cd mcp-config-scan
   ```
3. **Install** dependencies:
   ```bash
   go mod download
   ```
4. **Build** and test:
   ```bash
   go build -o mcp-scan
   ./mcp-scan demo
   ```

## 📋 Development Guidelines

### Code Style
- Follow standard Go conventions (`go fmt`, `go vet`)
- Use meaningful variable and function names
- Add comments for complex logic
- Keep functions focused and small

### Project Structure
```
mcp-scan/
├── cmd/              # CLI commands and flags
├── pkg/
│   ├── config/       # MCP configuration parsing
│   ├── scanner/      # Security detection logic  
│   └── report/       # Output formatting
├── demos/            # Sample configurations
└── main.go           # Entry point
```

### Testing
- Add tests for all new security checks
- Test with both valid and invalid configs
- Include edge cases and false positive prevention
- Run tests: `go test ./...`

## 🔍 Adding New Security Checks

### 1. Secret Pattern Detection
Add new patterns to `pkg/scanner/scanner.go`:

```go
// Add to secretPatterns map
"New Service API Key": regexp.MustCompile(`\bnewsvc_[A-Za-z0-9]{20,}\b`),
```

### 2. Configuration Checks
Add validation logic in `scanInsecureConfigs()`:

```go
if someCondition {
    issues = append(issues, Issue{
        Severity:       "high",
        Title:          "New security issue detected",
        Description:    "Detailed explanation",
        Recommendation: "How to fix it",
        Location:       location + ".setting",
    })
}
```

### 3. Testing New Checks
Create test cases in `demos/` directory:

```json
{
  "mcpServers": {
    "testServer": {
      "env": {
        "NEW_API_KEY": "newsvc_abc123def456ghi789jkl"
      }
    }
  }
}
```

## 🐛 Bug Reports

Good bug reports include:
- **Description**: What happened vs what you expected
- **Environment**: OS, Go version, MCP Scan version
- **Reproduction**: Minimal config that triggers the issue
- **Output**: Complete error message or unexpected behavior

**⚠️ Important**: Redact real secrets before sharing configs!

## 💡 Feature Requests

When suggesting features:
- **Use Case**: Why is this needed?
- **Examples**: Show what the feature would look like
- **Implementation**: Ideas on how it could work
- **Alternatives**: Other solutions you've considered

## 📝 Documentation

Help improve documentation by:
- Fixing typos and unclear explanations
- Adding usage examples
- Updating command-line help text
- Creating tutorials or guides

## 🔄 Pull Request Process

### Before Submitting
1. **Test thoroughly** with demo configs
2. **Run all tests**: `go test ./...`
3. **Check formatting**: `go fmt ./...`
4. **Verify builds**: `go build -o mcp-scan`
5. **Update docs** if needed

### PR Description Template
```markdown
## Summary
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Documentation update
- [ ] Performance improvement

## Testing
- [ ] Added/updated tests
- [ ] Tested with demo configs
- [ ] Verified no regressions

## Checklist
- [ ] Code follows project style
- [ ] Self-reviewed changes
- [ ] Updated documentation
```

### Review Process
1. Maintainers will review within 3-5 days
2. Address feedback promptly
3. Keep PR focused on single feature/fix
4. Be responsive to suggestions

## 🏷️ Issue Labels

| Label | Description |
|-------|-------------|
| `bug` | Something isn't working |
| `enhancement` | New feature or request |
| `good first issue` | Good for newcomers |
| `help wanted` | Extra attention needed |
| `security` | Security-related issue |
| `documentation` | Documentation improvements |

## ⚡ Priority Areas

We especially need help with:

### 🔐 **New Secret Patterns**
- Cloud providers (Mistral, Cohere, etc.)
- Database connection strings
- Container registry tokens
- CI/CD platform keys

### 🛡️ **Security Checks**
- Network security misconfigurations
- Container security issues
- File permission problems
- Injection attack vectors

### 🧪 **Testing & Quality**
- Unit tests for edge cases
- Integration tests
- Performance benchmarks
- Cross-platform compatibility

### 📦 **DevOps & Automation**
- GitHub Actions workflows
- Release automation
- Package managers (Homebrew, etc.)
- Docker containers

## 💬 Getting Help

- **GitHub Issues**: For bugs and feature requests
- **GitHub Discussions**: For questions and general discussion
- **Code Review**: Tag maintainers for urgent reviews

## 📜 Code of Conduct

### Our Pledge
We are committed to making participation in our project a harassment-free experience for everyone, regardless of experience level, gender, gender identity and expression, sexual orientation, disability, personal appearance, body size, race, ethnicity, age, religion, or nationality.

### Standards
- Be respectful and inclusive
- Focus on constructive feedback
- Accept responsibility for mistakes
- Learn from the community

### Enforcement
Unacceptable behavior can be reported to project maintainers. All complaints will be reviewed and investigated promptly and fairly.

---

Thank you for contributing to MCP Scan! Your efforts help make MCP configurations more secure for everyone. 🙏