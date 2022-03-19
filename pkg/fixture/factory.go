package fixture

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

type Model interface {
	path() string
	json() ([]byte, error)
}
type Data struct{}

func Create(m Model) ([]byte, error) {
	body, _ := m.json()
	path := fmt.Sprintf("%s%s", os.Getenv("BASEURL"), m.path())
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("error during http request creation: %v", err)
		return nil, err
	}

	res, err := Client.Do(req)
	if err != nil {
		fmt.Printf("error during request execution: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	fmt.Println(res.StatusCode)
	if res.StatusCode != 201 {
		fmt.Printf("entity not created, response: %v", res)
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, err
}
