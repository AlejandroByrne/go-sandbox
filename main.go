package main

import (
	"fmt"
	"net/http"

	"github.com/mmcdole/gofeed"
)

func getSubstackPosts() ([]*gofeed.Item, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://rubylarocca.substack.com/feed")

	if err != nil {
		return nil, err
	}

	return feed.Items, nil
}

// Helper to generate the HTML for the list of posts
func renderPosts(w http.ResponseWriter) {
	posts, _ := getSubstackPosts()
	for i, post := range posts {
		if i >= 3 {
			break
		}
		fmt.Fprintf(w, `
			<div class="bg-white p-6 rounded shadow hover:shadow-lg transition duration-500">
				<h2 class="text-xl font-semibold">%s</h2>
				<a href="%s" target="_blank" class="text-blue-500">Read &rarr;</a>
			</div>
		`, post.Title, post.Link)
	}
}

func main() {
	// 1. The Full Page Load
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<html>
		<head>
			<script src="https://cdn.tailwindcss.com"></script>
			<script src="https://unpkg.com/htmx.org@1.9.10"></script>
		</head>
		<body class="bg-gray-100 p-10">
			<div class="flex justify-between items-center mb-6">
				<h1 class="text-3xl font-bold">My Substack Feed</h1>
				
				<button hx-get="/refresh-feed" hx-target="#feed-container" 
						class="bg-black text-white px-4 py-2 rounded">
					Refresh Feed
				</button>
			</div>

			<div id="feed-container" class="grid gap-4">
				<p class="text-gray-500">Click refresh to load posts...</p>
			</div>
		</body>
		</html>`
		fmt.Fprint(w, html)
	})

	// 2. The Partial Endpoint (HTMX calls this)
	http.HandleFunc("/refresh-feed", func(w http.ResponseWriter, r *http.Request) {
		// Simulate network delay so you can see it happen
		// time.Sleep(1 * time.Second)
		renderPosts(w)
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
