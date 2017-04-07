package marsrover

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	Sols        []Sol  `json:"photos"`
}

type Sol struct {
	Sol         int      `json:"sol"`
	TotalPhotos int      `json:"total_photos"`
	Cameras     []string `json:"cameras"`
}

type solResponse struct {
	Photos []Photo
}

type Photo struct {
	Id        int
	Sol       int
	Camera    Camera
	ImgSrc    string `json:"img_src"`
	EarthDate string `json:"earth_date"`
	Rover     Rover
}

type Rover struct {
	Id          int
	Name        string
	LandingDate string `json:"landing_date"`
	LaunchDate  string `json:"launch_date"`
	Status      string
	MaxSol      int    `json:"max_sol"`
	MaxDate     string `json:"max_date"`
	TotalPhotos int    `json:"total_photos"`
	Cameras     []Camera
}

type Camera struct {
	Id        int    `json:"id, omitempty"`
	ShortName string `json:"name"`
	RoverId   int    `json:"rover_id, omitempty"`
	FullName  string `json:"full_name"`
}

func NewClient(key string, u string) *RoverClient {
	if key == "" {
		key = "DEMO_KEY"
	}
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

func (c *RoverClient) GetImagesBySol(rover string, sol int) ([]Photo, error) {
	url := fmt.Sprintf(c.URL+"/rovers/%s/photos?sol=%d&api_key=%s", rover, sol, c.Key)
	log.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data solResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data.Photos, nil
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
