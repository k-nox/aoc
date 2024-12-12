package gen_test

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/k-nox/aoc/gen"
	"github.com/stretchr/testify/suite"
)

type GeneratorTestSuite struct {
	suite.Suite
	dir        string
	testServer *httptest.Server
}

func (suite *GeneratorTestSuite) SetupSubTest() {
	dir, err := os.MkdirTemp("", "testing")
	suite.Require().NoError(err)
	suite.DirExists(dir)
	suite.dir = dir

	modfileContents := []byte("module github.com/k-nox/aoctesting\n\ngo 1.23.2")

	modfile, err := os.Create(filepath.Join(dir, "go.mod"))
	suite.Require().NoError(err)
	defer modfile.Close()
	_, err = modfile.Write(modfileContents)
	suite.Require().NoError(err)
	err = modfile.Sync()
	suite.Require().NoError(err)
}

func (suite *GeneratorTestSuite) TearDownSubTest() {
	if suite.testServer != nil {
		defer suite.testServer.Close()
	}

	err := os.RemoveAll(suite.dir)
	suite.Require().NoError(err)
	suite.NoDirExists(suite.dir)
	suite.dir = ""
}

func TestGeneratorTestSuite(t *testing.T) {
	suite.Run(t, new(GeneratorTestSuite))
}

func (suite *GeneratorTestSuite) TestGenerate() {
	subTests := []struct {
		name    string
		options func() []gen.Option
		setup   func()
		check   func(err error)
	}{
		{
			name: "should not write over existing input files if force is false",
			options: func() []gen.Option {
				return []gen.Option{
					gen.WithPath(suite.dir),
				}
			},
			setup: suite.precreateInputFiles,
			check: func(err error) { suite.ErrorIs(err, gen.ErrFileExists) },
		},
		{
			name: "should write over existing input files if force is true",
			options: func() []gen.Option {
				return []gen.Option{
					gen.WithPath(suite.dir),
					gen.WithForce(true),
				}
			},
			setup: suite.precreateInputFiles,
			check: suite.checkStandardGeneration,
		},
		{
			name: "should accept custom part template",
			options: func() []gen.Option {
				return []gen.Option{
					gen.WithPath(suite.dir),
					gen.WithPartTemplateFile(filepath.Join("testdata", "custom.tmpl")),
				}
			},
			setup: func() {},
			check: suite.checkCustomPartGeneration,
		},
		{
			name: "should accept custom main template",
			options: func() []gen.Option {
				return []gen.Option{
					gen.WithPath(suite.dir),
					gen.WithMainTemplateFile(filepath.Join("testdata", "custom.tmpl")),
				}
			},
			setup: func() {},
			check: suite.checkCustomMainGeneration,
		},
		{
			name: "should download input if session is given",
			options: func() []gen.Option {
				return []gen.Option{
					gen.WithPath(suite.dir),
					gen.WithSession("mock"),
					gen.WithBaseURL(suite.testServer.URL),
				}
			},
			setup: suite.setupTestServer,
			check: func(err error) {
				suite.checkStandardGeneration(err)
				suite.checkDownloadedInput()
			},
		},
	}

	for _, subTest := range subTests {
		suite.Run(subTest.name, func() {
			subTest.setup()
			generator, err := gen.New(subTest.options()...)
			suite.Require().NoError(err)
			subTest.check(generator.Generate(1, 2024))
		})
	}
}

func (suite *GeneratorTestSuite) precreateInputFiles() {
	suite.T().Helper()

	err := os.MkdirAll(filepath.Join(suite.dir, "input", "2024", "day01"), 0750)
	suite.Require().NoError(err)

	f, err := os.Create(filepath.Join(suite.dir, "input", "2024", "day01", "input.txt"))
	suite.Require().NoError(err)
	defer f.Close()
}

func (suite *GeneratorTestSuite) checkStandardGeneration(err error) {
	suite.T().Helper()

	suite.Require().NoError(err)
	suite.checkInputExists()
	suite.checkStandardPartFile()
	suite.checkStandardMainFile()
}

func (suite *GeneratorTestSuite) checkCustomPartGeneration(err error) {
	suite.T().Helper()
	suite.Require().NoError(err)

	suite.checkInputExists()
	expected := "hello"
	acutal, err := os.ReadFile(filepath.Join(suite.dir, "2024", "day01", "partone.go"))
	suite.Require().NoError(err)
	suite.Equal(expected, string(acutal))
	suite.checkStandardMainFile()
}

func (suite *GeneratorTestSuite) checkCustomMainGeneration(err error) {
	suite.T().Helper()
	suite.Require().NoError(err)

	suite.checkInputExists()
	suite.checkStandardPartFile()
	expected := "hello"
	acutal, err := os.ReadFile(filepath.Join(suite.dir, "2024", "main.go"))
	suite.Require().NoError(err)
	suite.Equal(expected, string(acutal))
}

func (suite *GeneratorTestSuite) checkInputExists() {
	suite.T().Helper()
	suite.FileExists(filepath.Join(suite.dir, "input", "2024", "day01", "input.txt"))
	suite.FileExists(filepath.Join(suite.dir, "input", "2024", "day01", "sample.txt"))
}

func (suite *GeneratorTestSuite) checkStandardPartFile() {
	suite.T().Helper()
	goldenPartFile, err := os.ReadFile(filepath.Join("testdata", "part.golden"))
	suite.Require().NoError(err)
	expectedPartFile := string(goldenPartFile)

	actualPartFile, err := os.ReadFile(filepath.Join(suite.dir, "2024", "day01", "partone.go"))
	suite.Require().NoError(err)
	suite.Equal(expectedPartFile, string(actualPartFile))
}

func (suite *GeneratorTestSuite) checkStandardMainFile() {
	suite.T().Helper()
	// because the main template has a timestamp generated, I'm doing this slightly hacky method to only compare everything after the first three lines
	expectedMainFile := suite.readFileWithoutNLines(3, filepath.Join("testdata", "main.golden"))
	actualMainFile := suite.readFileWithoutNLines(3, filepath.Join(suite.dir, "2024", "main.go"))
	suite.Equal(expectedMainFile, actualMainFile)
}

func (suite *GeneratorTestSuite) readFileWithoutNLines(n int, name string) string {
	suite.T().Helper()
	file, err := os.Open(name)
	suite.Require().NoError(err)
	scanner := bufio.NewScanner(file)
	contents := ""
	lines := 0
	for scanner.Scan() {
		if lines < n {
			continue
		}

		contents += scanner.Text()
	}

	return contents
}

func (suite *GeneratorTestSuite) setupTestServer() {
	suite.T().Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/2024/day/1/input"
		suite.Equal(expectedPath, r.URL.Path)

		expectedCookie := "mock"
		cookie, err := r.Cookie("session")
		suite.Require().NoError(err)
		suite.Equal(expectedCookie, cookie.Value)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`123`))
	}))

	suite.testServer = server
}

func (suite *GeneratorTestSuite) checkDownloadedInput() {
	suite.T().Helper()

	inpF, err := os.ReadFile(filepath.Join(suite.dir, "input", "2024", "day01", "input.txt"))
	suite.Require().NoError(err)
	suite.Equal(inpF, []byte(`123`))
}
