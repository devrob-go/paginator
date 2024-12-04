package paginator

import (
	"fmt"
	"math"
	"net/http"
)

// Paginator struct to hold pagination data.
type Paginator struct {
	Prev         *Page   // Previous page
	Next         *Page   // Next page
	Total        int     // Total number of items
	PerPage      int     // Items per page
	CurrentPage  int     // Current page
	ShowingRange string  // Range of items being displayed (e.g. 1-10)
	Pages        []Page  // List of page data
}

// NewPaginator creates a new Paginator instance based on the current page, items per page, and total items.
func NewPaginator(currentPage, perPage int, total int, req *http.Request) Paginator {
	p := Paginator{}
	p.PerPage = perPage
	p.Total = total
	p.CurrentPage = currentPage

	// Calculate the last page
	lastPage := int(math.Max(math.Ceil(float64(total)/float64(perPage)), 1))

	// Set the range of items being shown (e.g., 1-10)
	p.ShowingRange = getShowingRange(currentPage, perPage, total)

	// Generate the page URLs
	p.Pages = generatePages(currentPage, lastPage, req)

	// Add previous and next pages
	p.addPrevNext(currentPage, lastPage, req)

	return p
}

// getShowingRange returns the range of items being displayed on the current page.
func getShowingRange(currentPage, perPage, total int) string {
	numberOfItems := currentPage * perPage
	if numberOfItems > total {
		numberOfItems = total
	}
	return fmt.Sprintf("%d-%d", numberOfItems-perPage+1, numberOfItems)
}

// shouldIncludePage determines if a page should be included in the pagination.
func shouldIncludePage(i, currentPage, lastPage int) bool {
	if i == 1 || i == lastPage || (i >= currentPage-1 && i <= currentPage+1) {
		return true
	}
	if lastPage > 5 && (i == 2 || i == lastPage-1) {
		return true
	}
	return false
}
