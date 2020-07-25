package sitemap

import (
	"fmt"
	"net/url"
	"testing"
)

func TestFetch(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		u, _ := url.Parse("https://example.com")
		r, err := Fetch(u)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		r.Close()
	})
	t.Run("site not found", func(t *testing.T) {
		u, _ := url.Parse("https://lakdsjflkfaj.com")
		_, err := Fetch(u)
		if err == nil {
			t.Fatal("no error")
		}
	})
	t.Run("shit", func(t *testing.T) {
		u, _ := url.Parse("https://myanimelist.net")
		_, err := Fetch(u)
		fmt.Printf("Error type: %T\n", err)
		fmt.Printf("Error %v\n", err)
	})
}
