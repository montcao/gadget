package cmd

import (
    "github.com/spf13/cobra"
)

var Root = New()

func New() *cobra.Command {
    root := &cobra.Command{
        Use:   "gadget",
        Short: "Inspect your images",
        Long:  `A programatic package to inspect OCI compliant images.`,
        Run: func(cmd *cobra.Command, args []string) {
            cmd.Help() // Display help if no subcommands are provided
        },
    }
    addInspectCommand(root)
    return root
}
