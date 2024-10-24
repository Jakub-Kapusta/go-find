// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package cmd

import (
	"os"

	"github.com/Jakub-Kapusta/go-find/apps/find"
	"github.com/spf13/cobra"
)

var rootDir string
var unsafePrint bool
var print0 bool

var rootCmd = &cobra.Command{
	Version: "v0.0.1",
	Use:     "gofind TODO",
	Short:   "Find files and directories.",
	Long: `A partial GNU findutils replacement implemented ing GO.

This application is under construction.`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Run the custom validation logic
		//if myapp.IsValidColor(args[0]) {
		//	return nil
		//}
		return nil
	},
	// args[0] is the first actual argument, and not the name of the program.
	// Only arguments not caught by our flag definitions will be present.
	Run: func(cmd *cobra.Command, args []string) {
		find.Test(args, rootDir, unsafePrint, print0)
	},
}

func init() {
	// Global Flags
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-find.yaml)")
	rootCmd.PersistentFlags().StringVarP(&rootDir, "root_dir", "r", "./", "directory to search in")
	rootCmd.PersistentFlags().BoolVarP(&unsafePrint, "unsafe_print", "u", false, "output control characters to the terminal without checks")
	rootCmd.PersistentFlags().BoolVarP(&print0, "print0", "0", false, "print the full file name on the standard output, followed by a null character (instead of the default newline character)")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
