package main

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/skratchdot/open-golang/open"
	"os"
	"random-file-finder/search"
	"strings"
)

func main() {
	err := run()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}

func run() error {
	config := &search.Config{}

	if err := survey.AskOne(&survey.Confirm{
		Message: "Recursive search?",
	}, &config.Recursive); err != nil {
		return err
	}

	var extensions string

	if err := survey.AskOne(&survey.Input{
		Message: "File extensions?",
		Default: strings.Join(search.DefaultExtensions, " "),
		Help:    "Provide file extensions separated by space",
	}, &extensions); err != nil {
		return err
	}

	if err := survey.AskOne(&survey.Input{
		Message: "File name pattern?",
		Default: "",
		Help:    "Leave empty if not needed",
	}, &config.Pattern); err != nil {
		return err
	}

	config.Extensions = search.ParseExtensions(extensions)

	files, err := search.GetFiles(config)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New("not found files with provided settings")
	}

	var pick *search.File

	for {
		pick = search.PickFile(files)
		fmt.Println("👀", pick.Name)

		reload := false
		err := survey.AskOne(&survey.Confirm{
			Message: "Reload?",
			Default: false,
		}, &reload)

		if err != nil {
			return err
		}

		if !reload {
			break
		}
	}

	openFile := false
	err = survey.AskOne(&survey.Confirm{
		Message: "Open?",
		Default: true,
	}, &openFile)

	if err != nil {
		return err
	}

	if openFile {
		if err := open.Run(pick.Path); err != nil {
			return errors.New("cannot open file")
		}
	}
	return nil
}
