# Security Policy

> ## ⚠️ **EXPERIMENTAL EDUCATIONAL SOFTWARE** ⚠️
> 
> **MCP Scan is designed for educational and demonstration purposes ONLY.**
> 
> - **DO NOT use for production security assessments**
> - **DO NOT rely on for compliance or audit requirements**  
> - **Results may be inaccurate (false positives/negatives)**
> - **Always verify findings with security professionals**
> - **Not a replacement for enterprise security tools**

## Supported Versions

We provide security updates for the following versions of MCP Scan:

| Version | Supported          |
| ------- | ------------------ |
| 1.x.x   | ✅ Active support  |
| 0.x.x   | ❌ No longer supported |

## Reporting a Vulnerability

### 🚨 For Critical Security Issues

If you discover a **critical security vulnerability** in MCP Scan itself (not in MCP configurations it scans), please report it privately:

1. **Use GitHub's Private Vulnerability Reporting**: [Report Privately](https://github.com/AndreaGriffiths11/mcp-config-scan/security/advisories/new)
2. **Or email us**: security@[maintainer-domain] (if available)

**Please DO NOT** create a public issue for critical vulnerabilities.

### 📋 What to Include

When reporting security issues, please include:

- **Description**: Clear explanation of the vulnerability
- **Impact**: What could an attacker accomplish?
- **Steps to reproduce**: Minimal example (without real secrets!)
- **Environment**: OS, Go version, MCP Scan version
- **Suggested fix**: If you have ideas for remediation

### 🔍 For Detection Issues (False Positives/Negatives)

For issues with MCP Scan's security detection (missed threats or false alarms):

- **Use public GitHub issues** with the "security" label
- **Redact all real secrets** from example configurations
- These help improve the tool for everyone!

## Response Timeline

- **Critical vulnerabilities**: Response within 24-48 hours
- **Other security issues**: Response within 5 business days
- **Public disclosure**: After fix is available (coordinated disclosure)

## Security Best Practices

### For MCP Scan Users

1. **Keep MCP Scan Updated**: Always use the latest version
2. **Validate Results**: Review findings before taking action
3. **Secure Your Configs**: Use environment variables for secrets
4. **Regular Scanning**: Integrate into your CI/CD pipeline

### For Contributors

1. **No Hardcoded Secrets**: Never commit real API keys or tokens
2. **Secure Development**: Follow secure coding practices
3. **Dependency Updates**: Keep dependencies current and secure
4. **Test Security Features**: Verify detection accuracy

## Known Security Considerations

### False Positives
- MCP Scan may flag legitimate configuration patterns
- Always review findings in context
- Use `--mask-secrets` to safely share scan results

### False Negatives  
- New secret formats may not be detected immediately
- Complex obfuscation techniques might bypass detection
- Regular updates improve detection coverage

### Data Handling
- MCP Scan processes configurations locally
- No data is sent to external services
- Scan results may contain sensitive information

## Security Features

### ✅ What MCP Scan Protects Against

- **Exposed API Keys**: 30+ service types detected
- **Credential Leaks**: AWS, GitHub, OpenAI, and more
- **Path Traversal**: Dangerous filesystem access patterns
- **Command Injection**: Shell metacharacters in arguments  
- **Insecure Settings**: Debug mode, disabled TLS, etc.

### ⚠️ What MCP Scan Cannot Detect

- **Novel Secret Formats**: New or custom API key patterns
- **Encrypted Secrets**: Base64 or other encoded credentials
- **Dynamic Vulnerabilities**: Runtime-specific security issues
- **Network Security**: Firewall rules, network policies
- **Authentication Logic**: MCP server implementation flaws

## Responsible Disclosure

We appreciate security researchers and users who help improve MCP Scan's security:

### Recognition
- Security contributors will be credited in release notes
- Significant findings may be featured in project documentation  
- We maintain a security hall of fame for notable contributions

### Guidelines
- Provide reasonable time for fixes before public disclosure
- Avoid accessing others' data or disrupting services
- Don't perform testing on systems you don't own
- Follow applicable laws and regulations

## Security Updates

### How We Handle Vulnerabilities

1. **Assessment**: Evaluate severity and impact
2. **Development**: Create and test security patches
3. **Release**: Deploy fixes in patch releases
4. **Communication**: Notify users through GitHub releases
5. **Documentation**: Update security advisories

### Staying Informed

- **Watch**: GitHub repository for security announcements
- **Subscribe**: To release notifications
- **Follow**: Project updates and security bulletins

## Contact Information

- **Private Reports**: Use GitHub's private vulnerability reporting
- **Public Discussion**: GitHub issues with "security" label
- **General Questions**: GitHub discussions

## Security Commitment

MCP Scan is committed to:

- **Transparency**: Open communication about security issues
- **Responsiveness**: Timely handling of vulnerability reports  
- **Improvement**: Continuous enhancement of security features
- **Education**: Helping users understand MCP security risks

---

**Last Updated**: January 2025

Thank you for helping keep MCP Scan and the MCP ecosystem secure! 🔒