package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/skratchdot/open-golang/open"
	"os"
	"random-file-finder/search"
	"strings"
)

func main() {
	// TODO regex matching (optional)
	var recursive bool

	if err := survey.AskOne(&survey.Confirm{
		Message: "Recursive search?",
	}, &recursive); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	var extensions string

	if err := survey.AskOne(&survey.Input{
		Message: "File extensions?",
		Default: strings.Join(search.DefaultExtensions, " "),
		Help:    "Provide file extensions separated by space",
	}, &extensions); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	files, err := search.GetFiles(recursive, search.ParseExtensions(extensions))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	if len(files) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "not found files with provided settings")
		return
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

	openFile := true
	_ = survey.AskOne(&survey.Confirm{
		Message: "Open?",
		Default: true,
	}, &openFile)

	if openFile {
		if err := open.Run(pick.Path); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "cannot open file")
		}
	}
}
