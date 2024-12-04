package paginator

import "net/http"

// Page represents an individual page in the pagination system.
type Page struct {
	Order   int    // The page number
	URL     string // The URL associated with the page
	Current bool   // True if this page is the current page
}

// generatePages generates the list of page numbers and URLs based on the current page and last page.
func generatePages(currentPage, lastPage int, req *http.Request) []Page {
	var pages []Page

	for i := 1; i <= lastPage; i++ {
		pageURL := buildPageURL(req, i)

		// Add page to the list based on various conditions (e.g., current page, first, last, etc.)
		if shouldIncludePage(i, currentPage, lastPage) {
			pages = append(pages, Page{
				Order:   i,
				URL:     pageURL,
				Current: i == currentPage,
			})
		}
	}

	return pages
}
