package api

import (
	"fmt"
	"time"
)

func Initialize(roverName string, days int, apiKey string) error {

	fetcher := PhotoFetcher{}
	manifest := NasaManifest{}
	photoPrinter := PhotoPrinter{}
	var photos = make([]PhotoContainer, 0, days)
	err := manifest.fetchDateLimits(roverName, apiKey)
	if err != nil {
		return fmt.Errorf("NasaManifest error: %w\n", err)
	}

	date := time.Now().Local()

	for i := 0; i < days; i++ {
		photo, err := fetcher.fetch(roverName, apiKey, date.Format("2006-01-02"))
		if err != nil {
			return fmt.Errorf("PhotoFetcher error: %w\n", err)
		}

		photos = append(photos, *photo)
		date = date.AddDate(0, 0, -1)
	}

	photoPrinter.Print(photos)

	return nil
}
