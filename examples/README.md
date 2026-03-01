# ClawCLI Examples

This directory contains practical examples of using ClawCLI for various tasks.

## Contents

- `example.go` - Sample Go file for explaining/reviewing
- `USAGE_EXAMPLES.md` - Detailed usage examples

## Quick Examples

### Ask a Question

```bash
clawcli ask "What is a pointer in Go?"
clawcli ask "How do I sort a slice in Go?"
clawcli ask "Explain the defer keyword"
```

### Code Explanation

```bash
clawcli explain example.go
```

### Code Review

```bash
clawcli review example.go
```

### Interactive Chat

```bash
clawcli chat

# Then in the chat:
> How do I create a goroutine?
> What's the difference between channels and mutexes?
> exit
```

## Most Common Use Cases

### 1. Quick Question

Perfect for quick, one-shot questions:

```bash
clawcli ask "What does GOMAXPROCS do?"
```

### 2. Understanding Code

When you need to understand what a piece of code does:

```bash
clawcli explain main.go
```

### 3. Code Review

Get suggestions on improving your code:

```bash
clawcli review handler.go
```

### 4. Interactive Learning

Have a conversation to learn about a topic:

```bash
clawcli chat
> Tell me about Go concurrency patterns
> Can you give me an example with channels?
> How would that compare to using sync.WaitGroup?
```

## Tips

- Use `ask` for quick questions
- Use `explain` to understand existing code
- Use `review` to improve your code
- Use `chat` for exploratory learning and multi-turn conversations

## Configuration

All examples assume you have:
1. Set up your `.env` file with your Anthropic API key
2. Built the binary: `go build -o clawcli.exe`
3. Added the binary location to your PATH (optional)
