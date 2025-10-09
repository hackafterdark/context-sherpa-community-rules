# AI Agent Guidelines for Context Sherpa Community Rules

This document provides specific guidelines for AI agents contributing to the ast-grep rule repository. These rules help ensure consistent, high-quality rule creation and testing practices.

## Rule Creation Guidelines

### 1. Rule Structure Requirements

**All rules must follow this exact YAML structure:**

```yaml
id: language-category-rule-name
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
- `id` must follow format: `{language}-{category}-{descriptive-name}`
- `author` field MUST be in `metadata` section, not at root level
- `message` should be concise but specific enough to understand the issue
- `tags` should be lowercase, comma-separated, relevant keywords

### 2. Language Support Validation

**Before creating rules, verify language support:**

Valid languages for ast-grep: `C`, `Cpp`, `CSharp`, `Css`, `Go`, `Html`, `Java`, `JavaScript`, `Kotlin`, `Lua`, `Python`, `Rust`, `Scala`, `Swift`, `Thrift`, `Tsx`, `TypeScript`

**Rule:** If an unsupported language is requested, inform the user and suggest alternatives.

### 3. Pattern Writing Best Practices

#### Atomic Rules: The Building Blocks
Atomic rules match single AST nodes. Use them for precision.

- **`pattern`**: Match code structure. For ambiguous syntax (e.g., a JS class field `a = 1`), use a `pattern` object with `context` and `selector`.
  ```yaml
  # Match a class field, not an assignment
  rule:
    pattern:
      selector: field_definition
      context: class A { $FIELD = $INIT }
  ```
- **`kind`**: Match a node by its type name (e.g., `if_statement`, `type_identifier`). This is essential when a pattern is too broad or hard to write.
- **`regex`**: Match a node's text content. **Always combine with `kind` or `pattern`** to avoid performance issues and target the correct node.

#### Composite Rules: Combining Logic
Apply Boolean logic to a **single target node**.

- **`all`**: The node must satisfy ALL sub-rules (AND logic).
- **`any`**: The node must satisfy AT LEAST ONE sub-rule (OR logic).
- **`not`**: The node must NOT satisfy the sub-rule.

#### Relational Rules: Understanding Context
Filter nodes based on their relationship to surrounding nodes. The structure is always `target_node relation surrounding_node`.

- **`inside`**: The target node is a descendant of the surrounding node.
- **`has`**: The target node has a descendant that is the surrounding node.
- **`precedes`**: The target node appears before a sibling surrounding node.
- **`follows`**: The target node appears after a sibling surrounding node.

**Fine-Tuning Relational Rules:**
- **`stopBy: end`**: This is a critical option. By default, relational rules only search immediate neighbors. Use `stopBy: end` to make the search traverse the entire AST (e.g., up to the root for `inside`, down to the leaves for `has`). This is often required for rules that need to understand the broader context of a file.
- **`field`**: Restricts the search to a specific named child field (e.g., `key`, `body`). Only for `inside` and `has`.

---

### 4. Common Patterns & Solutions

#### **Problem: Applying a `regex` to a captured metavariable (`$NAME`)**
**Solution:** Use a composite `all` rule. The first part captures the node with a `pattern`, and the second part applies a `regex` to the captured metavariable.

```yaml
# ✅ Correct: Find type names ending in DTO, DAO, etc.
rule:
  all:
    - any: # First, find all type declarations and capture the name in $NAME
        - pattern: type $NAME struct { $$$ }
        - pattern: type $NAME interface { $$$ }
    - pattern: $NAME # Then, apply the regex to the captured $NAME node
      regex: ".+(DTO|DAO|Impl|Utils|Helper)$"
```

#### **Problem: Excluding matches within a specific context (e.g., inside test functions)**
**Solution:** Use an `all` rule that combines a `pattern` with a `not` > `inside` clause. The `stopBy: end` is crucial here to ensure the entire file context is checked.

```yaml
# ✅ Correct: Find `panic` calls that are NOT inside main, init, or test functions
rule:
  all:
    - pattern: panic($$$) # Find all panic calls
    - not: # Exclude the ones that are...
        inside: # ...inside any of these function patterns
          any:
            - pattern: func main() { $$$ }
            - pattern: func init() { $$$ }
            - pattern: func TestMain(m *testing.M) { $$$ }
            - pattern: func $F(t *testing.T) { $$$ }
            - pattern: func $F(b *testing.B) { $$$ }
            - pattern: func $F(f *testing.F) { $$$ }
          stopBy: end # IMPORTANT: Search all the way up to the file root
```

#### **Problem: Targeting only specific kinds of identifiers (e.g., type names vs. variable names)**
**Solution:** Combine a `regex` with a `kind` rule. Use `any` if multiple kinds are valid targets.

```yaml
# ✅ Correct: Find identifiers with specific suffixes, but only if they are type names
rule:
  all:
    - regex: ".+(DTO|DAO|Impl|Utils|Helper)$"
    - any:
        - kind: type_identifier
        - kind: identifier # Can be included for broader matching if needed
```

## Test File Requirements

### 4. Test Structure Mandate

**Every rule MUST have a corresponding YAML test file.** This project uses `ast-grep`'s inline testing feature, where valid and invalid code snippets are embedded directly within the test YAML.

The test file must be located at: `rule-tests/[language]/[category]/[rule-name]-test.yml`

### 5. Test Content Guidelines

The YAML test file must have the following structure:

```yaml
id: [rule-id] # Must match the id in the rule file
language: [target-language]
rule: [path/to/the/rule.yml]
valid:
  - |
    # Code snippet that should NOT trigger the rule.
    # Include multiple valid cases.
    # This code should be syntactically correct.
  - |
    # Another valid code snippet.
invalid:
  - |
    # Code snippet that SHOULD trigger the rule.
    # Demonstrate the exact anti-pattern.
  - |
    # Another invalid code snippet to show variations.
```

#### `valid` blocks (Should Pass)
- Include realistic code examples that follow best practices.
- Show the correct way to handle the situation the rule detects.
- Include multiple variations when relevant.

#### `invalid` blocks (Should Trigger Rule)
- Demonstrate the exact anti-pattern the rule detects.
- Include multiple instances of the problematic pattern if the rule should catch them all.
- Show realistic mistakes developers might make.

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

**Security Rules Requirements:**
For security-related rules, additional metadata fields are required:
- **OWASP and CWE classifications**: Include relevant OWASP Top 10 categories and CWE identifiers when applicable
- **Security tags**: Always include "security" and "vulnerability" tags for security rules
- **Example security rule metadata** (replace placeholders with appropriate classifications for your specific rule):
  ```yaml
  metadata:
    author: "contributor-username"
    tags: "security,vulnerability,injection"
    description: "Detects potential security vulnerabilities"
    owasp:
      - "AXX:YYYY - Category Name"  # Replace with relevant OWASP classification
    cwe:
      - "CWE-ID: Description"       # Replace with relevant CWE classification
  ```

**Tag Categories:**
- **Security**: authentication, injection, validation, etc.
- **Performance**: optimization, memory, efficiency
- **Best Practices**: style, maintainability, design
- **Error Handling**: exceptions, validation, resilience
- **Code Quality**: complexity, duplication, readability

### 8.1 Testing with Context-Sherpa MCP Server

**When the context-sherpa MCP server is available, AI agents should use it for enhanced rule testing:**

The `context-sherpa` MCP server provides the `scan_path` tool which is ideal for testing rules against their test directories:

```yaml
# Example usage for testing a specific rule's test cases
scan_path:
  path: "ast-grep/tests/go/security/rule-name/"
```

**Benefits of MCP Server Testing:**
- **Integrated Testing**: Test rules directly against their designated test directories
- **Multiple File Support**: Efficiently scan entire test directories containing both valid and invalid examples
- **Consistent Environment**: Ensures rules work correctly in the project's testing environment
- **Early Validation**: Catch issues before manual testing or CLI-based validation

**When to Use MCP Server Testing:**
- After creating or modifying rules
- When validating rule effectiveness against test cases
- Before finalizing rule contributions
- For regression testing when updating existing rules

## Contribution Workflow

### 9. Step-by-Step Process for AI Agents

1. **Analyze Request**: Understand what anti-pattern or issue to detect
2. **Research Existing Rules**: Check if similar rules already exist
3. **Design Pattern**: Create precise ast-grep pattern
4. **Write Rule YAML**: Follow exact structure requirements
5. **Create Test Cases**: Both valid and invalid examples
6. **Validate Locally**: Test with ast-grep CLI if possible
7. **Test with MCP Server**: If context-sherpa MCP server is available, use `scan_path` tool to validate rule against test cases
8. **Document Clearly**: Explain what the rule does and why it matters

### 10. Common Pitfalls to Avoid

**Pattern Issues:**
- ❌ Overly broad patterns that match too much.
- ❌ Patterns that don't account for edge cases.
- ❌ Missing proper escaping in string patterns.
- ❌ Using `kind` and `pattern` as separate rules to filter a pattern. Use a `pattern` object with `context` and `selector` instead.
- ❌ Mixing meta-variables with prefixes/suffixes (e.g., `use$HOOK`). Use `constraints` with `regex` instead.

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

**For MCP server testing (when context-sherpa is available):**

```yaml
# Example: Test a rule using context-sherpa MCP server
use_mcp_tool:
  server_name: "context-sherpa"
  tool_name: "scan_path"
  arguments:
    path: "ast-grep/tests/go/security/rule-name/"
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

## Advanced Topics & Debugging

### Meta-Variable Naming
- Start meta-variables with a `$` sign, followed by uppercase letters (A-Z), underscores (`_`), or digits (1-9).
- A meta-variable must represent a single AST node.
- By default, meta-variables match named AST nodes. Use `$$` (e.g., `$$UNNAMED`) to match unnamed nodes.

### Rule Matching Order
- Rule matching is sequential. The order in an `all` block matters, as meta-variables captured in earlier rules constrain later ones.
- An unordered rule object (not in an `all` block) has an implementation-defined order. For explicit ordering, always use `all`.

### Debugging Rules
- **Use the Playground**: Test patterns and rules in the [ast-grep playground](https://ast-grep.github.io/playground.html).
- **Simplify**: Reduce the rule to the minimal reproducible case.
- **Check CLI vs. Playground**: Discrepancies can arise from different parser versions or text encodings. Use `ast-grep run -p <PATTERN> --debug-query ast` to inspect the AST as seen by the CLI.