package sitemap

import (
	"strings"
	"testing"
)

func TestParseSitemapIndex(t *testing.T) {
	r := strings.NewReader(sitemapindex)
	idx, err := ParseSitemapIndex(r)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestParseSitemap(t *testing.T) {
	r := strings.NewReader(urlset)
	_, err := ParseSitemap(r)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

var urlset = `
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url><loc>https://myanimelist.net</loc><changefreq>daily</changefreq><priority>1.0</priority></url>
  <url><loc>https://myanimelist.net/reviews.php</loc><changefreq>daily</changefreq><priority>0.8</priority></url>
  <url><loc>https://myanimelist.net/anime.php?letter=.</loc><changefreq>weekly</changefreq><priority>0.8</priority></url>
</urlset>
`

var sitemapindex = `
<?xml version="1.0" encoding="UTF-8"?>
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <sitemap>
    <loc>https://myanimelist.net/sitemap/main.xml</loc>
    <lastmod>2016-09-06</lastmod>
  </sitemap>
  <sitemap>
    <loc>https://myanimelist.net/sitemap/anime-000.xml</loc>
    <lastmod>2016-09-06</lastmod>
  </sitemap>
</sitemapindex>
`
