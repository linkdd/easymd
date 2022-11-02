package easymd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"

	"github.com/linkdd/easymd/pkg/server"
)

var rootDocument string
var bindAddress net.IP
var bindPort int

var rootCmd = &cobra.Command{
	Use:   "easymd",
	Short: "easymd - a simple server rendering markdown documents to HTML",
	Run: func(cmd *cobra.Command, args []string) {
		server.Serve(rootDocument, bindAddress, bindPort)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(
		&rootDocument,
		"root",
		"r",
		".",
		"Root directory to scan for markdown documents",
	)
	rootCmd.Flags().IPVarP(
		&bindAddress,
		"bind",
		"b",
		net.IPv4zero,
		"IP address to listen on",
	)
	rootCmd.Flags().IntVarP(
		&bindPort,
		"port",
		"p",
		8000,
		"Port to listen on",
	)
}
