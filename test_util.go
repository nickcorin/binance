package binance

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func readTestData(t *testing.T) ([]byte, error) {
	path := filepath.Join("testdata", filepath.FromSlash(t.Name()+".json"))

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func createTestServer(t *testing.T, statusCode int) (*httptest.Server, error) {
	response, err := readTestData(t)
	if err != nil {
		return nil, err
	}

	h := func(w http.ResponseWriter, r *http.Request) {
		t.Log("In the handler")
		w.WriteHeader(statusCode)
		w.Write(response)
	}

	srv := httptest.NewServer(http.HandlerFunc(h))
	t.Logf("test server running on %s", srv.URL)
	return srv, nil
}
