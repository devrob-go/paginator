package paginator

import (
	"net/http"
	"testing"
)

// Mock a basic HTTP request for testing purposes
func mockRequest() *http.Request {
	req, _ := http.NewRequest("GET", "http://example.com/products", nil)
	return req
}

// Test the paginator for a simple case
func TestPaginator_SimplePagination(t *testing.T) {
	req := mockRequest()
	p := NewPaginator(2, 10, 50, req)

	// Check the current page is 2
	if p.CurrentPage != 2 {
		t.Errorf("Expected current page to be 2, got %d", p.CurrentPage)
	}

	// Check the range of items being shown
	expectedRange := "11-20"
	if p.ShowingRange != expectedRange {
		t.Errorf("Expected showing range to be '%s', got '%s'", expectedRange, p.ShowingRange)
	}

	// Check pages (should have page 1, 2, and 3)
	if len(p.Pages) != 3 {
		t.Errorf("Expected 3 pages, got %d", len(p.Pages))
	}
	if p.Pages[0].Order != 1 {
		t.Errorf("Expected page 1, got page %d", p.Pages[0].Order)
	}
	if p.Pages[1].Order != 2 {
		t.Errorf("Expected page 2, got page %d", p.Pages[1].Order)
	}
	if p.Pages[2].Order != 3 {
		t.Errorf("Expected page 3, got page %d", p.Pages[2].Order)
	}
}

// Test the paginator for when the current page is the first page
func TestPaginator_FirstPage(t *testing.T) {
	req := mockRequest()
	p := NewPaginator(1, 10, 50, req)

	// Check the current page is 1
	if p.CurrentPage != 1 {
		t.Errorf("Expected current page to be 1, got %d", p.CurrentPage)
	}

	// Check the range of items being shown
	expectedRange := "1-10"
	if p.ShowingRange != expectedRange {
		t.Errorf("Expected showing range to be '%s', got '%s'", expectedRange, p.ShowingRange)
	}

	// Check pages (should have page 1 and 2)
	if len(p.Pages) != 2 {
		t.Errorf("Expected 2 pages, got %d", len(p.Pages))
	}
	if p.Pages[0].Order != 1 {
		t.Errorf("Expected page 1, got page %d", p.Pages[0].Order)
	}
	if p.Pages[1].Order != 2 {
		t.Errorf("Expected page 2, got page %d", p.Pages[1].Order)
	}

	// Check that the Prev page is nil
	if p.Prev != nil {
		t.Errorf("Expected Prev to be nil, got %v", p.Prev)
	}
}

// Test the paginator for when the current page is the last page
func TestPaginator_LastPage(t *testing.T) {
	req := mockRequest()
	p := NewPaginator(5, 10, 50, req)

	// Check the current page is 5 (last page)
	if p.CurrentPage != 5 {
		t.Errorf("Expected current page to be 5, got %d", p.CurrentPage)
	}

	// Check the range of items being shown
	expectedRange := "41-50"
	if p.ShowingRange != expectedRange {
		t.Errorf("Expected showing range to be '%s', got '%s'", expectedRange, p.ShowingRange)
	}

	// Check pages (should have pages 3, 4, 5)
	if len(p.Pages) != 3 {
		t.Errorf("Expected 3 pages, got %d", len(p.Pages))
	}
	if p.Pages[0].Order != 3 {
		t.Errorf("Expected page 3, got page %d", p.Pages[0].Order)
	}
	if p.Pages[1].Order != 4 {
		t.Errorf("Expected page 4, got page %d", p.Pages[1].Order)
	}
	if p.Pages[2].Order != 5 {
		t.Errorf("Expected page 5, got page %d", p.Pages[2].Order)
	}

	// Check that the Next page is nil
	if p.Next != nil {
		t.Errorf("Expected Next to be nil, got %v", p.Next)
	}
}

// Test the paginator for when there is only one page
func TestPaginator_OnePage(t *testing.T) {
	req := mockRequest()
	p := NewPaginator(1, 10, 5, req)

	// Check the current page is 1
	if p.CurrentPage != 1 {
		t.Errorf("Expected current page to be 1, got %d", p.CurrentPage)
	}

	// Check the range of items being shown
	expectedRange := "1-5"
	if p.ShowingRange != expectedRange {
		t.Errorf("Expected showing range to be '%s', got '%s'", expectedRange, p.ShowingRange)
	}

	// Check that there is only 1 page
	if len(p.Pages) != 1 {
		t.Errorf("Expected 1 page, got %d", len(p.Pages))
	}
	if p.Pages[0].Order != 1 {
		t.Errorf("Expected page 1, got page %d", p.Pages[0].Order)
	}

	// Check that Prev and Next are nil
	if p.Prev != nil {
		t.Errorf("Expected Prev to be nil, got %v", p.Prev)
	}
	if p.Next != nil {
		t.Errorf("Expected Next to be nil, got %v", p.Next)
	}
}

// Test the paginator when total number of items is less than perPage
func TestPaginator_LessThanPerPage(t *testing.T) {
	req := mockRequest()
	p := NewPaginator(1, 10, 8, req)

	// Check the current page is 1
	if p.CurrentPage != 1 {
		t.Errorf("Expected current page to be 1, got %d", p.CurrentPage)
	}

	// Check the range of items being shown
	expectedRange := "1-8"
	if p.ShowingRange != expectedRange {
		t.Errorf("Expected showing range to be '%s', got '%s'", expectedRange, p.ShowingRange)
	}

	// Check that there is only 1 page
	if len(p.Pages) != 1 {
		t.Errorf("Expected 1 page, got %d", len(p.Pages))
	}
	if p.Pages[0].Order != 1 {
		t.Errorf("Expected page 1, got page %d", p.Pages[0].Order)
	}

	// Check that Prev and Next are nil
	if p.Prev != nil {
		t.Errorf("Expected Prev to be nil, got %v", p.Prev)
	}
	if p.Next != nil {
		t.Errorf("Expected Next to be nil, got %v", p.Next)
	}
}
