package sitemap

import (
	"net/http"

	"github.com/temoto/robotstxt"
)

// fetchRobots fetches robots.txt file
// url should have a trailing slash
func fetchRobots(u string) (*robotstxt.RobotsData, error) {
	resp, err := http.Get(u + "robots.txt")
	// TODO: check if not found and shit
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return robotstxt.FromResponse(resp)
}
