package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func NewListCmd(t []byte, dir string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists the architecture decision records",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			files, err := os.ReadDir(dir)
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range files {
				fmt.Println(file.Name())
			}
		},
	}
	return cmd
}
