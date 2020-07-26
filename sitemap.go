package sitemap

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"time"
)

// type definitions are stolen from https://github.com/kataras/sitemap

//revive:disable
// SitemapIndex is a <sitemapindex>
type SitemapIndex struct {
	//revive:enable
	XMLName    xml.Name `xml:"sitemapindex"`
	Xmlns      string   `xml:"xmlns,attr"`
	XmlnsXhtml string   `xml:"xmlns:xhtml,attr,omitempty"`

	URLs []URL `xml:"sitemap"`
}

// Sitemap is a <urlset>
type Sitemap struct {
	XMLName    xml.Name `xml:"urlset"`
	Xmlns      string   `xml:"xmlns,attr"`
	XmlnsXhtml string   `xml:"xmlns:xhtml,attr,omitempty"`

	URLs []URL `xml:"url"`
}

// URL is the parent tag for each URL entry.
type URL struct {
	// Loc is required. It defines the URL of the page.
	// This URL must begin with the protocol (such as http) and end with a trailing slash,
	// if your web server requires it. This value must be less than 2,048 characters.
	// Read more at: https://www.sitemaps.org/protocol.html#location
	Loc string `xml:"loc"`
	// LastMod is optional. It is the date of last modification of the file.
	LastMod time.Time `xml:"-"`
	// LastModStr do NOT set it directly,
	// other solution would be to use ptr or custom time marshaler but this will ruin the API's expressiveness.
	//
	// See internal `sitemap#Add`.
	LastModStr string `xml:"lastmod,omitempty"`
	// ChangeFreq is optional. Defines how frequently the page is likely to change.
	// This value provides general information to search engines and may not correlate exactly to how often they crawl the page.
	// Valid values are:
	// "always"
	// "hourly"
	// "daily"
	// "weekly"
	// "monthly"
	// "yearly"
	// "never"
	ChangeFreq string `xml:"changefreq,omitempty"`
	// Priority is optional. It defines the priority of this URL relative to other URLs on your site.
	// Valid values range from 0.0 to 1.0.
	//
	// The default priority of a page is 0.5.
	Priority float32 `xml:"priority,omitempty"`

	Links []Link `xml:"xhtml:link,allowempty,omitempty"`
}

// Link is the optional child element of a URL.
// It can be used to list every alternate version of the page.
//
// Read more at: https://support.google.com/webmasters/answer/189077?hl=en.
type Link struct {
	Rel      string `xml:"rel,attr"`
	Hreflang string `xml:"hreflang,attr"`
	Href     string `xml:"href,attr"`
}

// ParseSitemapIndex parses the <sitemapindex>
func ParseSitemapIndex(r io.Reader) (*SitemapIndex, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	s := &SitemapIndex{}
	err = xml.Unmarshal(b, s)
	return s, err
}

// ParseSitemap parses the <sitemapindex>
func ParseSitemap(r io.Reader) (*Sitemap, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	s := &Sitemap{}
	err = xml.Unmarshal(b, s)
	return s, err
}
