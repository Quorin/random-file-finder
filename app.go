package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
	"random-file-finder/search"
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

	files, err := search.GetFiles(recursive)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}

	for i, file := range files {
		fmt.Println("i:", i, " file:", file)
	}

}
