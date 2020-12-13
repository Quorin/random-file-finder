package search

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
)

func GetFiles(recursive bool) ([]*File, error) {
	if recursive {
		return getRecursiveFiles()
	}

	return getNonRecursiveFiles()
}

func PickFile(files []*File) *File {
	return files[rand.Intn(len(files)-1)]
}

func getRecursiveFiles() ([]*File, error) {
	var files []*File

	err := filepath.Walk(".", func(filepath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, &File{
			Path: filepath,
			Name: info.Name(),
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func getNonRecursiveFiles() ([]*File, error) {
	var files []*File

	dir, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, s := range dir {
		if s.IsDir() {
			continue
		}

		files = append(files, &File{
			Path: s.Name(),
			Name: s.Name(),
		})
	}

	return files, nil
}
