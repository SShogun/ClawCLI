# Contributing to ClawCLI

First off, thank you for considering contributing to ClawCLI! It's people like you that make ClawCLI such a great tool.

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the issue list as you might find out that you don't need to create one. When you are creating a bug report, please include as many details as possible:

* **Use a clear and descriptive title**
* **Describe the exact steps which reproduce the problem**
* **Provide specific examples to demonstrate those steps**
* **Describe the behavior you observed after following the steps**
* **Explain which behavior you expected to see instead and why**
* **Include your .env file configuration** (without API key!)
* **Include your OS version and Go version**

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, please include:

* **Use a clear and descriptive title**
* **Provide a step-by-step description of the suggested enhancement**
* **Provide specific examples to demonstrate the steps**
* **Describe the current behavior and the expected behavior**
* **Explain why this enhancement would be useful**

### Pull Requests

* Fill in the required template
* Follow the Go code style guidelines
* Include appropriate test cases
* End all files with a newline
* Avoid platform-specific code when possible

## Development Setup

### Prerequisites

- Go 1.21 or later
- Git
- Make (optional but recommended)
- An Anthropic API key

### Setup Steps

1. Fork the repository and clone it locally:
   ```bash
   git clone https://github.com/your-username/ClawCLI.git
   cd ClawCLI
   ```

2. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. Set up your environment:
   ```bash
   cp .env.example .env
   # Edit .env with your Anthropic API key
   ```

4. Install dependencies:
   ```bash
   go mod tidy
   ```

5. Run tests to ensure everything works:
   ```bash
   go test ./...
   ```

### Making Changes

1. **Write your code** following the guidelines below
2. **Run tests** to ensure existing functionality still works:
   ```bash
   go test ./...
   ```
3. **Format your code**:
   ```bash
   go fmt ./...
   ```
4. **Add/update tests** for new features
5. **Update README.md** if needed

### Code Style Guidelines

- Follow the standard Go code style (use `gofmt`)
- Use meaningful variable and function names
- Keep functions small and focused on a single responsibility
- Add comments for exported functions and types
- Write tests for all new functionality
- Keep comments concise and clear

### Git Commit Messages

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests liberally after the first line

Examples:
```
Add ask command for one-shot questions
Fix off-by-one error in token calculation
Refactor config loading to use godotenv
```

### Testing

We expect all new code to include tests. Here are some guidelines:

- Use table-driven tests where appropriate
- Test both success and error cases
- Name test functions as `TestXxx` where `Xxx` is what's being tested
- Use `t.Run()` for sub-tests
- Aim for meaningful test coverage

Run tests:
```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Submitting Changes

1. **Push to your fork**:
   ```bash
   git push origin feature/your-feature-name
   ```

2. **Create a Pull Request** with:
   - Clear description of changes
   - Reference to related issues
   - Screenshots if UI changes
   - Test coverage for new features

3. **Address feedback** from code review
4. **Ensure CI passes** before merge

## Project Structure

```
ClawCLI/
├── cmd/               # CLI commands
├── internal/
│   ├── ai/           # AI client logic
│   ├── config/       # Configuration
│   ├── types/        # Data types
│   ├── ui/           # Terminal UI
│   └── utils/        # Utilities
├── main.go
├── go.mod
└── README.md
```

## Additional Notes

### Code Organization

- Keep related code together
- Use packages to organize functionality
- Avoid circular dependencies
- Make functions testable and reusable

### Error Handling

- Return errors explicitly
- Wrap errors with context using `fmt.Errorf`
- Don't ignore errors unless you have a good reason

### Documentation

- Document exported types and functions
- Add examples to complex functions
- Keep comments up to date with code

## Questions?

Feel free to:
- Open an issue with the question tag
- Create a discussion in GitHub Discussions
- Reach out to maintainers

## Contributor License Agreement

By contributing to ClawCLI, you agree that your contributions will be licensed under its MIT license.

## Recognition

Contributors will be recognized in the README.md file. Thank you for helping make ClawCLI better!
