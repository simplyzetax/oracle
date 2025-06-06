package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

  $ source <(oracle completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ oracle completion bash > /etc/bash_completion.d/oracle
  # macOS:
  $ oracle completion bash > $(brew --prefix)/etc/bash_completion.d/oracle

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ oracle completion zsh > "${fpath[1]}/_oracle"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ oracle completion fish | source

  # To load completions for each session, execute once:
  $ oracle completion fish > ~/.config/fish/completions/oracle.fish

PowerShell:

  PS> oracle completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> oracle completion powershell > oracle.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			_ = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			_ = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			_ = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			_ = cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
