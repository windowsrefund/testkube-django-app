package main_test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	url := "https://https://django-test.streethawk.com/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code error: expected %d but got %d", http.StatusOK, res.StatusCode)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(body), "The install worked successfully! Congratulations") {
		t.Errorf("Response body does not contain expected string")
	}
}
