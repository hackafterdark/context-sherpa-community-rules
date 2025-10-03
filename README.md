# Context Sherpa Community Rules

A community-driven repository of ast-grep rules for code analysis and linting. This repository hosts a curated collection of rules that help developers identify potential issues, enforce best practices, and improve code quality across multiple programming languages.

## Features

- **Public & Transparent**: All rules are hosted in a public GitHub repository for easy browsing and verification
- **No Authentication Required**: Access rules anonymously without API keys or tokens
- **Community-Driven**: Contributions welcome through standard GitHub pull requests
- **Multi-Language Support**: Organized by linter tool and programming language for easy discovery
- **Automated Quality Assurance**: CI/CD pipeline validates all contributions before acceptance

## Repository Structure

```
context-sherpa-community-rules/
├── ast-grep/
│   ├── rules/                    # ast-grep rules organized by language/tool
│   │   └── go/security/
│   │       └── sql-injection.yml
│   └── tests/                   # Test files for rule validation
│       └── go/security/
│           └── sql-injection/
│               ├── valid.go      # Code that should pass the rule
│               └── invalid.go    # Code that should trigger the rule
├── index.json                   # Auto-generated catalog of all rules
└── README.md
```

## Contributing

We welcome contributions from the community! This section explains how to add new rules or improve existing ones.

### Getting Started

1. **Explore Existing Rules**: Browse the `ast-grep/rules/` directory to understand the current rule structure and identify gaps where new rules would be valuable.

2. **Choose a Language**: Rules are organized by programming language under `ast-grep/rules/`. Create a new subdirectory if your language isn't represented yet and is supported by the tool. See [supported languages](https://ast-grep.github.io/reference/yaml.html#language) for valid options (C, C++, C#, CSS, Go, HTML, Java, JavaScript, Kotlin, Lua, Python, Rust, Scala, Swift, Thrift, TSX, TypeScript).

3. **Fork and Clone**: Fork this repository and clone it locally to start working on your contribution.

### Rule Format

Each rule must be a YAML file with the following structure:

```yaml
id: unique-rule-identifier
language: target-language
message: "Brief description of what this rule detects"
severity: error|warning|info
metadata:
  author: your-github-username
  tags: tag1, tag2, tag3
  description: "Detailed explanation of the rule's purpose and rationale"
rule:
  # ast-grep pattern definition
  pattern: |
    # Your pattern here
```

**Required Fields:**
- `id`: Unique identifier (format: `tool-language-category-rule-name`)
- `language`: Target programming language
- `message`: Concise description shown when rule matches
- `severity`: Impact level (error/warning/info)
- `rule`: The ast-grep pattern definition

**Metadata Fields:**
- `author`: Your GitHub username (for contribution tracking)
- `tags`: Comma-separated keywords for discoverability (e.g., "security, performance, style")
- `description`: Detailed explanation of the rule's purpose

For complete documentation on all available fields and advanced configuration options, see the [official ast-grep YAML reference](https://ast-grep.github.io/reference/yaml.html).

### Testing Requirements

Every rule **must** include corresponding test files to ensure it works correctly:

- **`tests/[language]/[category]/[rule-name]/valid.ext`**: Code that should **pass** the rule (no violations)
- **`tests/[language]/[category]/[rule-name]/invalid.ext`**: Code that should **trigger** the rule (has violations)

Test files must use the same file extension as your target language (.go, .py, .js, etc.).

### Submission Process

1. **Create Your Rule**: Write your YAML rule file following the format above
2. **Add Test Files**: Create `valid` and `invalid` test cases for your rule
3. **Test Locally**: Run `ast-grep scan` on your test files to verify the rule works as expected:
   ```bash
   ast-grep scan --rule path/to/your-rule.yml path/to/test/files/
   ```
4. **Submit a Pull Request**: Create a PR with your new rule and test files

### Validation Process

All pull requests are automatically validated through GitHub Actions:

1. **Test Validation**: CI checks that test files exist for every rule
2. **Rule Testing**: Runs ast-grep against test files:
   - Must find violations in `invalid` files
   - Must find no violations in `valid` files
3. **Auto-indexing**: After merge, the `index.json` is automatically updated

**Note**: Pull requests cannot be merged until all CI checks pass.

### Best Practices

- **Clear Rule IDs**: Use descriptive, hyphen-separated identifiers (e.g., `go-security-sql-injection`)
- **Specific Patterns**: Write focused rules that detect one specific issue rather than broad patterns
- **Good Test Cases**: Include realistic code examples in your tests that reflect actual usage
- **Documentation**: Provide clear descriptions and metadata to help users understand when to use your rule
- **Validation**: Test your rules thoroughly before submitting

### Example Contribution

Here's how you might contribute a new rule for detecting unused variables in Python:

**Rule file** (`ast-grep/rules/python/best-practices/unused-variable.yml`):
```yaml
id: python-best-practices-unused-variable
language: python
author: your-username
message: "Variable is defined but never used"
severity: warning
metadata:
  tags: best-practices, unused-code, maintenance
  description: "Detects variables that are assigned but never referenced, helping reduce code complexity"
rule:
  pattern: |
    def $FUNC( $$$ ):
      $VAR = $VAL
      return $RET
```

**Test files**:
- `tests/python/best-practices/unused-variable/valid.py` (uses the variable)
- `tests/python/best-practices/unused-variable/invalid.py` (defines but doesn't use the variable)

For more examples, see the existing rules in the `ast-grep/rules/` directory.