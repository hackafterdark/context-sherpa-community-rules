# Contributing to Context Sherpa Community Rules

Thank you for your interest in contributing to the Context Sherpa Community Rules! This repository hosts a curated collection of ast-grep rules that help developers identify potential issues, enforce best practices, and improve code quality across multiple programming languages.

## 🌟 Why Contribute?

Your contributions help both human developers and AI coding agents by:
- **Expanding Rule Coverage**: Adding rules for new languages, frameworks, or coding patterns
- **Improving Code Quality**: Creating rules that catch common mistakes and security issues
- **AI Agent Enhancement**: Rules contributed here can be discovered and used by AI agents through the Context Sherpa MCP server
- **Community Knowledge**: Sharing your expertise and coding standards with the broader developer community

## 📋 Table of Contents

- [Getting Started](#getting-started)
- [Understanding the Project](#understanding-the-project)
- [Rule Development Process](#rule-development-process)
- [Rule Format and Structure](#rule-format-and-structure)
- [Testing Requirements](#testing-requirements)
- [Submission Guidelines](#submission-guidelines)
- [Validation and CI/CD](#validation-and-cicd)
- [Best Practices](#best-practices)
- [Community Guidelines](#community-guidelines)
- [Getting Help](#getting-help)

## 🚀 Getting Started

### 1. Explore the Repository

Before creating new rules, familiarize yourself with the existing structure:

```bash
# Clone the repository
git clone https://github.com/your-username/context-sherpa-community-rules.git
cd context-sherpa-community-rules

# Explore existing rules by language and category
ls rules/

# Check out some example rules to understand the patterns
cat rules/go/security/sql-injection.yml
```

### 2. Prerequisites

- **ast-grep CLI**: Install ast-grep to test your rules locally
  ```bash
  # Using npm (recommended)
  npm install -g @ast-grep/cli

  # Or using cargo (if you have Rust installed)
  cargo install ast-grep
  ```
- **Git**: For version control and submitting contributions
- **Text Editor**: Any editor that supports YAML syntax highlighting

### 3. Choose Your Contribution Area

Rules are organized by:
- **Language**: C, C++, C#, CSS, Go, HTML, Java, JavaScript, Kotlin, Lua, Python, Rust, Scala, Swift, Thrift, TSX, TypeScript
- **Category**: security, performance, best-practices, style, error-handling, etc.
- **Complexity**: Start with well-defined patterns before tackling complex multi-pattern rules

## 🎯 Understanding the Project

### Project Goals

This repository serves as a comprehensive rule database for:
- **Static Code Analysis**: Automated detection of code smells, bugs, and security issues
- **AI-Powered Development**: Rules that AI coding agents can discover and apply
- **Community Standards**: Shared best practices across different coding communities
- **Educational Resource**: Examples of how to write effective ast-grep patterns

### How AI Agents Use These Rules

The Context Sherpa MCP server enables AI agents to:
- **Search Community Rules**: Find existing rules for common problems
- **Import Rules**: Add community rules to local projects
- **Create Custom Rules**: Generate new rules from natural language descriptions
- **Validate Code**: Check generated code against established patterns

## 🔧 Rule Development Process

### Step 1: Identify the Problem

Start by identifying a specific code pattern or anti-pattern you want to detect:

**Good Examples:**
- "SQL injection vulnerabilities in database queries"
- "Unused variables in functions"
- "Missing error handling for file operations"
- "Hardcoded secrets or credentials"

**Avoid:**
- Overly broad patterns that match too many cases
- Subjective style preferences without clear rationale
- Platform-specific patterns that don't generalize

### Step 2: Research Existing Rules

Check if similar rules already exist:

```bash
# Search for similar rules by pattern
grep -r "sql-injection" rules/

# Check rules in the same language/category
ls rules/go/security/
```

### Step 3: Design Your Pattern

Create a precise ast-grep pattern that:
- **Matches the specific issue** you want to detect
- **Minimizes false positives** with appropriate constraints
- **Handles edge cases** with `not` conditions where needed
- **Uses proper ast-grep syntax** for your target language

### Step 4: Create Test Cases

Develop realistic test cases that demonstrate:
- **Valid code** that should pass your rule
- **Invalid code** that should trigger your rule
- **Edge cases** that test the boundaries of your pattern

### Step 5: Test Locally

Validate your rule works as expected:

```bash
# Test against your test files
ast-grep scan --rule path/to/your-rule.yml path/to/test/files/

# Debug pattern matching
ast-grep scan --rule path/to/your-rule.yml --debug
```

## 📋 Rule Format and Structure

All rules must follow this exact YAML structure:

```yaml
id: tool-language-category-rule-name
language: target-language
message: "Brief, actionable description of what this rule detects"
severity: error|warning|info
metadata:
  author: github-username
  tags: tag1, tag2, tag3
  description: "Detailed explanation of the rule's purpose and rationale"
rule:
  # ast-grep pattern definition using proper syntax
  pattern: |
    # Your pattern here
```

### Required Fields

| Field | Description | Example |
|-------|-------------|---------|
| `id` | Unique identifier using format: `{tool}-{language}-{category}-{descriptive-name}` | `ast-grep-go-security-sql-injection` |
| `language` | Target programming language | `go`, `python`, `javascript` |
| `message` | Concise description shown when rule matches | `"Potential SQL injection vulnerability detected"` |
| `severity` | Impact level | `error`, `warning`, `info` |
| `rule` | The ast-grep pattern definition | See [ast-grep documentation](https://ast-grep.github.io/reference/yaml.html) |

### Metadata Fields

| Field | Description | Example |
|-------|-------------|---------|
| `author` | Your GitHub username | `john-doe` |
| `tags` | Comma-separated keywords for discoverability | `security, database, injection` |
| `description` | Detailed explanation of the rule's purpose | `"Detects SQL injection vulnerabilities by identifying string concatenation in SQL queries"` |

### Rule ID Naming Convention

Follow this format: `{tool}-{language}-{category}-{descriptive-name}`

**Examples:**
- `ast-grep-go-security-sql-injection`
- `ast-grep-python-best-practices-unused-variable`
- `ast-grep-javascript-error-handling-missing-catch`

## 🧪 Testing Requirements

Every rule **must** include corresponding test files to ensure it works correctly and continues to work as the codebase evolves.

### Test File Structure

```
rule-tests/
└── [language]/
    └── [category]/
        └── [rule-name]-test.yml  # YAML file with valid/invalid test cases
```

### Test Files (YAML Format)

Test cases are embedded directly in YAML files:

```yaml
# rule-tests/go/security/sql-injection-test.yml
id: ast-grep-go-sql-injection
valid:
  - |
    package main
    import "database/sql"
    func goodExample() {
        query := "SELECT * FROM users WHERE id = ?"
        row := db.QueryRow(query, userID)  // ✅ Safe parameterized query
    }
invalid:
  - |
    package main
    import "fmt"
    func badExample() {
        query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", userID)  // ❌ SQL injection
        row := db.QueryRow(query)
    }
```

**Requirements:**
- Show realistic mistakes developers might make
- Include multiple instances of the problematic pattern
- Ensure violations are clear and unambiguous
- Test edge cases and variations

## 📤 Submission Guidelines

### 1. Create Your Rule

Write your YAML rule file following the exact format specified above. Place it in the appropriate directory:

```
rules/[language]/[category]/[your-rule-name].yml
```

### 2. Create Test File

Create a test file with `valid` and `invalid` test cases:

```
rule-tests/[language]/[category]/[your-rule-name]-test.yml
```

**Test file format:**
```yaml
id: your-rule-id
valid:
  - |
    // Code that should NOT trigger the rule
invalid:
  - |
    // Code that SHOULD trigger the rule
```

### 3. Test Locally

Validate your rule works as expected:

```bash
# Test using the ast-grep test framework
ast-grep test --skip-snapshot-tests

# Or test individual rules
ast-grep scan --rule rules/go/security/my-rule.yml --debug
```

### 4. Submit a Pull Request

1. **Fork** the repository on GitHub
2. **Create** a new branch for your contribution:
   ```bash
   git checkout -b add-new-rule-feature
   ```
3. **Commit** your changes:
   ```bash
   git add .
   git commit -m "Add new rule: ast-grep-go-security-my-pattern"
   ```
4. **Push** your branch:
   ```bash
   git push origin add-new-rule-feature
   ```
5. **Create** a Pull Request with:
   - Clear title describing the rule
   - Detailed description of what the rule does
   - Examples of code it catches/fixes
   - Rationale for why this rule is valuable

## ✅ Validation and CI/CD

All pull requests are automatically validated through GitHub Actions:

### Automated Checks

1. **Test Validation**: Ensures test files exist for every rule
2. **Rule Testing**: Runs `ast-grep test --skip-snapshot-tests`:
   - Validates all rules against their embedded test cases
   - Reports pass/fail status with detailed diagnostics
3. **YAML Validation**: Checks rule and test file format and syntax
4. **Auto-indexing**: Updates `index.json` after successful merge

### Manual Review

Maintainers will review your contribution for:
- **Pattern Effectiveness**: Does it catch real issues without too many false positives?
- **Test Quality**: Are the test cases realistic and comprehensive?
- **Documentation**: Is the rule clearly explained and justified?
- **Consistency**: Does it follow the established patterns and conventions?

**Note**: Pull requests cannot be merged until all CI checks pass.

## 🌟 Best Practices

### Pattern Writing

- **Be Specific**: Write focused rules that detect one issue rather than broad patterns
- **Minimize False Positives**: Use `not` conditions and constraints to avoid matching valid code
- **Handle Edge Cases**: Consider different ways the pattern might appear in real code
- **Performance**: Avoid overly complex patterns that might slow down scans

### Rule Design

- **Clear Rule IDs**: Use descriptive, hyphen-separated identifiers
- **Actionable Messages**: Write messages that tell developers how to fix the issue
- **Appropriate Severity**: Use `error` for security/correctness issues, `warning` for best practices, `info` for minor issues
- **Good Tags**: Choose relevant, searchable tags that help others discover your rule

### Testing

- **Realistic Examples**: Use code patterns that actually occur in real projects
- **Multiple Test Cases**: Include various ways the pattern might appear
- **Edge Case Coverage**: Test boundary conditions and unusual cases
- **Comprehensive Coverage**: Ensure both valid and invalid cases are well-represented

### Documentation

- **Clear Descriptions**: Explain not just what the rule does, but why it matters
- **Usage Context**: Describe when and where this rule is most valuable
- **Examples**: Include code examples in your PR description
- **Rationale**: Explain the reasoning behind the pattern and constraints

## 🤝 Community Guidelines

### Code of Conduct

We expect all contributors to:
- **Be Respectful**: Treat other contributors with courtesy and professionalism
- **Be Collaborative**: Work together to improve rules and fix issues
- **Be Constructive**: Focus on improving code quality, not personal preferences
- **Be Inclusive**: Welcome contributions from developers of all backgrounds and experience levels

### Contribution Standards

- **Quality First**: Prioritize well-tested, effective rules over quantity
- **Maintainability**: Write rules that are easy to understand and modify
- **Relevance**: Focus on patterns that provide real value to developers
- **Originality**: Avoid duplicating existing rules unless you're improving them

### Reporting Issues

If you find problems with existing rules:
1. **Check Existing Issues**: Search for similar reports first
2. **Provide Details**: Include specific examples and expected vs actual behavior
3. **Suggest Improvements**: Propose fixes or enhancements when possible
4. **Be Patient**: Complex rule issues may take time to investigate and fix

## 🆘 Getting Help

### Resources

- **ast-grep Documentation**: [Official Guide](https://ast-grep.github.io/)
- **YAML Reference**: [Pattern Syntax](https://ast-grep.github.io/reference/yaml.html)
- **Testing Guide**: [Test Your Rules](https://ast-grep.github.io/guide/test-rule.html)
- **Existing Rules**: Browse `rules/` for examples
- **Test Files**: Study `rule-tests/` for testing patterns

### Getting Support

- **Issues**: Use GitHub Issues for bug reports and feature requests
- **Discussions**: Use GitHub Discussions for questions and community support
- **Pull Requests**: Welcome for bug fixes, improvements, and new rules

### Common Questions

**Q: How do I test my rule locally?**
```bash
# Using the ast-grep test framework
ast-grep test --skip-snapshot-tests

# Or test individual rules
ast-grep scan --rule rules/go/security/my-rule.yml --debug
```

**Q: What languages are supported?**
A: All languages supported by ast-grep: C, C++, C#, CSS, Go, HTML, Java, JavaScript, Kotlin, Lua, Python, Rust, Scala, Swift, Thrift, TSX, TypeScript

**Q: Can I modify existing rules?**
A: Yes! Improvements to existing rules are welcome. Just ensure your changes don't break existing functionality.

**Q: How do I choose the right category for my rule?**
A: Use these general guidelines:
- `security`: Authentication, injection, validation issues
- `performance`: Optimization, memory, efficiency concerns
- `best-practices`: Style, maintainability, design patterns
- `error-handling`: Exceptions, validation, resilience
- `code-quality`: Complexity, duplication, readability

---

Thank you for contributing to the Context Sherpa Community Rules! Your efforts help make code better for everyone. 🎉