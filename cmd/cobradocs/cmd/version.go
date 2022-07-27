package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string
var commit string
var tool string

func init() {
	RootCmd.AddCommand(versionCmd)
}

func _versionConvention(tool string, version string, commit string) string {
	if version == "" {
		if commit == "" {
			return fmt.Sprintf("%s/nover git/nostamp", tool)
		} else {
			return fmt.Sprintf("%s/nover git/%s", tool, commit)
		}
	}

	if commit == "" {
		return fmt.Sprintf("%s/%s git/nostamp", tool, version)
	}
	return fmt.Sprintf("%s/%s git/%s", tool, version, commit)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `All software has versions. This is mine.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(_versionConvention(tool, version, commit))
	},
}
