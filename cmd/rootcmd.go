package cmd

import "github.com/spf13/cobra"

func RootCmd() *cobra.Command {
	rCmd := &cobra.Command{
		Use: "socksproxy",
	}

	start := &cobra.Command{
		Use: "start",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	rCmd.AddCommand(start)

	return &cobra.Command{}
}
