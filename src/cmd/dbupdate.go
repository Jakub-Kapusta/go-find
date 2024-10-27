/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Jakub-Kapusta/go-find/apps/dbupdate"
	"github.com/spf13/cobra"
)

var dbupdateCmd = &cobra.Command{
	Version: "v0.0.1",
	Use:     "dbfind TODO",
	Short:   "Find files and directories and cache them in a database.",
	Long: `A partial GNU findutils replacement implemented ing GO.

This application is under construction.`,
	Run: func(cmd *cobra.Command, args []string) {
		dbupdate.DbUpdate(ctx, args, rootDir, isSearchPath, searchPath)
	},
}

func init() {
	rootCmd.AddCommand(dbupdateCmd)
}
