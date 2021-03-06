package crawler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

type Fetcher interface {
	Fetch(url string) (res *Result, err error)
}

type Result struct {
	Body       string
	StatusCode int
	Urls       []string
}

type Fetchy map[string]*Result

func (f Fetchy) Fetch(url string) (r *Result, err error) {
	var result = Result{
		Body: url,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("unexpected error: %v", err)
	}
	res, err := Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unexpected error: %v", err)
	}

	urls, err := getUrlsFromResponse(res)
	if err != nil {
		return nil, fmt.Errorf("unexpected error: %v", err)
	}

	result.Urls = urls
	result.StatusCode = res.StatusCode

	return &result, nil
}

func getUrlsFromResponse(res *http.Response) ([]string, error) {

	var result []string
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unexpected error: %v", err)
	}

	links := doc.Find("a")

	for _, link := range links.Nodes {
		for _, attr := range link.Attr {
			if attr.Key == "href" {
				result = append(result, attr.Val)
			}
		}
	}

	return result, nil

}

func Crawl(w io.Writer, url string, depth int, fetcher *Fetchy) {
	// TODO: Add mutex so it doesn't short out
	// TODO: Add concurrency
	if depth <= 0 {
		return
	}

	res, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println("Unexpected error")
	}
	(*fetcher)[url] = res
	fmt.Fprintf(w, "Url: %s; Status code: %d\n", res.Body, res.StatusCode)
	for _, u := range res.Urls {
		Crawl(w, u, depth-1, fetcher)
	}
}

func Merge(maps ...Fetchy) Fetchy {
	res := make(Fetchy)

	for _, m := range maps {
		for k, v := range m {
			res[k] = v
		}
	}

	return res
}
