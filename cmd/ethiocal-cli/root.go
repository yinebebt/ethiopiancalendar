package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ethiocal",
	Short: "Ethiocal — Ethiopian Calendar (ባሕረ-ሐሳብ) and date converter",
	Long: `Ethiocal provides fasting and holiday dates based on the Ethiopian Orthodox
church calendar, and converts dates between the Ethiopian and Gregorian calendars.

Use subcommands (bahir, convert) for CLI access.
The graphical app ships separately as the Ethiocal desktop and mobile build.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return cmd.Help()
	},
}

// Execute runs the root command: bahir/convert subcommands; no args prints usage.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(bahirCmd)
	rootCmd.AddCommand(convertCmd)
}
