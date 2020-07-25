package sitemap

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/temoto/robotstxt"
)

// Sitemap is a sitemap
type Sitemap struct{}

// ErrNotFound is returned when a site does not have a sitemap
// var ErrNotFound = errors.New("sitemap.xml not found")

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

// Fetch tries to fetch the sitemap for a given website
func Fetch(u *url.URL) (io.ReadCloser, error) {
	//TODO: pass string as an argument since http.Get() converts it anyway

	// try fetch robots first
	uString := u.String()
	if uString[len(uString)-1] != '/' {
		uString += "/"
	}
	respRobots, err := http.Get(uString + "robots.txt")
	// TODO: check if not found and shit
	if err != nil {
		return nil, FetchError{err}
	}
	defer respRobots.Body.Close()

	// TODO: parseRobots
	robots, err := robotstxt.FromResponse(respRobots)
	if err != nil {
		// continue I guess?
		fmt.Fprintln(os.Stderr, "could not parse the robots.txt")
	}

	sitemaps := robots.Sitemaps
	if len(sitemaps) == 0 {
		sitemaps = defaultSitemaps
	}
	for _, su := range sitemaps {
		// then look for default links
		resp, err := http.Get(su)
		if err != nil {
			return nil, FetchError{err}
		}
		// TODO: might not be efficient to pass resp.Body
		return resp.Body, nil
	}
	return nil, FetchError{errors.New("no sitemaps found")}
}

// Parse parses the sitemap
func Parse(r io.Reader) (Sitemap, error) {
	return Sitemap{}, nil
}
