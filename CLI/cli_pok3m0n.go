package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	log.SetLevel(log.InfoLevel)
	var start = &cobra.Command{
		Use:              "start [LAB_NAME]",
		TraverseChildren: true,
		Short:            "Use for stop many labs.",
		Args:             cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			LabToBE(args[0], "start")
		},
	}
	var stop = &cobra.Command{
		Use:              "stop [LAB_NAME]",
		TraverseChildren: true,
		Short:            "Use for start many labs.",
		Args:             cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			LabToBE(args[0], "stop")
		},
	}

	var sign_up = &cobra.Command{
		Use:              "sign_up -u [USERNAME] -p [PASSWORD]",
		Short:            "Sign up account.",
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.Flag("username").Value.String() == "empty" || cmd.Flag("password").Value.String() == "empty" {
				fmt.Println("\nPassword or Username not been set.")
			} else {
				UserToBE("sign_up", cmd.Flag("username").Value.String(), cmd.Flag("password").Value.String())
			}
		},
	}

	var sign_in = &cobra.Command{
		Use:              "sign_in -u [USERNAME] -p [PASSWORD]",
		Short:            "Sign in account",
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.Flag("username").Value.String() == "empty" || cmd.Flag("password").Value.String() == "empty" {
				fmt.Println("\nPassword or Username not been set.")
			} else {
				UserToBE("sign_in", cmd.Flag("username").Value.String(), cmd.Flag("password").Value.String())
			}
		},
	}

	tab := []*cobra.Command{sign_up, sign_in}
	for _, p := range tab {
		p.PersistentFlags().StringP("username", "u", "empty", "--username=<user's name>")
		p.PersistentFlags().StringP("password", "p", "empty", "--password=<password>")
	}

	var rootCmd = &cobra.Command{Use: "cli_pok3m0n"}
	rootCmd.AddCommand(start, stop, sign_in, sign_up)
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
