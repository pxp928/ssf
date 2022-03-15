/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package ssfrepo

import (
	"fmt"

	"github.com/spf13/cobra"
)

func deleteCommand() *cobra.Command {
	eg := `Create a Task from ClusterTask 'foo' in namespace 'ns':
	tkn task create --from foo
or
	tkn task create foobar --from=foo -n ns`

	c := &cobra.Command{
		Use: "delete",
		//ValidArgsFunction: formatted.ParentCompletion,
		Short:   "Delete pipeline repo",
		Example: eg,
		Annotations: map[string]string{
			"commandType": "main",
		},
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("create called")
		},
	}
	//c.Flags().StringVarP(&opts.From, "from", "", "", "Create a ClusterTask from Task in a particular namespace")
	return c
}
