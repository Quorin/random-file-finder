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
	Delimiter         = " "
	AllExtensionsChar = "*"
)

func GetFiles(recursive bool, extensions []string, pattern string) ([]*File, error) {
	if recursive {
		return getRecursiveFiles(extensions, pattern)
	}

	return getNonRecursiveFiles(extensions, pattern)
}

func PickFile(files []*File) *File {
	if len(files) == 1 {
		return files[0]
	}
	rand.Seed(time.Now().UnixNano())
	return files[rand.Intn(len(files)-1)]
}

func ParseExtensions(extensions string) []string {
	if len(extensions) == 0 {
		return DefaultExtensions
	}

	// star == every extension
	if strings.TrimSpace(extensions) == AllExtensionsChar {
		return []string{AllExtensionsChar}
	}

	var exts []string
	split := strings.Split(extensions, Delimiter)

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

		exts = append(exts, e)
	}

	// if every extension is incorrect then return defautl extensions
	if len(exts) == 0 {
		return DefaultExtensions
	}

	return exts
}

func getRecursiveFiles(extensions []string, pattern string) ([]*File, error) {
	var files []*File

	err := filepath.Walk(".", func(fp string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !FindAnyExtension(extensions) && !SliceContain(extensions, filepath.Ext(info.Name())) {
			return nil
		}

		if len(pattern) > 0 {
			if !strings.Contains(strings.ToLower(filepath.Base(fp)), strings.ToLower(pattern)) {
				return nil
			}
		}

		files = append(files, &File{
			Path: fp,
			Name: info.Name(),
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func getNonRecursiveFiles(extensions []string, pattern string) ([]*File, error) {
	var files []*File

	dir, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, s := range dir {
		if s.IsDir() {
			continue
		}

		if !FindAnyExtension(extensions) && !SliceContain(extensions, filepath.Ext(s.Name())) {
			continue
		}

		if len(pattern) > 0 {
			if !strings.Contains(strings.ToLower(s.Name()), strings.ToLower(pattern)) {
				continue
			}
		}

		files = append(files, &File{
			Path: s.Name(),
			Name: s.Name(),
		})
	}

	return files, nil
}
