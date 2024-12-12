package gen

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func getInput(session string, day int, year int) ([]byte, error) {
	url, err := url.Parse("https://adventofcode.com")
	if err != nil {
		return nil, fmt.Errorf("error parsing url: %w", err)
	}
	url = url.JoinPath(strconv.Itoa(year), "day", strconv.Itoa(day), "input")
	cookie := &http.Cookie{
		Name:  "session",
		Value: session,
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.AddCookie(cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}
	return body, nil
}
