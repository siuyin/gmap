// Package bike converts Singapore Land Transport Authority Bicycle Parking responses to geojson
package bike

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/siuyin/dflt"
)

type Response struct {
	URL   string    `json:"odata.metadata"`
	Value []Parking `json:"value"`
}

type Parking struct {
	Description      string  `json:"Description"`
	Latitude         float64 `json:"Latitude"`
	Longitude        float64 `json:"Longitude"`
	RackType         string  `json:"RackType"`
	RackCount        int     `json:"RackCount"`
	ShelterIndicator string  `json:"ShelterIndicator"`
}

func load(r io.Reader) []Parking {
	var res Response
	dec := json.NewDecoder(r)
	if err := dec.Decode(&res); err != nil {
		log.Fatal(err)
	}
	return res.Value
}

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string   `json:"type"`
	Geometry   Geometry `json:"geometry"`
	Properties Parking  `json:"properties"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func init() {
	log.Println("lta/bike package loaded with LTA_ACCOUNT_KEY=****")
}

// GeoJSON reads an LTA Bicycle Parking response and returns a json string.
func GeoJSON(r io.Reader) string {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	parkingSpots := load(r)
	fc := FeatureCollection{}
	fc.Type = "FeatureCollection"
	for _, p := range parkingSpots {
		fc.Features = append(fc.Features, Feature{
			Type: "Feature",
			Geometry: Geometry{
				Type:        "Point",
				Coordinates: []float64{p.Longitude, p.Latitude},
			},
			Properties: p,
		})
	}

	if err := enc.Encode(&fc); err != nil {
		log.Fatal(err)
	}
	return b.String()
}

func ParkingSpots(lat, lng float64) string {
	key := dflt.EnvString("LTA_ACCOUNT_KEY", "your-account-key")
	ltaURL := fmt.Sprintf("https://datamall2.mytransport.sg/ltaodataservice/BicycleParkingv2?Lat=%f&Long=%f", lat, lng)
	client := &http.Client{}

	req, err := http.NewRequest("GET", ltaURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("accountKey", key)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	return GeoJSON(res.Body)
}
