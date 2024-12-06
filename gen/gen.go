package gen

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"text/template"
	"time"
)

func Generate(d int, moduleName string) error {
	pkgName := fmt.Sprintf("day%02d", d)
	err := createInputs(pkgName)
	if err != nil {
		return fmt.Errorf("error creating inputs: %w", err)
	}

	err = genPkg(d, pkgName)
	if err != nil {
		return fmt.Errorf("error generating package: %w", err)
	}

	err = genRegistry(d, moduleName)
	if err != nil {
		return fmt.Errorf("error generating registry: %w", err)
	}

	return nil
}

func createInputs(pkgName string) error {
	inputDir := fmt.Sprintf("input/%s", pkgName)
	err := createDirIfNotExist(inputDir)
	if err != nil {
		return err
	}

	err = createFileInDir(inputDir, "input.txt")
	if err != nil {
		return err
	}

	err = createFileInDir(inputDir, "sample.txt")
	if err != nil {
		return err
	}

	return nil
}

func createDirIfNotExist(name string) error {
	err := os.Mkdir(name, 0750)
	if err != nil && !errors.Is(err, fs.ErrExist) {
		return fmt.Errorf("failed to create %s dir: %w", name, err)
	}
	return nil
}

func createFileInDir(dir string, name string) error {
	fullPath := fmt.Sprintf("%s/%s", dir, name)
	f, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", fullPath, err)
	}
	defer f.Close()
	return nil
}

func genPkg(d int, pkgName string) error {
	err := createDirIfNotExist(pkgName)
	if err != nil {
		return err
	}

	for _, part := range []string{"One", "Two"} {
		err := genPartFile(d, pkgName, part)
		if err != nil {
			return err
		}
	}

	return nil
}

func genPartFile(d int, pkgName string, part string) error {
	tmpl, err := template.New("part").Parse(PartTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}
	f, err := os.Create(fmt.Sprintf("%s/part%s.go", pkgName, strings.ToLower(part)))
	if err != nil {
		return fmt.Errorf("error creating part%s.go: %w", strings.ToLower(part), err)
	}
	defer f.Close()
	err = tmpl.Execute(f, struct {
		Day  int
		Part string
	}{
		Day:  d,
		Part: part,
	})

	if err != nil {
		return fmt.Errorf("error execturing template: %w", err)
	}
	return nil
}

func genRegistry(d int, moduleName string) error {
	tmpl, err := template.New("registry").Parse(RegistryTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	f, err := os.Create("cli/registry.go")
	if err != nil {
		return fmt.Errorf("error creating registry.go: %w", err)
	}
	defer f.Close()

	days := []int{}

	for i := 1; i <= d; i++ {
		days = append(days, i)
	}

	// err = tmpl.Execute(f, struct{Days: days})
	err = tmpl.Execute(f, struct {
		Timestamp  time.Time
		Days       []int
		ModuleName string
	}{
		Timestamp:  time.Now(),
		Days:       days,
		ModuleName: moduleName,
	})

	if err != nil {
		return fmt.Errorf("error execturing template: %w", err)
	}

	return nil
}
