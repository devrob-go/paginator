package main

import (
	"fmt"
	"net/http"
	"github.com/devrob-go/paginator"
)

func main() {
	// Example of how to use the Paginator package

	// Mock HTTP request (replace with actual request in a real scenario)
	req, _ := http.NewRequest("GET", "http://example.com/products", nil)

	// Simulate a paginator (current page 3, items per page 10, total items 50)
	paginator := paginator.NewPaginator(3, 10, 50, req)

	// Output the generated pagination details
	fmt.Printf("Showing range: %s\n", paginator.ShowingRange)

	// Display page URLs
	for _, page := range paginator.Pages {
		fmt.Printf("Page %d: %s (Current: %v)\n", page.Order, page.URL, page.Current)
	}

	// Display Prev/Next page URLs
	if paginator.Prev != nil {
		fmt.Printf("Previous page: %s\n", paginator.Prev.URL)
	}
	if paginator.Next != nil {
		fmt.Printf("Next page: %s\n", paginator.Next.URL)
	}
}
