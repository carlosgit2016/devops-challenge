package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestHandler(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "testfiles")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after the test.

	testFileName := "index.html"
	testFilePath := filepath.Join(tempDir, testFileName)
	testFileContent := "<!doctype html><html><head></head><body></body></html>"
	if err := ioutil.WriteFile(testFilePath, []byte(testFileContent), 0666); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	handler := Handler(tempDir)

	ts := httptest.NewServer(handler)
	defer ts.Close() // Close the server when test finishes.

	resp, err := http.Get(ts.URL + "/" + testFileName)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if string(body) != testFileContent {
		t.Errorf("Expected response body to be %q, got %q", testFileContent, string(body))
	}
}
