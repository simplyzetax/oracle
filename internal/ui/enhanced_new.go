package ui

import (
	"fmt"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

// Enhanced color palette
var (
	// Brand colors
	enhancedBlue   = lipgloss.Color("#00D9FF")
	enhancedPurple = lipgloss.Color("#8B5CF6")
	enhancedGold   = lipgloss.Color("#F59E0B")
	enhancedGreen  = lipgloss.Color("#10B981")
	enhancedOrange = lipgloss.Color("#F97316")

	// Neutral colors
	enhancedCharcoal = lipgloss.Color("#374151")
	enhancedSlate    = lipgloss.Color("#64748B")
	enhancedPearl    = lipgloss.Color("#F8FAFC")
)

// Enhanced styles
var (
	bannerStyle = lipgloss.NewStyle().
			Foreground(enhancedBlue).
			Bold(true).
			Align(lipgloss.Center).
			MarginBottom(1)

	enhancedResponseStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(enhancedSlate).
				Padding(1, 2).
				MarginTop(1).
				MarginBottom(1)
)

// Glamour renderer for markdown
var enhancedMarkdownRenderer *glamour.TermRenderer

func init() {
	var err error
	enhancedMarkdownRenderer, err = glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(80),
	)
	if err != nil {
		enhancedMarkdownRenderer = nil // Fallback to nil if glamour fails
	}
}

// ShowBanner displays the Oracle banner with enhanced styling
func ShowBanner() {
	banner := `
 ██████╗ ██████╗  █████╗  ██████╗██╗     ███████╗
██╔═══██╗██╔══██╗██╔══██╗██╔════╝██║     ██╔════╝
██║   ██║██████╔╝███████║██║     ██║     █████╗  
██║   ██║██╔══██╗██╔══██║██║     ██║     ██╔══╝  
╚██████╔╝██║  ██║██║  ██║╚██████╗███████╗███████╗
 ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚══════╝╚══════╝`

	taglineStyle := lipgloss.NewStyle().
		Foreground(enhancedSlate).
		Italic(true).
		Align(lipgloss.Center).
		MarginBottom(2)

	fmt.Println(bannerStyle.Render(banner))
	fmt.Println(taglineStyle.Render("AI-Powered Terminal Assistant with Command Execution"))
}

// StreamMarkdownText outputs streaming text (placeholder for future enhancement)
func StreamMarkdownText(text string) {
	fmt.Print(text)
}

// RenderFinalResponse renders the complete response with markdown support
func RenderFinalResponse(fullText string) {
	if enhancedMarkdownRenderer == nil {
		// Fallback to basic rendering
		fmt.Println(enhancedResponseStyle.Render(fullText))
		return
	}

	rendered, err := enhancedMarkdownRenderer.Render(fullText)
	if err != nil {
		// Fallback on error
		fmt.Println(enhancedResponseStyle.Render(fullText))
		return
	}

	fmt.Println(enhancedResponseStyle.Render(rendered))
}

// Enhanced ConfirmExecution using huh for better UX
func EnhancedConfirmExecution(command string) bool {
	var confirm bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Execute this command?").
				Description(fmt.Sprintf("Command: %s", command)).
				Affirmative("Yes, execute").
				Negative("No, skip").
				Value(&confirm),
		),
	)

	err := form.Run()
	if err != nil {
		// Fallback to original method
		return ConfirmExecution(command)
	}

	return confirm
}

// Enhanced ConfirmContinueOnError using huh
func EnhancedConfirmContinueOnError() bool {
	var continueExec bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Command failed! Continue with remaining commands?").
				Affirmative("Yes, continue").
				Negative("No, stop").
				Value(&continueExec),
		),
	)

	err := form.Run()
	if err != nil {
		// Fallback to original method
		return ConfirmContinueOnError()
	}

	return continueExec
}

// ShowEnhancedCommandsTable displays commands in a beautiful table
func ShowEnhancedCommandsTable(commands []string) {
	if len(commands) == 0 {
		return
	}

	header := lipgloss.NewStyle().
		Foreground(enhancedGold).
		Bold(true).
		SetString("⚡ Detected Executable Commands:")

	fmt.Println(header.Render())
	fmt.Println()

	// Create table rows
	rows := make([][]string, len(commands))
	for i, cmd := range commands {
		rows[i] = []string{
			fmt.Sprintf("%d", i+1),
			cmd,
		}
	}

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(enhancedGold)).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return lipgloss.NewStyle().
					Foreground(enhancedGold).
					Bold(true).
					Align(lipgloss.Center)
			}

			if col == 0 {
				return lipgloss.NewStyle().
					Foreground(enhancedBlue).
					Bold(true).
					Align(lipgloss.Center).
					Width(5)
			}

			return lipgloss.NewStyle().
				Padding(0, 1)
		}).
		Headers("#", "Command").
		Rows(rows...)

	fmt.Println(t.Render())
	fmt.Println()
}

// ShowEnhancedExecutionStatus shows beautiful execution status
func ShowEnhancedExecutionStatus(message string, statusType string) {
	var style lipgloss.Style
	var icon string

	switch statusType {
	case "success":
		style = lipgloss.NewStyle().Foreground(enhancedGreen).Bold(true)
		icon = "✅"
	case "error":
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6B6B")).Bold(true)
		icon = "❌"
	case "executing":
		style = lipgloss.NewStyle().Foreground(enhancedOrange).Bold(true)
		icon = "⚡"
	case "warning":
		style = lipgloss.NewStyle().Foreground(enhancedGold).Bold(true)
		icon = "⚠️"
	default:
		style = lipgloss.NewStyle().Foreground(enhancedBlue).Bold(true)
		icon = "ℹ️"
	}

	fmt.Println(style.Render(fmt.Sprintf("%s %s", icon, message)))
}
