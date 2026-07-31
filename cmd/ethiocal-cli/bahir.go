package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yinebebt/ethiocal/bahirehasab"
)

var bahirCmd = &cobra.Command{
	Use:   "bahir [year]",
	Short: "Get Ethiopian fasting and religious festival dates for a given year",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		year, err := strconv.Atoi(args[0])
		if err != nil || year < 0 {
			fmt.Println("Please provide a valid Ethiopian year.")
			return
		}

		festival, err := bahirehasab.NewFestival(year)
		if err != nil {
			fmt.Println("Error fetching bahire-hasab:", err)
			return
		}

		printFestivalInfo(festival)
	},
}

func printFestivalInfo(festival bahirehasab.Festival) {
	fmt.Printf("\nBahire-hasab Calendar for year %d\n", festival.Year.Year)
	fmt.Println("Year Information:")
	fmt.Printf("  Evangelist: %s\n", festival.Year.Evangelist)
	fmt.Printf("  New Year falls on: %s\n", festival.Year.NewYearWeekday)

	fmt.Println("\nBasic Information:")
	fmt.Printf("  Medeb: %d\n", festival.Basic.Medeb)
	fmt.Printf("  Wenber: %d\n", festival.Basic.Wenber)
	fmt.Printf("  Abektie: %d\n", festival.Basic.Abektie)
	fmt.Printf("  Metiq: %d\n", festival.Basic.Metiq)
	fmt.Printf("  Beale Metiq: %d\n", festival.Basic.BealeMetiq)
	fmt.Printf("  Mebaja Hamer: %d\n", festival.Basic.MebajaHamer)
	fmt.Printf("  Nenewie: %s\n", festival.Basic.Nenewie)

	fmt.Println("\nFasting Dates:")
	fmt.Printf("  Abiy Tsome: %s\n", festival.Fasting.Abiy)
	fmt.Printf("  Debre Zeit: %s\n", festival.Fasting.DebreZeit)
	fmt.Printf("  Hosanna: %s\n", festival.Fasting.Hosanna)
	fmt.Printf("  Siklet: %s\n", festival.Fasting.Siklet)
	fmt.Printf("  Tinsaye: %s\n", festival.Fasting.Tinsaye)
	fmt.Printf("  Rkbe Kahnat: %s\n", festival.Fasting.RkbeKahnat)
	fmt.Printf("  Dihnet: %s\n", festival.Fasting.Dihnet)
	fmt.Printf("  Hawariyat: %s\n", festival.Fasting.Hawariyat)
	fmt.Printf("  Erget: %s\n", festival.Fasting.Erget)
	fmt.Printf("  Peraklitos: %s\n", festival.Fasting.Peraklitos)
	fmt.Println("\nFixed Fasts:")
	fmt.Printf("  Nebiyat (ኅዳር 15): %s\n", festival.Fasting.Nebiyat)
	fmt.Printf("  Filseta (ነሐሴ 1): %s\n", festival.Fasting.Filseta)
	fmt.Printf("  Gehad (ጥር 10): %s\n", festival.Fasting.Gehad)
}
