package cli

import "github.com/c-bata/go-prompt"

func Completer(d prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest
	switch d.Text {
	case "e", "ex", "exi", "exit":
		s = []prompt.Suggest{
			{Text: "exit", Description: "exit the program"},
		}

	case "c", "cl", "cle", "clea", "clear":
		s = []prompt.Suggest{
			{Text: "clear", Description: "clear th cli"},
		}

	case "h", "he", "hel", "help":
		s = []prompt.Suggest{
			{Text: "help", Description: "see all features"},
		}

	case "w", "we", "wea", "weat", "weath", "weathe", "weather":
		s = []prompt.Suggest{
			{Text: "weather ", Description: "get current weather default Tashkent"},
		}
	case "g", "gy", "gym":
		s = []prompt.Suggest{
			{Text: "gym", Description: "Get gym with 30 day workout"},
		}
	case "cr", "cry", "crypt", "crypto":
		s = []prompt.Suggest{
			{Text: "crypto ", Description: "get any crypto prices"},
		}
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
