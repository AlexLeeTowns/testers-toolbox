package crawler

import (
	"fmt"
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
	Body string
	Urls []string
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
