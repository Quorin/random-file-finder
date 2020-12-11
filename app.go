package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
	"random-file-finder/search"
)

func main() {
	recursive := false
	_ = survey.AskOne(&survey.Confirm{
		Message: "Recursive search?",
	}, &recursive)

	files, err := search.GetFiles(recursive)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}

	var pick *search.File

	for {
		pick = search.PickFile(files)
		fmt.Println("File:", pick.Name)

		reload := false
		_ = survey.AskOne(&survey.Confirm{
			Message: "Reload?",
			Default: false,
		}, &reload)

		if !reload {
			break
		}
	}

	open := true
	_ = survey.AskOne(&survey.Confirm{
		Message: "Open?",
		Default: true,
	}, &open)
}
