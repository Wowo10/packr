package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDummyFlowHandler(t *testing.T) {
	s := &Server{}
	mux := http.NewServeMux()
	mux.HandleFunc("/getPacks", s.GetPacksHandler)
	mux.HandleFunc("/solution", s.GetSolutionHandler)
	mux.HandleFunc("/addPacks", s.PostPacksHandler)
	mux.HandleFunc("/delPacks", s.DeletePacksHandler)

	server := httptest.NewServer(mux)
	defer server.Close()
	resp, err := http.Get(server.URL + "/getPacks")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	expected := "{\"packs\":[]}"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != string(body) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}

	_, err = http.Get(server.URL + "/addPacks?pack=420")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	_, err = http.Get(server.URL + "/addPacks?pack=7357")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	_, err = http.Get(server.URL + "/addPacks?pack=69")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}

	resp2, err := http.Get(server.URL + "/getPacks")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp2.Body.Close()
	if resp2.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp2.Status)
	}
	expected = "{\"packs\":[7357,420,69]}"
	body, err = io.ReadAll(resp2.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != string(body) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}

	_, err = http.Get(server.URL + "/delPacks?pack=7357")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}

	resp3, err := http.Get(server.URL + "/getPacks")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp3.Body.Close()
	if resp3.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp3.Status)
	}
	expected = "{\"packs\":[420,69]}"
	body, err = io.ReadAll(resp3.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != string(body) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}

	resp4, err := http.Get(server.URL + "/solution?amount=1")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp4.Body.Close()
	if resp4.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp4.Status)
	}
	expected = "{\"solution\":{\"69\":1}}"
	body, err = io.ReadAll(resp4.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != string(body) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}
}
