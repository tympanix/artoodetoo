package api_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Tympanix/artoodetoo/api"
)

var (
	server *httptest.Server
)

func init() {
	server = httptest.NewServer(api.API)
}

func url(resource string) string {
	return server.URL + resource
}

func login(t *testing.T) string {
	login := `{"username":"admin","password":"admin"}`

	request, err := http.NewRequest("POST", url("/login"), strings.NewReader(login))

	if err != nil {
		t.Error(err)
	}

	response, _ := http.DefaultClient.Do(request)

	if response.StatusCode != 200 {
		t.Error("Wrong username or password")
		t.FailNow()
	}

	auth, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	return string(auth)
}

func TestEvents(t *testing.T) {
	// request, _ := http.NewRequest("GET", url("/all_events"), nil)
	// request.Header.Set("Authentication", login(t))
	//
	// res, err := http.DefaultClient.Do(request)
	//
	// if err != nil {
	// 	t.Error(err)
	// }
	//
	// if res.StatusCode != 200 {
	// 	t.Errorf("Success expected: %d", res.StatusCode)
	// }
}
