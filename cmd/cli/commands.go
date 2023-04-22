package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func helloCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "hello",
		Short: "A brief description of the hello command",
		Long:  "A longer description of the hello command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello!")
		},
	}
}
