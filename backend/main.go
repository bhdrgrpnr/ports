package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"saasteamtest/backend/internal"

	log "github.com/sirupsen/logrus"
	"saasteamtest/backend/logging"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})
	logging.SetLogOutput()
	level, err := log.ParseLevel("debug")
	if err != nil {
		fmt.Println("parse log level:", err)
		os.Exit(1)
	}
	log.SetLevel(level)

	if err := rootCmd.Execute(); err != nil {
		log.Error("Error calling rootCmd.Execute():", err)
		os.Exit(1)
	}
	if err := recover(); err != nil {
		fmt.Println("Error:", err)
	}
}

var rootCmd = &cobra.Command{
	Use:          "backend",
	Short:        "Serve the SAAS Backend Test API",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		logging.LogInfo(ctx, "start api", "SAAS Backend Test running on port 8000")
		log.Fatal(http.ListenAndServe(
			":8000",
			internal.RouterInitializer(),
		))
		return nil
	},
}
