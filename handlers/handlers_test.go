package handlers

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHome(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("for home page, expected status 200 but got %d", resp.StatusCode)
	}

	// This part is supposed to fail in order to demonstrate the screenshot feature
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(bodyText), "awesome") {
		vel.TakeScreenShot(ts.URL, "HomeTest", 1500, 1000)
		t.Error("Did not find expected string in body")
	}
}
