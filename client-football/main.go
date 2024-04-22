package main

import (
	"fmt"
	pkg "github.com/franklaercio/middleware-football/client-football/pkg/stub"
	"github.com/gookit/color"
)

const (
	PremierLeague = "39"
	LaLiga        = "140"
	SerieA        = "71"
)

func main() {

	app := pkg.NewMatchService()

	for {
		fmt.Println(color.FgGreen.Sprintf("Select a league:"))
		fmt.Println("1. Premier League")
		fmt.Println("2. La Liga")
		fmt.Println("3. Brasileirão Série A")
		fmt.Println("0. Exit")

		var choice string
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			return
		}

		switch choice {
		case "1":
			app.GetMatchDay(PremierLeague)
		case "2":
			app.GetMatchDay(LaLiga)
		case "3":
			app.GetMatchDay(SerieA)
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.")
			continue
		}
	}
}
