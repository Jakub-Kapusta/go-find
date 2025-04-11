// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package cmd

import "os"

var (
	// Directory to base the search in.
	rootDir string
	// Query string for thew full path.
	isSearchPath bool
	searchPath   string
	unsafePrint  bool
	print0       bool
)

func validateFlags() error {
	if searchPath == "" {
		isSearchPath = false
	} else {
		isSearchPath = true
	}
	return nil
}

func setFlags() {
	// Global Flags
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-find.yaml)")
	rootCmd.PersistentFlags().StringVarP(&rootDir, "root_dir", "r", ".", "directory to search in")
	rootCmd.PersistentFlags().StringVar(&searchPath, "path", "", "Query string for the full path. If an empty path is set, all paths will match.")
	rootCmd.PersistentFlags().BoolVarP(&unsafePrint, "unsafe_print", "u", false, "output control characters to the terminal without checks")
	rootCmd.PersistentFlags().BoolVarP(&print0, "print0", "0", false, "print the full file name on the standard output, followed by a null character (instead of the default newline character)")

	if err := validateFlags(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
