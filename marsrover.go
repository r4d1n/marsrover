package marsrover

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://api.nasa.gov/mars-photos/api/v1"

type RoverClient struct {
	Key string
	URL string
}

type manifestResponse struct {
	Manifest Manifest `json:"photo_manifest"`
}

// The manifest contains details about a rover's mission

type Manifest struct {
	Name        string
	LandingDate string `json:"landing_date"`
	LaunchDate  string `json:"launch_date"`
	Status      string
	MaxSol      int    `json:"max_sol"`
	MaxDate     string `json:"max_date"`
	TotalPhotos int    `json:"total_photos"`
	Photos      Photos `json:"photos"`
}

type Photos []*Sol

type Sol struct {
	Sol         int      `json:"sol"`
	TotalPhotos int      `json:"total_photos"`
	Cameras     []string `json:"cameras"`
}

func NewClient(key string, u string) *RoverClient {
	if u == "" {
		u = baseURL
	}
	return &RoverClient{
		Key: key,
		URL: u,
	}
}

func (c *RoverClient) GetManifest(rover string) (*Manifest, error) {
	url := fmt.Sprintf(c.URL+"/manifests/%s?api_key=%s", rover, c.Key)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data manifestResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data.Manifest, nil
}

func (c *RoverClient) GetImagesBySol(rover string, sol int) (*Photos, error) {
	url := fmt.Sprintf(c.URL+"/manifests/%s?api_key=%s", rover, c.Key)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Photos
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *RoverClient) doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
