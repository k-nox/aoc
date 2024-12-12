package gen

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func getInput(session string, URL string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: session,
	}

	req.AddCookie(cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 status received from adventofcode: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}

	return body, nil
}

func buildUrl(base string, day int, year int) (*url.URL, error) {
	URL, err := url.Parse(base)
	if err != nil {
		return nil, fmt.Errorf("error parsing url: %w", err)
	}

	return URL.JoinPath(strconv.Itoa(year), "day", strconv.Itoa(day), "input"), nil
}
