package gen

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"golang.org/x/mod/modfile"
)

var (
	ErrNoModule        = errors.New("no go.mod found")
	ErrFileExists      = errors.New("unable to create file as it already exists")
	ErrNoValidTemplate = errors.New("no valid template found")
)

type Option func(*Generator)

func WithPath(path string) Option {
	return func(g *Generator) {
		g.path = path
	}
}

func WithForce(force bool) Option {
	return func(g *Generator) {
		g.force = force
	}
}

func WithModuleName(name string) Option {
	return func(g *Generator) {
		g.moduleName = name
	}
}

func WithPartTemplateFile(tmpl string) Option {
	return func(g *Generator) {
		g.partTemplate = templ{
			file: tmpl,
			name: "part",
		}
	}
}

func WithMainTemplateFile(tmpl string) Option {
	return func(g *Generator) {
		g.mainTemplate = templ{
			file: tmpl,
			name: "main",
		}
	}
}

func WithSession(session string) Option {
	return func(g *Generator) {
		g.session = session
	}
}

func WithBaseURL(baseURL string) Option {
	return func(g *Generator) {
		g.baseURL = baseURL
	}
}

type Generator struct {
	path         string
	force        bool
	moduleName   string
	partTemplate templ
	mainTemplate templ
	session      string
	baseURL      string
}

func New(opts ...Option) (*Generator, error) {
	generator := &Generator{
		mainTemplate: templ{
			text: MainTemplate,
			name: "main",
		},
		partTemplate: templ{
			text: PartTemplate,
			name: "part",
		},
		baseURL: "https://adventofcode.com",
	}

	for _, o := range opts {
		o(generator)
	}

	if generator.moduleName == "" {
		moduleName, err := generator.defaultModuleName()
		if err != nil {
			return nil, err
		}
		generator.moduleName = moduleName
	}

	return generator, nil
}

func (g *Generator) defaultModuleName() (string, error) {
	modFile, err := os.ReadFile(g.concatPath("go.mod"))
	if err != nil {
		return "", ErrNoModule
	}

	moduleName := modfile.ModulePath(modFile)
	if moduleName == "" {
		return "", ErrNoModule
	}

	return moduleName, nil
}

func (g *Generator) Generate(day int, year int) error {
	err := g.createInputs(day, year)
	if err != nil {
		return err
	}

	err = g.generateDailyPackage(day, year)
	if err != nil {
		return err
	}

	err = g.generateMain(year)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) createInputs(day int, year int) error {
	inputDir := g.concatPath("input", strconv.Itoa(year), formatDay(day))

	err := createDirIfNotExist(inputDir)
	if err != nil {
		return err
	}

	inpF, err := g.createFile(filepath.Join(inputDir, "input.txt"))
	if err != nil {
		return err
	}
	defer inpF.Close()

	if g.session != "" {
		URL, err := buildUrl(g.baseURL, day, year)
		if err != nil {
			return err
		}
		data, err := getInput(g.session, URL.String())
		if err != nil {
			return err
		}
		_, err = inpF.Write(data)
		if err != nil {
			return fmt.Errorf("error writing input file: %w", err)
		}
		err = inpF.Sync()
		if err != nil {
			return fmt.Errorf("error flushing input file to disk: %w", err)
		}
	}

	sampleF, err := g.createFile(filepath.Join(inputDir, "sample.txt"))
	if err != nil {
		return err
	}
	defer sampleF.Close()

	return nil
}

func (g *Generator) generateDailyPackage(day int, year int) error {
	path := g.concatPath(strconv.Itoa(year), formatDay(day))
	err := createDirIfNotExist(path)
	if err != nil {
		return err
	}

	for _, part := range []string{"One", "Two"} {
		err := g.generatePartFile(day, year, part)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) generatePartFile(day int, year int, part string) error {
	partTemplate, err := g.partTemplate.parse()
	if err != nil {
		return ErrNoValidTemplate
	}
	path := g.concatPath(strconv.Itoa(year), formatDay(day), fmt.Sprintf("part%s.go", strings.ToLower(part)))
	f, err := g.createFile(path)
	if err != nil {
		return fmt.Errorf("error creating part%s.go: %w", strings.ToLower(part), err)
	}
	defer f.Close()
	err = partTemplate.Execute(f, struct {
		Day  int
		Year int
		Part string
	}{
		Day:  day,
		Year: year,
		Part: part,
	})

	if err != nil {
		return fmt.Errorf("error execturing template: %w", err)
	}
	return nil
}

func (g *Generator) generateMain(year int) error {
	mainTemplate, err := g.mainTemplate.parse()
	if err != nil {
		return ErrNoValidTemplate
	}

	path := g.concatPath(strconv.Itoa(year), "main.go")
	mainFile, err := os.Create(path) // main will always get overwritten in current design
	if err != nil {
		return fmt.Errorf("error creating main.go: %w", err)
	}
	defer mainFile.Close()

	dirContents, err := os.ReadDir(g.concatPath(strconv.Itoa(year)))
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}
	days := []string{}
	for _, dirEntry := range dirContents {
		if dirEntry.IsDir() {
			days = append(days, dirEntry.Name())
		}
	}

	err = mainTemplate.Execute(mainFile, struct {
		Timestamp  time.Time
		Days       []string
		ModuleName string
		Year       int
	}{
		Timestamp:  time.Now(),
		Days:       days,
		ModuleName: g.moduleName,
		Year:       year,
	})

	if err != nil {
		return fmt.Errorf("error execturing template: %w", err)
	}

	return nil
}

func (g *Generator) concatPath(paths ...string) string {
	paths = append([]string{g.path}, paths...)
	return filepath.Join(paths...)
}

func formatDay(day int) string {
	return fmt.Sprintf("day%02d", day)
}
