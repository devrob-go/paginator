# Paginator Package

A simple and flexible pagination package for Go that generates pagination data, including page numbers, previous/next links, and the range of items shown on the current page. The paginator is designed to be modular and easily extendable for future use.

## Installation

To install the package, run the following command:

```bash
go get github.com/devrob-go/paginator
```

## Usage


### How to Use

- **Install**: `go get github.com/devrob-go/paginator`
- **Import**: In your Go code, import the paginator package with `import "github.com/devrob-go/paginator"`.
- **Create a Paginator**: Call `NewPaginator()` with your pagination parameters to generate the paginator instance.
  
This package helps generate clean pagination logic in web applications, particularly useful when dealing with large datasets. It also provides flexible page URL generation based on the current HTTP request.


### Example

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/devrob-go/paginator"
)

func main() {
	// Create a new HTTP request (replace with an actual request in a real scenario)
	req, _ := http.NewRequest("GET", "http://example.com/products", nil)

	// Initialize paginator (current page 3, items per page 10, total items 50)
	paginator := paginator.NewPaginator(3, 10, 50, req)

	// Display the range of items being shown
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

```

## Functions
NewPaginator(currentPage, perPage, total int, req *http.Request) Paginator

Creates a new Paginator instance with the specified current page, items per page, and total number of items. The request URL is used to generate the correct page URLs.

    - currentPage: The current page number.
    - perPage: The number of items to display per page.
    - total: The total number of items to paginate.
    - req: The HTTP request, used to build the page URLs.

## Structs

    Paginator: Contains all pagination data:
        - Prev: A pointer to the previous page (if any).
        - Next: A pointer to the next page (if any).
        -Total: The total number of items.
        - PerPage: The number of items per page.
        - CurrentPage: The current page number.
        - ShowingRange: The range of items displayed on the current page (e.g., "1-10").
        - Pages: A list of Page structs representing the page links.

    Page: Represents a single page:
        - Order: The page number.
        - URL: The URL for the page.
        - Current: A boolean flag indicating whether this page is the current page.

## Example Output

For a paginator with 50 items, 10 items per page, and the current page set to 3:

```
Showing range: 21-30
Page 1: http://example.com/products?page=1 (Current: false)
Page 2: http://example.com/products?page=2 (Current: false)
Page 3: http://example.com/products?page=3 (Current: true)
Page 4: http://example.com/products?page=4 (Current: false)
Page 5: http://example.com/products?page=5 (Current: false)
Previous page: http://example.com/products?page=2
Next page: http://example.com/products?page=4
```

## Unit Tests

The paginator package includes unit tests to ensure proper functionality. You can run them using the Go test command:

```
go test -v
```

## Test Cases:

    - TestPaginator_SimplePagination: Validates pagination for a typical case with multiple pages.
    - TestPaginator_FirstPage: Validates pagination when the current page is the first page.
    - TestPaginator_LastPage: Validates pagination when the current page is the last page.
    - TestPaginator_OnePage: Tests pagination when there is only one page due to fewer items than the items per page.
    - TestPaginator_LessThanPerPage: Validates pagination when the total number of items is less than the number of items per page.

## License

This package is open-source and available under the MIT license. See LICENSE for more details.
