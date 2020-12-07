package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

func main() {
	fmt.Println("Hello, world!")

	recursive := false
	_ = survey.AskOne(&survey.Confirm{
		Message: "Recursive search?",
	}, &recursive)

	reload := false
	_ = survey.AskOne(&survey.Confirm{
		Message: "Reload?",
	}, &reload)

	open := true
	_ = survey.AskOne(&survey.Confirm{
		Message: "Open?",
		Default: true,
	}, &open)

}
