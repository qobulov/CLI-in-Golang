package cli

import (
	"cli/weather"
	"cli/crypto"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
	"time"
)

var History []string

func Executor(s string) {
	History = append(History, s)

	s = strings.TrimSpace(s)
	commands := strings.Split(s, " ")

	switch commands[0] {

	case "exit":
		fmt.Println("bye bye !!!")
		os.Exit(0)

	case "help":
		fmt.Println("[gym] -> gym with 30 day workout")
		fmt.Println("[weather <<country name>>] -> current weather any country")
		fmt.Println("[crypto <<crypto name>>] -> get any crypto prices")

	case "clear":
		print("\033[2J")
		print("\033[H")
		print("\033[3J")

	case "weather":
		var country string
		if len(commands) >= 2 {
			country = commands[1]
		} else {
			country = "Tashkent"
		}
		wh := weather.GetWeather(country)

		location, current, hours := wh.Location, wh.Current, wh.Forecast.Forecastday[0].Hour

		fmt.Printf(
			"%s, %s: %.0fC, %s\n",
			location.Name,
			location.Country,
			current.TempC,
			current.Condition.Text,
		)

		for _, hour := range hours {
			date := time.Unix(hour.TimeEpoch, 0)

			if date.Before(time.Now()) {
				continue
			}

			message := fmt.Sprintf(
				"%s - %.0fC°, Chance of rain: %.0f%%, %s\n",
				date.Format("15:04"),
				hour.TempC,
				hour.ChanceOfRain,
				hour.Condition.Text,
			)

			if hour.ChanceOfRain < 40 {
				fmt.Print(message)
			} else {
				color.Red(message)
			}
		}

	case "crypto":
		var cryptoName string
		if len(commands) >= 2 {
			cryptoName = commands[1]
		} else {
			cryptoName = "bitcoin" // Default qiymat
		}

		cryptoInfo, err := crypto.GetCryptoPrice(cryptoName)
		if err != nil {
			fmt.Printf("Xatolik: %v\n", err)
		} else {
			fmt.Printf("%s uchun joriy narx: %.2f USD\n", cryptoInfo.Symbol, cryptoInfo.Price)
		}

	default:
		fmt.Println("command not found")
	}
}
