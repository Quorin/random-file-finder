package search

import (
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

	dir, err := filepath.Glob("*")
	if err != nil {
		return nil, err
	}

	for _, s := range dir {
		files = append(files, &File{
			Path: s,
			Name: s,
		})
	}

	return files, nil
}
