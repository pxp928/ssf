/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package ssfrepo

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/thesecuresoftwarefactory/ssf/cli/pkg/ssfrepo"
)

func createCommand() *cobra.Command {
	eg := `Create pipeline based on config.yaml
	ssf-cli ssfrepo create config.yaml`

	c := &cobra.Command{
		Use:     "create",
		Short:   "Create pipeline with tasks specified",
		Example: eg,
		Annotations: map[string]string{
			"commandType": "main",
		},
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				return errors.New("too many values passed in")
			}
			if len(args) != 0 {
				return ssfrepo.Create(args[0])
			}
			return errors.New("create new ssf repo failed to run")
		},
	}
	return c
}
