# AI Agent Guidelines for Context Sherpa Community Rules

This document provides specific guidelines for AI agents contributing to the ast-grep rule repository. These rules help ensure consistent, high-quality rule creation and testing practices.

## Rule Creation Guidelines

### 1. Rule Structure Requirements

**All rules must follow this exact YAML structure:**

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
```

**Critical Points:**
- `id` must follow format: `{tool}-{language}-{category}-{descriptive-name}`
- `author` field MUST be in `metadata` section, not at root level
- `message` should be concise but specific enough to understand the issue
- `tags` should be lowercase, comma-separated, relevant keywords

### 2. Language Support Validation

**Before creating rules, verify language support:**

Valid languages for ast-grep: `C`, `Cpp`, `CSharp`, `Css`, `Go`, `Html`, `Java`, `JavaScript`, `Kotlin`, `Lua`, `Python`, `Rust`, `Scala`, `Swift`, `Thrift`, `Tsx`, `TypeScript`

**Rule:** If an unsupported language is requested, inform the user and suggest alternatives.

### 3. Pattern Writing Best Practices

#### Use Specific Patterns
```yaml
# ✅ Good: Specific and targeted
rule:
  pattern: $DB.Exec(..., fmt.Sprintf($$$), ...)

# ❌ Avoid: Too broad or generic
rule:
  pattern: fmt.Sprintf($$$)
```

#### Leverage ast-grep Features
- Use `any` for multiple patterns that should trigger the rule
- Use `all` for compound conditions
- Use `not` to exclude false positives
- Consider `kind` constraints for more precise matching

#### Common Pattern Examples
```yaml
# Multiple patterns (OR logic)
rule:
  any:
    - pattern: $DB.Query(..., fmt.Sprintf($$$), ...)
    - pattern: $DB.Exec(..., fmt.Sprintf($$$), ...)

# Compound conditions (AND logic)
rule:
  all:
    - pattern: $VAR = $VAL
    - not:
        pattern: $VAR := $VAL

# Error detection pattern
rule:
  pattern: |
    $ERR := $FUNC( $$$ )
  not:
    follows:
      pattern: |
        if $ERR != nil {
```

## Test File Requirements

### 4. Test Structure Mandate

**Every rule MUST have corresponding test files:**

```
tests/
└── [language]/
    └── [category]/
        └── [rule-name]/
            ├── valid.ext      # Code that should NOT trigger rule
            └── invalid.ext    # Code that SHOULD trigger rule
```

### 5. Test Content Guidelines

#### Valid Files (Should Pass)
- Include realistic code examples that follow best practices
- Show the correct way to handle the situation the rule detects
- Include multiple variations when relevant
- Ensure code is syntactically correct

#### Invalid Files (Should Trigger Rule)
- Demonstrate the exact anti-pattern the rule detects
- Include multiple instances of the problematic pattern
- Show realistic mistakes developers might make
- Ensure violations are clear and unambiguous

#### Example Test Cases
```go
// Valid test case - good practices
func goodExample() {
    query := "SELECT * FROM users WHERE id = ?"
    row := db.QueryRow(query, userID)  // ✅ Safe parameterized query
}

// Invalid test case - violations
func badExample() {
    query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", userID)  // ❌ SQL injection
    row := db.QueryRow(query)
}
```

## Quality Assurance Rules

### 6. Pre-Submission Validation

**Before suggesting rule creation, AI agents must:**

1. **Check existing rules** in the target category to avoid duplicates
2. **Validate YAML syntax** of proposed rules
3. **Test patterns locally** if possible using ast-grep CLI
4. **Ensure test files exist** for every rule
5. **Verify file extensions** match the target language

### 7. Rule Effectiveness Metrics

**Good rules should:**
- Have **high precision**: Few false positives
- Have **high recall**: Catch most instances of the problem
- Be **performant**: Not slow down scans significantly
- Be **maintainable**: Easy to understand and modify

### 8. Metadata Completeness

**All rules must include:**

```yaml
metadata:
  author: "contributor-username"  # Required for attribution
  tags: "relevant,categories"     # Required for discoverability
  description: "Clear explanation" # Required for understanding
```

**Tag Categories:**
- **Security**: authentication, injection, validation, etc.
- **Performance**: optimization, memory, efficiency
- **Best Practices**: style, maintainability, design
- **Error Handling**: exceptions, validation, resilience
- **Code Quality**: complexity, duplication, readability

## Contribution Workflow

### 9. Step-by-Step Process for AI Agents

1. **Analyze Request**: Understand what anti-pattern or issue to detect
2. **Research Existing Rules**: Check if similar rules already exist
3. **Design Pattern**: Create precise ast-grep pattern
4. **Write Rule YAML**: Follow exact structure requirements
5. **Create Test Cases**: Both valid and invalid examples
6. **Validate Locally**: Test with ast-grep CLI if possible
7. **Document Clearly**: Explain what the rule does and why it matters

### 10. Common Pitfalls to Avoid

**Pattern Issues:**
- ❌ Overly broad patterns that match too much
- ❌ Patterns that don't account for edge cases
- ❌ Missing proper escaping in string patterns

**Test Issues:**
- ❌ Test files that don't actually trigger/avoid the rule
- ❌ Invalid syntax in test files
- ❌ Missing file extensions or incorrect structure

**Metadata Issues:**
- ❌ Missing required fields in metadata
- ❌ Incorrect field placement (author at root level)
- ❌ Inconsistent or irrelevant tags

## Validation Commands

**For local testing, AI agents should be familiar with:**

```bash
# Test a specific rule against test files
ast-grep scan --rule path/to/rule.yml path/to/test/files/

# Validate YAML syntax
ast-grep scan --rule path/to/rule.yml --debug

# Check for existing similar rules
ast-grep scan --rule path/to/rule.yml path/to/existing/files/
```

## Success Criteria

**A well-written rule contribution should:**
- ✅ Follow exact YAML structure and field requirements
- ✅ Include comprehensive test coverage
- ✅ Have clear, descriptive metadata
- ✅ Pass automated CI validation
- ✅ Address a real development concern
- ✅ Be maintainable and understandable

These guidelines ensure AI agents can effectively contribute high-quality, validated rules that help developers write better code while maintaining consistency with the repository's standards and automation requirements.