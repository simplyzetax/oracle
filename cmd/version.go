package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version will be set during build time
	Version = "dev"
	// Commit will be set during build time  
	Commit = "none"
	// Date will be set during build time
	Date = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Oracle version and feature information",
	Long:  `Display version information and available features for Oracle CLI tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Oracle CLI Tool")
		fmt.Printf("Version: %s\n", Version)
		fmt.Printf("Commit: %s\n", Commit)
		fmt.Printf("Built: %s\n", Date)
		fmt.Println("Features:")
		fmt.Println("  - AI-powered question answering with Google Gemini")
		fmt.Println("  - Terminal UI with styled output")
		fmt.Println("  - Real-time streaming responses")
		fmt.Println("  - Interactive question prompts")
		fmt.Println("  - Command execution (with --execute flag)")
		fmt.Println("  - Safe command detection and confirmation")
		fmt.Println()
		fmt.Println("Supported Models:")
		fmt.Println("  - gemini-2.0-flash-exp (default)")
		fmt.Println("  - gemini-pro")
		fmt.Println("  - gemini-pro-vision")
		fmt.Println()
		fmt.Println("Repository: https://github.com/simplyzetax/oracle")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
