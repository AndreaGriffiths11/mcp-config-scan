# SwampUp 2025 — AI Security Demo & Slides

Welcome to the SwampUp 2025 presentation materials for:  
## **When AI Agents Go Rogue: Securing the New Attack Surface**

## About the Talk

**The problem:** 84% of developers now use AI agents, but only 29% trust their accuracy. We're adopting tools faster than we can secure them.

**The reality:** AI agents aren't just writing code anymore. Through the Model Context Protocol (MCP), they're managing your repositories, deploying to production, and accessing sensitive infrastructure with your credentials. Hundreds of MCP servers launched on GitHub this year alone—each one a potential attack vector that traditional security tools can't see.

**What's different:** These aren't your typical vulnerabilities. AI agents can be compromised through their reasoning process itself:
- **Prompt injection attacks** disguised as GitHub issues that trick agents into leaking secrets
- **Tool poisoning** where legitimate MCP servers get hijacked to execute malicious commands  
- **Intent breaking** that turns your helpful code reviewer into a data exfiltrator
- **Token passthrough attacks** where one compromised agent becomes a trojan horse for your entire system

**The solution:** This talk demonstrates how GitHub became the first platform to integrate secret scanning directly into MCP tool calls, stopping prompt injection attacks before secrets leave your control. You'll see real attacks blocked in real-time, learn how to audit your existing MCP installations, and discover how end-to-end provenance from GitHub Actions to JFrog Evidence creates cryptographic proof that your AI agents aren't going rogue.

**What you'll leave with:** Specific commands to audit your MCP setup tonight, a framework for securing AI agents without killing productivity, and actionable strategies that work whether you're defending against sophisticated attacks or just trying to prove to compliance that your binaries came from trusted source code.

This isn't about fearing AI—it's about making AI agents both powerful and provable in your environment.

---

## Contents

- [Slides](#slides)
- [Demo Instructions](#demo-instructions)
- [Demo Configurations](#demo-configurations)
- [References & Further Reading](#references--further-reading)


---

## 🖥️ Slides

[View SwampUp 2025 Slides](AIAgentsGoRogue.pdf)

---

## 🛡️ Demo Instructions

To run the live demo of MCP Scan:

```bash
# Build the scanner
go build -o mcp-scan ../

# Run a demo scan on intentionally vulnerable configs
./mcp-scan demo

# Scan the provided configs in this folder
./mcp-scan -c vulnerable-config.json -c secure-config.json
```

- The `demos/` folder contains realistic configurations for both best-practices and intentionally unsafe setups.

---

## ⚙️ Demo Configurations

- **secure-config.json:** Strong security practices
- **vulnerable-config.json:** Multiple security flaws
- **mixed-config.json:** Secure and risky patterns
- **development-config.json:** Common pitfalls in dev environments

---

## 🔗 References & Further Reading


- [GitHub Secret Scanning](https://docs.github.com/en/code-security/secret-scanning)
- [Copilot Secret Scanning (AI-powered detection)](https://docs.github.com/en/code-security/secret-scanning/copilot-secret-scanning)
- [Model Context Protocol (MCP) Overview](https://github.com/AndreaGriffiths11/mcp-scan)
- [SwampUp 2025 Conference](https://swampup.jfrog.com/)
Got it — here’s the full slide list with all URLs included:

---

### MCP Security & Adoption Resources

* [MCP: What it is and Why it Matters](https://addyo.substack.com/p/mcp-what-it-is-and-why-it-matters) – High-level overview of MCP
* [Why MCP Won – Latent.Space](https://www.latent.space/p/why-mcp-won) – Industry analysis on adoption and impact
* [The MCP GitHub Vulnerability – A Deep Dive into Agentic Threats](https://andreagriffiths11.github.io/mcp-vulnerability-deep-dive) – In-depth forensics of threats
* [GitHub MCP Server Guide (GitHub Blog)](https://github.blog/engineering/github-mcp-server-guide) – Practical reference and usage guidance
* [Safeguarding VS Code (GitHub Blog)](https://github.blog/security/safeguarding-vscode-against-prompt-injections) – Defensive patterns and prompt injection examples
* [MCP Community Registry](https://modelcontextprotocol.org/registry) – Central hub for trusted server integrations
* [MCP Horror Stories](https://securitylab.github.com/research/mcp-horror-stories) – Real-world exploits with threat walkthroughs
* [CVE-2025-6514: Remote Command Injection (RCE)](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2025-6514) – Canonical post-build RCE case
* [MCP Kanban Server CVE](https://nvd.nist.gov/vuln/detail/CVE-2025-6514-kanban) – Tool poisoning command injection example
* [MCP Vulnerabilities Every Developer Should Know (Composio)](https://composio.dev/blog/mcp-vulnerabilities) – Summary of classes of risk
* [What is MCP? (GitHub Blog)](https://github.blog/engineering/what-is-mcp-and-why-it-matters) – Plain-English walkthrough of MCP’s origin
* [OWASP FinBot Agentic AI CTF](https://owasp.org/www-project-finbot/) – Overview and access to live CTF

### Security Practices & Conference Content

* [Session: Securing MCP in an Agentic World (Cisco)](https://cisco.com/events/securing-mcp-agentic-world) – Governance & defense talk
* [MCP Spec: Authorization & Security Best Practices](https://modelcontextprotocol.org/spec) – Official protocol guidance
* [MCP & Its Critical Vulnerabilities (Strobes)](https://strobes.co/research/mcp-critical-vulnerabilities) – Breakdown of attack surfaces and CVEs
* [Dagger/container-use](https://dagger.io/docs/container-use) – Secure multi-agent dev environments
* [OWASP ASI (Feb 2025)](https://owasp.org/www-project-agentic-security-initiative/) – Threats & mitigations for autonomous AI
* [Prioritizing Dependabot Alerts Using Production Context](https://docs.github.com/en/code-security/securing-your-organization/understanding-your-organizations-exposure-to-vulnerabilities/prioritizing-dependabot-alerts-using-production-context) – GitHub guide for contextual vulnerability prioritization
* [DevSecOps in the AI Era: JFrog Powers Agentic Remediation](https://www.businesswire.com/news/home/20250909861213/en/DevSecOps-in-the-AI-Era-JFrog-Powers-Agentic-Remediation-with-Self-Healing-Software-Supply-Chain) – JFrog's approach to self-healing software supply chain
### Vulnerability & Risk Catalogs

* [MCP Security Risks & Best Practices (WorkOS)](https://workos.com/blog/mcp-security-risks) – Enterprise-focused guide
* [The Vulnerable MCP Project](https://vulnerablemcp.org) – Ongoing vulnerability catalog
* [MCP Security Vulnerabilities: Weekend List](https://console.dev/mcp-vulnerabilities-list) – Concise updates on threats & mitigations
* [MCP Security: Best Practices & Pitfalls](https://securityweek.com/mcp-best-practices-pitfalls) – Beginner-friendly audit checklist
* [Datadog Guide: MCP Security Risks](https://www.datadoghq.com/blog/mcp-security-risks) – Operational risks and monitoring tips
* [Understanding the Security Landscape (Apideck)](https://apideck.com/blog/mcp-security-landscape) – End-to-end threat modeling

### Broader Context

* [REvil – Darknet Diaries (Ep. 126)](https://darknetdiaries.com/episode/126) – Ransomware, supply chain & extortion tactics

---

**Questions?**  
Open an issue on [mcp-scan](https://github.com/AndreaGriffiths11/mcp-scan) or contact the author via GitHub.

---
