package cmd

import (
	"github.com/spf13/cobra"
)

func SetVersionInfo(version string) { _ = "STUB: not implemented"; return }

// NewVersionCmd return a versionCmd instance.
func NewVersionCmd() *cobra.Command {
	_ = "STUB: not implemented"
	// versionCmd represents the version command.
	return nil
}

func init() {
	rootCmd.AddCommand(NewVersionCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
