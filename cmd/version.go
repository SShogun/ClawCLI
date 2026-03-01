package cmd

import (
	"fmt"
	"runtime"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	Version   = "1.0.0"
	BuildDate = "unknown"
	GitCommit = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run:   runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(cmd *cobra.Command, args []string) {
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("12"))

	label := lipgloss.NewStyle().
		Foreground(lipgloss.Color("8"))

	fmt.Println(header.Render("\n🤖 CLAW CLI Assistant"))
	fmt.Println(dividerStyle.Render("────────────────────────────────────────"))
	fmt.Printf("  %s %s\n", label.Render("Version:"), Version)
	fmt.Printf("  %s %s\n", label.Render("Build:  "), BuildDate)
	fmt.Printf("  %s %s\n", label.Render("Commit: "), GitCommit)
	fmt.Printf("  %s %s\n", label.Render("Go:     "), runtime.Version())
	fmt.Printf("  %s %s/%s\n", label.Render("OS/Arch:"), runtime.GOOS, runtime.GOARCH)
	fmt.Println()
}
