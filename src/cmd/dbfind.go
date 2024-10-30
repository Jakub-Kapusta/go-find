// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package cmd

import (
	"github.com/Jakub-Kapusta/go-find/apps/dbfind"
	"github.com/spf13/cobra"
)

var dbfindCmd = &cobra.Command{
	Version: version,
	Use:     "dbfind TODO",
	Short:   "Find files and directories in the databse cache.",
	Long: `A partial GNU findutils replacement implemented ing GO.

This application is under construction.`,
	Run: func(cmd *cobra.Command, args []string) {
		dbfind.DbFind(ctx, args, unsafePrint, print0, isSearchPath, searchPath)
	},
}

func init() {
	rootCmd.AddCommand(dbfindCmd)
}
