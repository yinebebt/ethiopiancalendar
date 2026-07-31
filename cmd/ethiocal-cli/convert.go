package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/yinebebt/ethiocal/dateconverter"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert dates between Ethiopian and Gregorian calendars",
}

// parseDate parses a YYYY-MM-DD date string.
func parseDate(s string) (year, month, day int, err error) {
	parts := strings.Split(s, "-")
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid date %q, expected YYYY-MM-DD", s)
	}
	fields := []struct {
		name string
		out  *int
	}{
		{"year", &year},
		{"month", &month},
		{"day", &day},
	}
	for i, f := range fields {
		if *f.out, err = strconv.Atoi(parts[i]); err != nil {
			return 0, 0, 0, fmt.Errorf("invalid %s in %q", f.name, s)
		}
	}
	return year, month, day, nil
}

var gtoeCmd = &cobra.Command{
	Use:     "gtoe [date]",
	Short:   "Convert Gregorian date to Ethiopian date",
	Long:    "Convert a Gregorian date (YYYY-MM-DD) to Ethiopian.",
	Args:    cobra.ExactArgs(1),
	Example: "ethiocal-cli convert gtoe 2025-02-1",
	Run: func(cmd *cobra.Command, args []string) {
		year, month, day, err := parseDate(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		etDate, err := dateconverter.Ethiopian(year, month, day)
		if err != nil {
			fmt.Println("Error converting date:", err)
			return
		}

		fmt.Printf("\nGregorian Date: %s\n", dateconverter.FormatGregorian(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)))
		fmt.Printf("  %04d-%02d-%02d\n", year, month, day)
		if long, err := dateconverter.FormatEthiopian(etDate.Year(), int(etDate.Month()), etDate.Day()); err == nil {
			fmt.Println("Converted Ethiopian Date:", long)
		}
		fmt.Println(" ", etDate.Format("2006-01-02"))
	},
}

var etogCmd = &cobra.Command{
	Use:     "etog [date]",
	Short:   "Convert Ethiopian date to Gregorian date",
	Long:    "Convert an Ethiopian date (YYYY-MM-DD) to Gregorian.",
	Args:    cobra.ExactArgs(1),
	Example: "ethiocal-cli convert etog 2017-5-25",
	Run: func(cmd *cobra.Command, args []string) {
		year, month, day, err := parseDate(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		gregDate, err := dateconverter.Gregorian(year, month, day)
		if err != nil {
			fmt.Println("Error converting date:", err)
			return
		}
		if long, err := dateconverter.FormatEthiopian(year, month, day); err == nil {
			fmt.Println("\nEthiopian Date:", long)
		}
		fmt.Printf("  %04d-%02d-%02d\n", year, month, day)
		fmt.Println("Converted Gregorian Date:", dateconverter.FormatGregorian(gregDate))
		fmt.Println(" ", gregDate.Format("2006-01-02"))
	},
}

func init() {
	convertCmd.AddCommand(gtoeCmd)
	convertCmd.AddCommand(etogCmd)
}
