package marsrover_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/r4d1n/marsrover"
)

func TestGetManifest(t *testing.T) {
	// return sample manifest data
	manifestHandler := func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile("./testdata/manifest_response.json")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}

	// create test server with handler
	server := httptest.NewServer(http.HandlerFunc(manifestHandler))
	defer server.Close()

	// instantiate client
	c := marsrover.NewClient("DEMO_KEY")
	c.OverrideBaseURL(server.URL)
	result, err := c.GetManifest("curiosity")

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// confirm result
	if result.Name != "Curiosity" {
		t.Errorf("Unexpected result: %v", result)
	}
	if result.MaxSol != 1658 {
		t.Errorf("Unexpected result: %v", result)
	}
}

func TestGetImagesBySol(t *testing.T) {
	// return sample photo data for a rover on a specific sol
	solHandler := func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile("./testdata/sol_response.json")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}

	// create test server with handler
	server := httptest.NewServer(http.HandlerFunc(solHandler))
	defer server.Close()

	// instantiate client
	c := marsrover.NewClient("DEMO_KEY")
	c.OverrideBaseURL(server.URL)
	photos, err := c.GetImagesBySol("curiosity", 1004)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// confirm result
	if len(photos) != 4 {
		t.Errorf("Unexpected result: %v", photos)
	}
	if photos[0].ID != 102685 {
		t.Errorf("Unexpected result: %v", photos)
	}
}
