package fixture_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/AlexLeeTowns/testers-toolbox/pkg/fixture"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

//TODO: Tests can't parallelize because of global variable dependency.
var (
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func init() {
	fixture.Client = &MockClient{}
	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		s := fmt.Sprintf(`{"firstName":%q,"lastName":%q}`, "Alex", "Towns")
		return &http.Response{
			Status:     "201 CREATED",
			StatusCode: http.StatusCreated,
			Body:       io.NopCloser(bytes.NewBufferString(s)),
		}, nil
	}
}

func TestUserCreate(t *testing.T) {
	t.Run("Should create user", func(t *testing.T) {
		got, err := fixture.CreateUser("Alex", "Towns")
		if err != nil {
			t.Errorf("Error during user creation: %v", err)
		}

		if (*got).UserData.FirstName != "Alex" {
			t.Errorf("got %q, expected %q", (*got).UserData.FirstName, "Alex")
		}
		if (*got).UserData.LastName != "Towns" {
			t.Errorf("got %q, expected %q", (*got).UserData.LastName, "Towns")
		}
	})
}
