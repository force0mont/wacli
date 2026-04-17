// Package cmd provides the CLI commands for wacli.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Version is set at build time via ldflags.
	Version = "dev"
	// Commit is the git commit hash set at build time.
	Commit = "none"
	// Date is the build date set at build time.
	Date = "unknown"
)

// rootCmd is the base command for the CLI.
var rootCmd = &cobra.Command{
	Use:   "wacli",
	Short: "wacli — WhatsApp CLI client",
	Long: `wacli is a command-line interface for interacting with WhatsApp.
It allows you to send messages, manage contacts, and more from your terminal.`,
	SilenceUsage:  true,
	SilenceErrors: true, // handle errors ourselves for cleaner output
}

// versionCmd prints version information.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("wacli %s (commit: %s, built: %s)\n", Version, Commit, Date)
	},
}

// Execute runs the root command.
func Execute(version, commit, date string) {
	Version = version
	Commit = commit
	Date = date

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
