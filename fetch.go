package sitemap

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// FetchError is returned by FetchFromURL when something goes wrong.
// Inner error can be Unwrapped
type FetchError struct {
	inner error
}

func (e FetchError) Error() string {
	return fmt.Sprintf("Could not fetch the sitemap: %v", e.inner)
}

// Unwrap returns the inner error of FetchError for inspection
func (e FetchError) Unwrap() error {
	return e.inner
}

var defaultSitemaps = []string{
	"/sitemap.xml",
}

// Fetch tries to fetch the sitemap for a given website.
// Before fetching the sitemap,
func Fetch(u string) (io.ReadCloser, error) {
	// try fetch robots first
	if u[len(u)-1] != '/' {
		u += "/"
	}

	sitemaps := []string{}
	robots, err := fetchRobots(u)
	if err == nil {
		sitemaps = robots.Sitemaps
		if len(sitemaps) == 0 {
			sitemaps = defaultSitemaps
		}
	}

	for _, su := range sitemaps {
		// then look for default links
		r, err := FetchFromURL(su)
		if err != nil {
			continue
		}
		return r, nil
	}
	return nil, FetchError{errors.New("no sitemaps found")}
}

// FetchFromURL is a shorthand for FetchFromURLwithClient(u, &http.Client{})
func FetchFromURL(u string) (io.ReadCloser, error) {
	return FetchFromURLwithClient(u, &http.Client{})
}

// FetchFromURLwithClient fetches the given page
// Example: FetchFromURL("https://myanimelist.net/sitemap/index.xml")
func FetchFromURLwithClient(u string, c *http.Client) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, FetchError{err}
	}
	// TODO: use something smarter
	req.Header.Add("User-Agent", "Mozilla/5.0")
	resp, err := c.Do(req)
	if err != nil {
		return nil, FetchError{err}
	}
	if resp.StatusCode != http.StatusOK {
		return nil, FetchError{
			errors.New("Got status: " + strconv.Itoa(resp.StatusCode) + " when fetching sitemap")}
	}
	// TODO: might not be efficient to pass resp.Body
	return resp.Body, nil
}
