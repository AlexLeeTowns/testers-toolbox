package crawler_test

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/AlexLeeTowns/testers-toolbox/pkg/crawler"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var (
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func TestFetch(t *testing.T) {
	crawler.Client = &MockClient{}
	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(`<a href="fml">Hello, world</a>`)),
		}, nil
	}
	t.Run("Test Fetch", func(t *testing.T) {
		// Arrange
		f := make(crawler.Fetchy, 0)
		want := []string{"fml"}
		var got []string

		// Act
		res, err := f.Fetch("http://fake.url")
		if err != nil {
			t.Errorf("unexpected error :%v", err)
		}
		got = res.Urls

		// Assert
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestCrawl(t *testing.T) {
	url := "http://fake.url"
	crawler.Client = &MockClient{}
	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Body: io.NopCloser(bytes.NewBufferString(`<a href="http://fml.com">Hello, world</a>
<a href="http://lmao.lol">bla</a>`)),
		}, nil
	}
	t.Run("Should generate urls from Fetch", func(t *testing.T) {
		// Arrange
		f := crawler.Fetchy{}
		got := &f
		var buf bytes.Buffer
		want := map[string]*crawler.Result{
			url: {
				Body: url,
				Urls: []string{"http://fml.com", "http://lmao.lol"},
			},
		}

		// Act
		crawler.Crawl(&buf, url, 1, &f)

		// Assert
		if !reflect.DeepEqual((*got)[url].Urls, want[url].Urls) {
			t.Errorf("Got %v, expected %v", (*got)[url].Urls, want[url].Urls)
		}
	})

	t.Run("Should run Crawl for each url and return map", func(t *testing.T) {
		// Arrange
		f := crawler.Fetchy{}
		got := &f
		var buf bytes.Buffer
		want := map[string]*crawler.Result{
			url: {
				Body: url,
				Urls: []string{"http://fml.com", "http://lmao.lol"},
			},
			"http://fml.com": {
				Body: "http://fml.com",
				Urls: []string{"http://fml.com", "http://lmao.lol"},
			},
			"http://lmao.lol": {
				Body: "http://lmao.lol",
				Urls: []string{"http://fml.com", "http://lmao.lol"},
			},
		}

		// Act
		crawler.Crawl(&buf, url, 2, &f)
		// Assert
		for key := range want {
			if (*got)[key] == nil {
				t.Errorf("key %s not in got map: %v", key, got)
			}
		}
	})
}

func TestMerge(t *testing.T) {
	// Arrange
	firstMap := crawler.Fetchy{
		"foo": &crawler.Result{
			Body: "bar",
			Urls: []string{"first url"},
		},
	}
	secondMap := crawler.Fetchy{
		"hello": &crawler.Result{
			Body: "baz",
			Urls: []string{"second url"},
		},
	}
	want := map[string]*crawler.Result{
		"foo": {
			Body: "bar",
			Urls: []string{"first url"},
		},
		"hello": {
			Body: "baz",
			Urls: []string{"second url"},
		},
	}

	// Act
	got := crawler.Merge(firstMap, secondMap)

	// Assert
	for key := range want {
		if got[key] == nil {
			t.Errorf("key %s not in got map: %v", key, got)
		}
	}
}
