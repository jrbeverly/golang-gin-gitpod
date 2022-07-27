package main

import (
	"fmt"
	"log"
	"os"

	cobradocs "github.com/jrbeverly/cobra-cmd-with-docs/cmd/cobradocs/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var docsDirectory string

func main() {
	var docCmd = &cobra.Command{
		Use: "cobradocs",
		Run: func(cmd *cobra.Command, args []string) {
			err := doc.GenMarkdownTree(cobradocs.RootCmd, docsDirectory)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	docCmd.Flags().StringVar(&docsDirectory, "dir", "", "directory of the docs")
	docCmd.MarkFlagRequired("dir")

	if err := docCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
