package paginator

import (
	"net/http"
	"strconv"
)

// buildPageURL is a utility function to construct a URL with the given page number.
func buildPageURL(req *http.Request, page int) string {
	pageURL := *req.URL
	params := pageURL.Query()
	params.Set("page", strconv.Itoa(page))
	pageURL.RawQuery = params.Encode()
	return pageURL.String()
}

// addPrevNext adds the previous and next page information to the paginator.
func (p *Paginator) addPrevNext(currentPage, lastPage int, req *http.Request) {
	// Previous page
	if currentPage > 1 {
		p.Prev = &Page{URL: buildPageURL(req, currentPage-1)}
	}

	// Next page
	if currentPage < lastPage {
		p.Next = &Page{URL: buildPageURL(req, currentPage+1)}
	}
}
