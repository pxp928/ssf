/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package ssfrepo

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Command instantiates the pipelinerun command
func Command() *cobra.Command {
	c := &cobra.Command{
		Use:   "repo",
		Short: "Manage Repos",
		Annotations: map[string]string{
			"commandType": "main",
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("create called")
		},
	}

	c.AddCommand(
		createCommand(),
		deleteCommand(),
	)

	return c
}
