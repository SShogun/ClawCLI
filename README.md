# ClawCLI}}

🤖 An AI-Powered Command-Line Assistant built with Go and Claude AI

ClawCLI brings the power of Claude AI directly to your terminal. Ask questions, review code, get explanations, and have interactive conversations with the AI—all from your CLI.

## Features

✨ **Interactive Chat** - Have multi-turn conversations with Claude AI  
❓ **Ask Questions** - Get quick answers to one-shot queries  
📄 **Explain Code** - Understand what your code does  
🔍 **Code Review** - Get AI-powered code reviews with suggestions  
⚙️ **Configurable** - Choose your Claude model, temperature, and token limits  
🎨 **Beautiful Output** - Color-coded terminal output with great styling  

## Prerequisites

- Go 1.21 or later
- An [Anthropic API key](https://console.anthropic.com/account/keys)
- Basic understanding of CLI tools

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/SShogun/ClawCLI.git
cd ClawCLI
```

### 2. Set up your environment

Copy the example environment file and add your API key:

```bash
cp .env.example .env
```

Edit `.env` and add your Anthropic API key:

```dotenv
CLAW_API_KEY=sk-ant-your_actual_key_here
CLAW_MODEL=claude-haiku-4-5-20251001
CLAW_MAX_TOKENS=4096
CLAW_TEMPERATURE=0.7
```

Get your API key from: https://console.anthropic.com/account/keys

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Build the project

```bash
go build -o clawcli.exe
```

Or use the Makefile:

```bash
make build
```

## Usage

### Interactive Chat

Start a multi-turn conversation with the AI:

```bash
./clawcli.exe chat
```

Commands in chat mode:
- `exit` or `quit` - Exit the chat
- `clear` - Clear conversation history  
- `history` - Show conversation history

### Ask a Question

Get a one-shot answer to a question:

```bash
./clawcli.exe ask "What is a goroutine in Go?"
./clawcli.exe ask "How to reverse a string in Go?"
```

### Explain Code

Have the AI explain what a file does:

```bash
./clawcli.exe explain main.go
./clawcli.exe explain src/handler.go
```

### Review Code

Get an AI code review with suggestions:

```bash
./clawcli.exe review main.go
./clawcli.exe review src/service.go
```

The review will cover:
- Summary of what the code does
- Potential bugs or issues
- Performance improvements
- Code style and best practices
- Security concerns

### Version Info

Check installed version and build info:

```bash
./clawcli.exe version
```

## Configuration

Edit your `.env` file to customize:

| Variable | Default | Description |
|----------|---------|-------------|
| `CLAW_API_KEY` | (required) | Your Anthropic API key |
| `CLAW_MODEL` | claude-haiku-4-5-20251001 | Claude model to use |
| `CLAW_MAX_TOKENS` | 4096 | Max response tokens (1-100000) |
| `CLAW_TEMPERATURE` | 0.7 | Response randomness (0-1) |

### Available Models

- **claude-haiku-4-5-20251001** (latest, cheapest, fastest)
- **claude-3-5-haiku-20241022** (budget-friendly)
- **claude-3-5-sonnet-20241022** (balanced performance)
- **claude-3-opus-20250219** (most capable)

## Development

### Project Structure

```
ClawCLI/
├── cmd/               # CLI commands
│   ├── ask.go        # Ask command
│   ├── chat.go       # Chat command
│   ├── explain.go    # Explain command
│   ├── review.go     # Review command
│   ├── root.go       # Root command
│   └── version.go    # Version command
├── internal/
│   ├── ai/           # AI client logic
│   │   ├── client.go     # Claude API client
│   │   └── history.go    # Chat history management
│   ├── config/       # Configuration
│   │   ├── config.go     # Config loader
│   │   └── defaults.go   # Default values
│   ├── types/        # Data types
│   │   └── types.go      # Shared types
│   ├── ui/           # Terminal UI
│   │   └── styles.go     # Terminal styles
│   └── utils/        # Utilities
│       ├── file.go       # File operations
│       └── helpers.go    # Helper functions
├── main.go           # Entry point
├── go.mod            # Go module file
├── .env.example      # Example environment file
├── .gitignore        # Git ignore rules
├── README.md         # This file
└── Makefile          # Build automation (optional)
```

### Building from Source

```bash
# Install dependencies
go mod tidy

# Build binary
go build -o clawcli

# Run tests
go test ./...

# Run with verbose output
go build -v -o clawcli
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run with coverage
go test -cover ./...
```

## Error Handling

### API Key Not Set

If you get an error that your API key is not set:
1. Make sure the `.env` file exists in your project directory
2. Verify `CLAW_API_KEY=` is set to your actual Anthropic API key
3. Check that there are no extra spaces or quotes around the key

### Credit Balance Error

If you see: "Your credit balance is too low..."
- Add credits to your Anthropic account at https://console.anthropic.com/account/billing/overview

### File Not Found

If you get "file not found" when using explain/review:
- Use the full or relative path to the file
- Example: `./clawcli explain ./main.go`

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

MIT License - See [LICENSE](LICENSE) file for details

## Support

- 📖 Check the [examples](examples/) directory for usage examples
- 🐛 Report bugs on GitHub Issues
- 💡 Suggest features on GitHub Discussions
- 📧 Contact: [your-contact-info]

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) CLI framework
- Styling with [Lipgloss](https://github.com/charmbracelet/lipgloss)
- Configuration with [Viper](https://github.com/spf13/viper)
- Powered by [Anthropic Claude AI](https://www.anthropic.com)


