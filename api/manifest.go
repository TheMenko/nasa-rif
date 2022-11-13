package api

import (
    "github.com/perimeterx/marshmallow"
    "io/ioutil"
    "net/http"
)

type NasaManifest struct {
    Manifest struct {
        MaxDate string      `json:"max_date"`
        LandingDate string  `json:"landing_date"`
    } `json:"photo_manifest"`
}

func (stimer *NasaManifest) fetchDateLimits(rover string, apiKey string) error {

    url := "https://api.nasa.gov/mars-photos/api/v1/manifests/"+ rover +"/?api_key=" + apiKey

    response, err := http.Get(url)
    if err != nil {
        return err
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return err
    }

    _, err = marshmallow.Unmarshal(responseData, stimer)
    if err != nil {
        return err
    }

    return nil
}
