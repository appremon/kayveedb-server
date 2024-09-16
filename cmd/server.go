
package cmd

import (
	"log"
	"github.com/appremon/kayveedb-server/daemon"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the KayveeDB server",
	Run: func(cmd *cobra.Command, args []string) {
		err := daemon.StartServer(port)
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	},
}

var port string

func init() {
	RootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&port, "port", "P", "3466", "Port to run the server on")
}
