package gen

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func (g *Generator) createFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err == nil {
		return f, nil
	}

	if !errors.Is(err, fs.ErrExist) {
		return nil, fmt.Errorf("error creating file %s: %w", path, err)
	}

	if !g.force {
		return nil, ErrFileExists
	}

	f, err = os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("error creating file %s: %w", path, err)
	}

	return f, nil
}

func createDirIfNotExist(name string) error {
	err := os.MkdirAll(name, 0750)
	if err != nil && !errors.Is(err, fs.ErrExist) {
		return fmt.Errorf("failed to create %s dir: %w", name, err)
	}

	return nil
}
