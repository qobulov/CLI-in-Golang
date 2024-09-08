package main

import (
	"cli/cli"
	"github.com/c-bata/go-prompt"
	_ "github.com/c-bata/go-prompt"
)

func main() {
	p := prompt.New(
		cli.Executor,
		cli.Completer,
		prompt.OptionTitle("multifunctional cli"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionHistory(cli.History),
		prompt.OptionSuggestionBGColor(prompt.Green),
		prompt.OptionSuggestionTextColor(prompt.Black),
		prompt.OptionDescriptionBGColor(prompt.LightGray),
		prompt.OptionDescriptionTextColor(prompt.Black),
		prompt.OptionSelectedSuggestionBGColor(prompt.Blue),
		prompt.OptionSelectedSuggestionTextColor(prompt.Black),
	)

	p.Run()
}
