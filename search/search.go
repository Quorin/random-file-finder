package search

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	DefaultExtensions = []string{".mkv", ".mp4", ".mpeg", ".avi", ".mov", ".m4ts", ".wmv", ".flv", ".webm", ".mpg"}
)

const (
	delimiter         = " "
	AllExtensionsChar = "*"
)

// GetFiles returns files depending on the configuration
func GetFiles(config *Config) ([]*File, error) {
	if config.Recursive {
		return getRecursiveFiles(config)
	}

	return getNonRecursiveFiles(config)
}

// PickFile returns random file
func PickFile(files []*File) *File {
	if len(files) == 1 {
		return files[0]
	}

	rand.Seed(time.Now().UnixNano())
	return files[rand.Intn(len(files)-1)]
}

func ParseExtensions(extensions string) []string {
	// if no provided extensions then return default
	if len(extensions) == 0 {
		return DefaultExtensions
	}

	// star == every extension
	if strings.TrimSpace(extensions) == AllExtensionsChar {
		return []string{AllExtensionsChar}
	}

	// split extensions by delimiter
	var parsedExtensions []string
	split := strings.Split(extensions, delimiter)

	for _, s := range split {
		// trim unneeded character
		e := strings.TrimFunc(s, func(r rune) bool {
			return r == ','
		})

		// trim space
		e = filepath.Ext(strings.TrimSpace(e))

		// omit if is not correct extension
		if len(e) == 0 {
			continue
		}

		parsedExtensions = append(parsedExtensions, e)
	}

	// if every extension is incorrect then return default extensions
	if len(parsedExtensions) == 0 {
		return DefaultExtensions
	}

	return parsedExtensions
}

// getRecursiveFiles returns all files in current directory and all subdirectories (recursively)
func getRecursiveFiles(config *Config) ([]*File, error) {
	var files []*File

	err := filepath.Walk(".", func(fp string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// return only files
		if info.IsDir() {
			return nil
		}

		// check extensions
		if !FindAnyExtension(config.Extensions) && !SliceContain(config.Extensions, filepath.Ext(info.Name())) {
			return nil
		}

		// check pattern
		if len(config.Pattern) > 0 {
			if !strings.Contains(strings.ToLower(filepath.Base(fp)), strings.ToLower(config.Pattern)) {
				return nil
			}
		}

		files = append(files, &File{
			Path: fp,
			Name: info.Name(),
		})

		// return nil if everything is correct
		return nil
	})

	// return error if filepath.Walk failed
	if err != nil {
		return nil, err
	}

	return files, nil
}

// getNonRecursiveFiles returns all files in current directory
func getNonRecursiveFiles(config *Config) ([]*File, error) {
	var files []*File

	dir, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, s := range dir {
		// return only files
		if s.IsDir() {
			continue
		}

		// check extensions
		if !FindAnyExtension(config.Extensions) && !SliceContain(config.Extensions, filepath.Ext(s.Name())) {
			continue
		}

		// check patterns
		if len(config.Pattern) > 0 {
			if !strings.Contains(strings.ToLower(s.Name()), strings.ToLower(config.Pattern)) {
				continue
			}
		}

		// path is not needed, because file is inside current directory
		files = append(files, &File{
			Path: s.Name(),
			Name: s.Name(),
		})
	}

	return files, nil
}
