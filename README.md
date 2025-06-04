# Oracle CLI Tool

Oracle is a beautiful CLI tool that allows you to ask questions to AI models (currently Gemini) with streaming responses and elegant UI styling.

## Features

- 🤖 **AI-Powered**: Chat with Google's Gemini models
- 🎨 **Beautiful UI**: Styled with Charm's Lipgloss for elegant terminal output
- 🔄 **Streaming**: Real-time response streaming
- 💭 **Interactive**: Prompt for questions if none provided
- ⚡ **Command Execution**: Oracle can detect and run shell commands (with --execute flag)
- 🛡️ **Safe Execution**: Command detection with user confirmation and safety checks
- ⚙️ **Configurable**: Multiple models and API key options

## Installation

1. Clone the repository
2. Set your Google AI API key:
   ```bash
   export GOOGLE_AI_API_KEY="your-api-key-here"
   ```
3. Build the application:
   ```bash
   go build -o oracle
   ```

## Usage

### Ask a question directly:
```bash
oracle ask "What is the meaning of life?"
oracle ask "Explain quantum computing in simple terms"
```

### Interactive mode (prompts for question):
```bash
oracle ask
```

### With command execution enabled:
```bash
oracle ask "How do I list all files in the current directory?" --execute
oracle ask "Show me the Git status and stage all changes" -x
```

### With custom model:
```bash
oracle ask "Write a haiku about coding" --model gemini-pro
```

### With API key flag:
```bash
oracle ask "Hello world" --api-key your-key-here
```

### Command execution examples:
```bash
# Oracle will detect commands and ask for confirmation
oracle ask "How do I create a new Git repository?" --execute

# Oracle might suggest and offer to run:
# $ git init
# $ git add .
# $ git commit -m "Initial commit"
```

## Project Structure

```
oracle/
├── main.go              # Application entry point
├── cmd/                 # Command definitions
│   ├── root.go         # Root command and global flags
│   ├── ask.go          # Ask command implementation
│   └── version.go      # Version command
├── internal/
│   ├── ai/             # AI client and interaction logic
│   │   └── client.go   # Gemini API client
│   ├── commands/       # Command execution system
│   │   └── executor.go # Command detection and execution
│   └── ui/             # User interface and styling
│       ├── display.go  # Output styling and display
│       └── input.go    # User input handling
└── pkg/
    └── types/          # Shared types and structures
        └── types.go    # Type definitions
```

## Available Models

- `gemini-2.0-flash-exp` (default)
- `gemini-pro`
- `gemini-pro-vision`

## Requirements

- Go 1.24+
- Google AI API key
- `gum` command-line tool (for interactive prompts)

## Dependencies

- `github.com/spf13/cobra` - CLI framework
- `google.golang.org/genai` - Google Generative AI SDK
- `github.com/charmbracelet/lipgloss` - Terminal styling
- `github.com/charmbracelet/gum` - Interactive CLI components
