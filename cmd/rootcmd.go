package cmd

import "github.com/spf13/cobra"

func RootCmd() {
	rCmd := &cobra.Command{
		Use: "socksproxy",
	}

	rCmd.AddCommand()
}
