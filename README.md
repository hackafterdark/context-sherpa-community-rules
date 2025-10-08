# Context Sherpa Community Rules

A community-driven repository of ast-grep rules for code analysis and linting. This repository hosts a curated collection of rules that help developers identify potential issues, enforce best practices, and improve code quality across multiple programming languages.

**🤖 Designed for AI Integration**: These rules are specifically crafted to work seamlessly with [Context Sherpa](https://github.com/hackafterdark/context-sherpa) - an AI-powered MCP server that enables coding agents to discover, validate, and dynamically create linting rules based on natural language feedback.

## Features

- **Public & Transparent**: All rules are hosted in a public GitHub repository for easy browsing and verification
- **No Authentication Required**: Access rules anonymously without API keys or tokens
- **Community-Driven**: Contributions welcome through standard GitHub pull requests
- **Multi-Language Support**: Organized by linter tool and programming language for easy discovery
- **Automated Quality Assurance**: CI/CD pipeline validates all contributions before acceptance
- **AI-Powered Integration**: Rules work seamlessly with AI coding agents through the Context Sherpa MCP server

## Built With

This project is powered by [**ast-grep**](https://ast-grep.github.io/) - a lightning-fast tool for searching and linting code using AST patterns. We're incredibly grateful to the ast-grep team for creating such a powerful and flexible platform that makes this community rules repository possible!

## Repository Structure

```
context-sherpa-community-rules/
├── rules/                        # ast-grep rules organized by language/category
│   └── go/security/
│       └── sql-injection.yml     # Rule definition with pattern matching
├── rule-tests/                   # Test files for rule validation
│   └── go/security/
│       └── sql-injection-test.yml # Test cases (valid/invalid examples)
├── index.json                    # Auto-generated catalog of all rules
├── sgconfig.yml                  # ast-grep configuration
└── README.md
```

## Contributing

We welcome contributions from the community! Help us expand our collection of ast-grep rules to improve code quality across multiple programming languages.

### Quick Start

1. **Read the detailed guide**: See [CONTRIBUTING.md](./CONTRIBUTING.md) for comprehensive contribution guidelines
2. **Explore existing rules**: Browse `ast-grep/rules/` to understand current patterns and identify gaps
3. **Choose your focus**: Pick a language and category where you can add value
4. **Submit a pull request**: All contributions are validated through automated CI/CD

### Key Requirements

- **Rule Format**: YAML files following the `tool-language-category-rule-name` naming convention
- **Testing**: Every rule must include `valid` and `invalid` test cases
- **Validation**: All PRs validated through GitHub Actions before merging
- **Languages**: Support for 16+ languages including Go, Python, JavaScript, Rust, and more

For detailed instructions on rule format, testing requirements, submission process, and best practices, please see our comprehensive [Contributing Guide](./CONTRIBUTING.md).