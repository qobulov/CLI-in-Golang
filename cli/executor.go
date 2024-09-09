package cli

import (
	"cli/gym/storage"
	"cli/weather"
	"cli/crypto"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

var History []string
var gymStorage storage.Storage // Assume storage is already initialized and passed during setup

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
		fmt.Println("[weather <<country name>>] -> current weather of any country")
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
		handleCrypto(commands)

	case "gym":
		handleGym(commands)

	default:
		fmt.Println("command not found")
	}
}

func handleGym(commands []string) {
	if len(commands) < 2 {
		fmt.Println("Usage: gym [all|Get done|not_done|next|day <number>|done <day number>]")
		return
	}

	action := commands[1]

	switch action {
	case "all":
		tasks, err := gymStorage.GetAllTasks()
		if err != nil {
			fmt.Println("Error fetching all tasks:", err)
			return
		}
		for _, task := range tasks {
			fmt.Printf("Day %d: %s - Done: %v\n", task.Day, task.Works, task.IsDone)
		}

	case "Get_done":
		tasks, err := gymStorage.GetDoneTasks()
		if err != nil {
			fmt.Println("Error fetching done tasks:", err)
			return
		}
		for _, task := range tasks {
			fmt.Printf("Day %d: %s - Done: %v\n", task.Day, task.Works, task.IsDone)
		}

	case "not_done":
		tasks, err := gymStorage.GetNotDoneTasks()
		if err != nil {
			fmt.Println("Error fetching not done tasks:", err)
			return
		}
		for _, task := range tasks {
			fmt.Printf("Day %d: %s - Done: %v\n", task.Day, task.Works, task.IsDone)
		}

	case "next":
		task, err := gymStorage.GetNextTasks()
		if err != nil {
			fmt.Println("Error fetching next task:", err)
			return
		}
		fmt.Printf("Next Task - Day %d: %s - Done: %v\n", task.Day, task.Works, task.IsDone)

	case "day":
		if len(commands) < 3 {
			fmt.Println("Usage: gym day <day number>")
			return
		}
		day, err := strconv.Atoi(commands[2])
		if err != nil {
			fmt.Println("Invalid day number:", err)
			return
		}
		task, err := gymStorage.GetByDay(day)
		if err != nil {
			fmt.Println("Error fetching task by day:", err)
			return
		}
		fmt.Printf("Day %d: %s - Done: %v\n", task.Day, task.Works, task.IsDone)

	case "done":
		if len(commands) < 3 {
			fmt.Println("Usage: gym done <day number>")
			return
		}
		day, err := strconv.Atoi(commands[2])
		if err != nil {
			fmt.Println("Invalid day number:", err)
			return
		}
		err = gymStorage.DoDone(day)
		if err != nil {
			fmt.Println("Error marking task as done:", err)
			return
		}
		fmt.Printf("Marked Day %d as done\n", day)

	default:
		fmt.Println("Unknown action for gym command.")
	}
}

func handleCrypto(commands []string) {
	if len(commands) < 2 {
		fmt.Println("Usage: crypto <crypto name>")
		return
	}

	currency := commands[1]
	price, err := crypto.GetCryptoPrice(currency)
	if err != nil {
		fmt.Printf("Error fetching price for %s: %v\n", currency, err)
		return
	}

	fmt.Printf("Current price of %s: %.2f USD\n", currency, price.Price)
}
