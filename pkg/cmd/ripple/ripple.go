package ripple

import (
	"github.com/spf13/cobra"
)

func NewCommand(name string) *cobra.Command {
	c := &cobra.Command{
		Use:   name,
		Short: "Starts simple HTTP server and waits for Stash push event",
		Long:  "bla bla bla",
	}

	return c
}
