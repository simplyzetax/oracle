package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Oracle version and feature information",
	Long:  `Display version information and available features for Oracle CLI tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔮 Oracle CLI Tool")
		fmt.Println("Version: 1.0.0")
		fmt.Println("Features:")
		fmt.Println("  ✨ AI-powered question answering with Google Gemini")
		fmt.Println("  🎨 Beautiful terminal UI with styled output")
		fmt.Println("  🔄 Real-time streaming responses")
		fmt.Println("  💭 Interactive question prompts")
		fmt.Println("  ⚡ Command execution capabilities (with --execute flag)")
		fmt.Println("  🛡️  Safe command detection and confirmation")
		fmt.Println()
		fmt.Println("Supported Models:")
		fmt.Println("  • gemini-2.0-flash-exp (default)")
		fmt.Println("  • gemini-pro")
		fmt.Println("  • gemini-pro-vision")
		fmt.Println()
		fmt.Println("Repository: https://github.com/simplyzetax/oracle")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
