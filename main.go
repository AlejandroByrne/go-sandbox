package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/mmcdole/gofeed"
)

// Helper to get posts
func getPosts() []*gofeed.Item {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://rubylarocca.substack.com/feed")
	if err != nil {
		return []*gofeed.Item{}
	}
	// Return only first 3
	if len(feed.Items) > 3 {
		return feed.Items[:3]
	}
	return feed.Items
}

func main() {
	// ROUTE 1: The Home Page (Substack)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 1. Fetch Data
		posts := getPosts()

		// 2. Render Component
		component := SubstackPage(posts)
		templ.Handler(component).ServeHTTP(w, r)
	})

	// ROUTE 2: The Hello Page
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// 1. Extract Query Param
		// The button sends /hello?city=Austin
		city := r.URL.Query().Get("city")

		if city == "" {
			city = "Unknown"
		}

		// 2. Render Component
		component := HelloPage(city)
		templ.Handler(component).ServeHTTP(w, r)
	})

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
