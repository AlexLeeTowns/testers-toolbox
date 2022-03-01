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
		f := make(crawler.Fetchy, 0)
		res, err := f.Fetch("http://fake.url")
		if err != nil {
			t.Errorf("unexpected error :%v", err)
		}
		got := res.Urls
		want := []string{"fml"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
