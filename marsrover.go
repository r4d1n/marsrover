// package marsrover provides a client for the NASA mars rover images API

package marsrover

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://api.nasa.gov/mars-photos/api/v1"

type Client struct {
	Key string
	URL string
}

type manifestResponse struct {
	Manifest Manifest `json:"photo_manifest"`
}

// The manifest contains details about a rover's mission

type Manifest struct {
	Name        string `json:"name"`
	LandingDate string `json:"landing_date"`
	LaunchDate  string `json:"launch_date"`
	Status      string `json:"status"`
	MaxSol      int    `json:"max_sol"`
	MaxDate     string `json:"max_date"`
	TotalPhotos int    `json:"total_photos"`
	Sols        []Sol  `json:"photos"`
}

// Sol contains data for the rover's photo activity on a given martian sol

type Sol struct {
	Sol         int      `json:"sol"`
	TotalPhotos int      `json:"total_photos"`
	Cameras     []string `json:"cameras"`
}

type photoResponse struct {
	Photos []Photo `json:"photos"`
}

// Photo represents an image and related metadata

type Photo struct {
	ID        int    `json:"id"`
	Sol       int    `json:"sol"`
	Camera    Camera `json:"camera"`
	ImgSrc    string `json:"img_src"`
	EarthDate string `json:"earth_date"`
	Rover     Rover  `json:"rover"`
}

// Rover contains information about a given rover

type Rover struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	LandingDate string   `json:"landing_date"`
	LaunchDate  string   `json:"launch_date"`
	Status      string   `json:"status"`
	MaxSol      int      `json:"max_sol"`
	MaxDate     string   `json:"max_date"`
	TotalPhotos int      `json:"total_photos"`
	Cameras     []Camera `json:"cameras"`
}

// Rover contains information about a rover camera

type Camera struct {
	ID        int    `json:"id, omitempty"`
	ShortName string `json:"name"`
	RoverID   int    `json:"rover_id, omitempty"`
	FullName  string `json:"full_name"`
}

func NewClient(key string) *Client {
	if key == "" {
		key = "DEMO_KEY"
	}
	return &Client{
		Key: key,
		URL: baseURL,
	}
}

// Fetch a rover mission manifest

func (c *Client) GetManifest(rover string) (*Manifest, error) {
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

// Fetch all photos taken by a specific rover on a particular martian sol

func (c *Client) GetImagesBySol(rover string, sol int) (*photoResponse, error) {
	url := fmt.Sprintf(c.URL+"/rovers/%s/photos?sol=%d&api_key=%s", rover, sol, c.Key)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data *photoResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Fetch all photos taken by a specific rover on a particular earth date

func (c *Client) GetImagesByEarthDate(rover string, date string) (*photoResponse, error) {
	url := fmt.Sprintf(c.URL+"/rovers/%s/photos?earth_date=%s&api_key=%s", rover, date, c.Key)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data *photoResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
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

// A convenience function for testing
// There must be a better way of doing this

func (c *Client) OverrideBaseURL(url string) {
	c.URL = url
}
