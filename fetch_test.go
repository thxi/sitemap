package sitemap

import (
	"testing"
)

// func TestFetch(t *testing.T) {
// 	t.Run("no error", func(t *testing.T) {
// 		u, _ := url.Parse("https://example.com")
// 		r, err := Fetch(u)
// 		if err != nil {
// 			t.Fatalf("error: %v", err)
// 		}
// 		r.Close()
// 	})
// 	t.Run("site not found", func(t *testing.T) {
// 		u, _ := url.Parse("https://lakdsjflkfaj.com")
// 		_, err := Fetch(u)
// 		if err == nil {
// 			t.Fatal("no error")
// 		}
// 	})
// 	t.Run("shit", func(t *testing.T) {
// 		u, _ := url.Parse("https://myanimelist.net")
// 		_, err := Fetch(u)
// 		fmt.Printf("Error type: %T\n", err)
// 		fmt.Printf("Error %v\n", err)
// 	})
// }

func TestFetchRobots(t *testing.T) {
	u := "https://myanimelist.net/"
	robots, err := fetchRobots(u)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	expected := []string{"https://myanimelist.net/sitemap/index.xml"}
	got := robots.Sitemaps
	if len(got) != len(expected) {
	}
	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("expected %v, got: %v", expected[i], got[i])
		}
	}
}

func TestFetchFromURL(t *testing.T) {
	t.Run("404", func(t *testing.T) {
		u := "https://reddit.com/sitemap.xml"
		// u := "http://localhost:1337/sitemap.xml"
		_, err := FetchFromURL(u)
		if err == nil {
			t.Fatal("no error")
		}
	})

	t.Run("200", func(t *testing.T) {
		u := "https://myanimelist.net/sitemap/index.xml"
		r, err := FetchFromURL(u)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		r.Close()
	})
}
