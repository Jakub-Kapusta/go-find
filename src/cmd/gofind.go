// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package cmd

import (
	"context"
	"os"
	"sync"

	"github.com/Jakub-Kapusta/go-find/apps/find"
	"github.com/Jakub-Kapusta/go-find/internal/signals"
	"github.com/spf13/cobra"
)

var (
	// WaitGroup for global goroutines.
	wg sync.WaitGroup
	// Global ctx: should only be cancelled by the signal handler.
	ctx, cf = context.WithCancel(context.Background())
	// Args.
	rootDir     string
	unsafePrint bool
	print0      bool
)

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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		signals.LaunchSignalHandler(ctx, cf, &wg)
	},
	// args[0] is the first actual argument, and not the name of the program.
	// Only arguments not caught by our flag definitions will be present.
	Run: func(cmd *cobra.Command, args []string) {
		find.Find(ctx, args, rootDir, unsafePrint, print0)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		cf()
		wg.Wait()
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
