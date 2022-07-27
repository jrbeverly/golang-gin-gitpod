package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/jrbeverly/cobra-cmd-with-docs/internal/backend"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		gin.SetMode(gin.ReleaseMode)

		router := backend.InitializeRoutes()
		router.LoadHTMLGlob("templates/*")

		router.Run()
	},
}
