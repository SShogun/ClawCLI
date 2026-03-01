# ClawCLI Usage Examples

A comprehensive guide to using ClawCLI with real-world examples.

## Table of Contents

1. [Installation & Setup](#installation--setup)
2. [Basic Usage](#basic-usage)
3. [Ask Command Examples](#ask-command-examples)
4. [Explain Command Examples](#explain-command-examples)
5. [Review Command Examples](#review-command-examples)
6. [Chat Command Examples](#chat-command-examples)
7. [Advanced Tips](#advanced-tips)

## Installation & Setup

### 1. Build the Project

```bash
cd ClawCLI
go build -o clawcli.exe
```

### 2. Create Your .env File

```bash
cp .env.example .env
# Edit .env with your Anthropic API key
```

### 3. Verify Installation

```bash
./clawcli.exe version
```

## Basic Usage

### Help Command

```bash
./clawcli.exe help
./clawcli.exe ask --help
./clawcli.exe explain --help
./clawcli.exe review --help
```

## Ask Command Examples

### Example 1: Language Question

```bash
./clawcli.exe ask "What is the difference between var, let, and const in JavaScript?"
```

### Example 2: Code Pattern Question

```bash
./clawcli.exe ask "How do I implement the observer pattern in Go?"
```

### Example 3: Troubleshooting

```bash
./clawcli.exe ask "Why would a Go goroutine leak memory?"
```

### Example 4: Best Practices

```bash
./clawcli.exe ask "What are the best practices for error handling in Go?"
```

### Example 5: Multi-word Query

```bash
./clawcli.exe ask "How do you implement OAuth 2.0 in a REST API?"
```

## Explain Command Examples

### Example 1: Explain Our Sample File

```bash
./clawcli.exe explain examples/example.go
```

This will explain:
- What the Calculator struct does
- How the mutex is used
- The purpose of each method
- How the concurrent operations work

### Example 2: Explain Your Own File

```bash
./clawcli.exe explain src/handlers/user.go
./clawcli.exe explain ./config.go
./clawcli.exe explain ../../utils/validation.js
```

## Review Command Examples

### Example 1: Review Our Sample File

```bash
./clawcli.exe review examples/example.go
```

Expected feedback:
- Code structure and organization
- Error handling approach
- Concurrency patterns
- Performance considerations
- Potential improvements

### Example 2: Review Your Code

```bash
./clawcli.exe review main.go
./clawcli.exe review src/service.go
```

## Chat Command Examples

### Example 1: Learning About Go Concurrency

```bash
./clawcli.exe chat

> What are the different ways to do concurrency in Go?
> Can you give me an example with channels?
> What about using goroutines with sync.WaitGroup?
> How would you choose between these approaches?
> exit
```

### Example 2: Debugging Help

```bash
./clawcli.exe chat

> I have a deadlock in my code with two goroutines accessing shared maps
> The deadlock happens intermittently
> Can you help me debug it?
> Here's my code pattern...
> What's the best way to fix this?
> exit
```

### Example 3: Code Design Discussion

```bash
./clawcli.exe chat

> I'm building a REST API. Should I use dependency injection?
> How would I structure the code?
> Can you give me an example of a handler with DI?
> exit
```

### Example 4: Explaining a Concept

```bash
./clawcli.exe chat

> Explain the concept of channels in Go
> How do you use them in practice?
> What about buffered channels? When would I use those?
> Can you give me three real-world examples?
> exit
```

### Special Chat Commands

```
In chat mode:

clear    - Clear conversation history
history  - Show conversation history
exit     - Exit the chat
quit     - Exit the chat
q        - Exit the chat (single letter)
```

## Advanced Tips

### Tip 1: Asking for Specific Output Format

```bash
./clawcli.exe ask "Give me a JSON example of a user object with the following fields: id, name, email, age"
```

### Tip 2: Asking for Code Examples

```bash
./clawcli.exe ask "Show me 3 different ways to implement a singleton pattern in Go"
```

### Tip 3: Performance Questions

```bash
./clawcli.exe ask "How would you optimize this code for better performance?"
```

### Tip 4: Security Questions

```bash
./clawcli.exe ask "What are the security implications of using eval() in JavaScript?"
```

### Tip 5: Learning Path

Use chat for a structured learning experience:

```bash
./clawcli.exe chat

> I want to learn Go. Where do I start?
> What should I learn first?
> Can you give me a practical project I can build?
> What libraries would be most useful for a web server?
> exit
```

## Workflow Examples

### Workflow 1: Understanding Unfamiliar Code

1. First, use `explain` to understand the code:
   ```bash
   ./clawcli.exe explain unknown_file.go
   ```

2. Then use `ask` for specific questions:
   ```bash
   ./clawcli.exe ask "What does the mutex in this code do?"
   ```

3. Finally, use `chat` for a deeper discussion:
   ```bash
   ./clawcli.exe chat
   > I found this code uses reflection. <paste code>
   > Why would they use reflection here?
   > What's the performance impact?
   ```

### Workflow 2: Code Review & Improvement

1. Get an AI review:
   ```bash
   ./clawcli.exe review my_code.go
   ```

2. Ask follow-up questions:
   ```bash
   ./clawcli.exe ask "How would I refactor this to use dependency injection?"
   ```

3. Have a chat session about improvements:
   ```bash
   ./clawcli.exe chat
   > Based on the review I got, should I refactor my handlers?
   > What's the best pattern for large Go projects?
   ```

### Workflow 3: Learning New Concepts

1. Ask introductory questions:
   ```bash
   ./clawcli.exe ask "What is the decorator pattern?"
   ```

2. Get a code explanation:
   ```bash
   ./clawcli.exe explain where_decorator_is_used.go
   ```

3. Have an interactive learning session:
   ```bash
   ./clawcli.exe chat
   > Tell me more about when to use decorators
   > How is it different from middleware?
   > Give me a practical example
   ```

## Configuration Tips

### Using Different Models

Edit your `.env` to try different models:

```dotenv
# Fast and cheap (default)
CLAW_MODEL=claude-haiku-4-5-20251001

# Balanced
CLAW_MODEL=claude-3-5-sonnet-20241022

# Most capable (slower, more expensive)
CLAW_MODEL=claude-3-opus-20250219
```

### Adjusting Response Creativity

```dotenv
# More deterministic (good for code)
CLAW_TEMPERATURE=0.3

# Balanced (default)
CLAW_TEMPERATURE=0.7

# More creative (good for brainstorming)
CLAW_TEMPERATURE=0.9
```

### Token Limits

```dotenv
# Shorter responses
CLAW_MAX_TOKENS=1024

# Default
CLAW_MAX_TOKENS=4096

# Longer, more detailed responses
CLAW_MAX_TOKENS=8192
```

## Troubleshooting

### Issue: API Key Not Found

Make sure:
1. `.env` file exists in the project directory
2. `CLAW_API_KEY=` is set correctly (no quotes)
3. No extra spaces around the key

### Issue: File Not Found

Use the correct path:
```bash
./clawcli.exe explain ./main.go          # Current directory
./clawcli.exe explain src/handler.go     # Relative path
./clawcli.exe explain /absolute/path.go  # Absolute path
```

### Issue: Insufficient Credits

Your Anthropic account needs available credits. Add credits at:
https://console.anthropic.com/account/billing/overview

## Summary

- **`ask`** - Quick, one-shot answers
- **`explain`** - Understand code files
- **`review`** - Get code improvement suggestions
- **`chat`** - Interactive, multi-turn conversations
- **`version`** - Check version info

Happy learning! 🚀
