package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func selectMode(cmd *cobra.Command, args []string) {
	//? Switch possible ?
	if cmd.Flag("verbose").Value.String() == "true" {
		log.SetLevel(log.InfoLevel)
	} else if cmd.Flag("debug").Value.String() == "true" {
		log.SetLevel(log.DebugLevel)
	} else if cmd.Flag("trace").Value.String() == "true" {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.ErrorLevel)

	}
}

func main() {
	log.SetLevel(log.InfoLevel)
	var start = &cobra.Command{
		Use:              "start [LAB_NAME]",
		TraverseChildren: true,
		Short:            "Use for stop many labs.",
		Args:             cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			selectMode(cmd, args)
			LabToBE(args[0], "start")
		},
	}
	var stop = &cobra.Command{
		Use:              "stop [LAB_NAME]",
		TraverseChildren: true,
		Short:            "Use for start many labs.",
		Args:             cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			selectMode(cmd, args)
			LabToBE(args[0], "stop")
		},
	}

	tab := []*cobra.Command{start, stop}
	for _, p := range tab {
		p.Flags().BoolP("trace", "t", false, "trace mode")
		p.Flags().BoolP("debug", "d", false, "debug mode")
		p.Flags().BoolP("verbose", "v", false, "verbose mode")
	}

	var rootCmd = &cobra.Command{Use: "cli_pok3m0n"}
	rootCmd.AddCommand(start, stop)
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
