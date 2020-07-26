package sitemap

import (
	"io"
)

// Sitemap is a sitemap
type Sitemap struct{}

// Parse parses the sitemap
func Parse(r io.Reader) (Sitemap, error) {
	return Sitemap{}, nil
}
