package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	apiKey   string
	verbose  bool
	model    string
	provider string
	baseURL  string
)

var rootCmd = &cobra.Command{
	Use:   "clawcli",
	Short: "An AI-Powered CLI Assistant",
	Long: `CLAW CLI is a command-line tool that provides
AI-powered code assistance, explanations, and chat capabilities.

Supports both local Ollama models and Anthropic Claude API.

Usage examples:
  clawcli chat                    Start interactive chat
  clawcli ask "explain pointers"  Ask a one-shot question
  clawcli explain main.go         Explain a file
  clawcli review main.go          Review code in a file
  clawcli version                 Show version info`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.claw-cli.yaml)")
	rootCmd.PersistentFlags().StringVar(&provider, "provider", "", "AI provider: ollama or anthropic (default: ollama)")
	rootCmd.PersistentFlags().StringVar(&baseURL, "base-url", "", "API base URL (default: http://localhost:11434)")
	rootCmd.PersistentFlags().StringVar(&apiKey, "api-key", "", "API key (required for anthropic, not needed for ollama)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	rootCmd.PersistentFlags().StringVar(&model, "model", "", "model to use (e.g. qwen2.5-coder, llama3.1:8b)")

	viper.BindPFlag("provider", rootCmd.PersistentFlags().Lookup("provider"))
	viper.BindPFlag("base-url", rootCmd.PersistentFlags().Lookup("base-url"))
	viper.BindPFlag("api-key", rootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("model", rootCmd.PersistentFlags().Lookup("model"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error finding home directory:", err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".claw-cli")
	}

	viper.SetEnvPrefix("CLAW")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}
}
