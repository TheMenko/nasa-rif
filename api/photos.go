package api

import (
	"fmt"
	"github.com/perimeterx/marshmallow"
	"io/ioutil"
	"net/http"
)

type OuterJsonObject struct {
	Photos []InnerJsonObject `json:"photos"`
}

type InnerJsonObject struct {
	ImgSrc string `json:"img_src"`
}

type PhotoContainer struct {
	Date   string
	Photos []InnerJsonObject
}

type PhotoFetcher struct{}

func (pf *PhotoFetcher) fetch(rover string, apiKey string, date string) (*PhotoContainer, error) {

	url := "https://api.nasa.gov/mars-photos/api/v1/rovers/" + rover + "/photos?earth_date=" + date + "&api_key=" + apiKey

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	resp := &OuterJsonObject{}

	_, err = marshmallow.Unmarshal(responseData, resp)
	if err != nil {
		return nil, err
	}

	container := &PhotoContainer{}

	container.Photos = resp.Photos
	container.Date = date

	return container, nil
}

type PhotoPrinter struct {
	photos map[string]string
}

func (ps *PhotoPrinter) Print(photoList []PhotoContainer) {

	for _, pl := range photoList {
		fmt.Printf("DATE: %s: [\n", pl.Date)
		for i, p := range pl.Photos {
			if i > 2 {
				break
			}
			fmt.Printf("\t%s\n", p.ImgSrc)
		}
		fmt.Println("]")
	}
}
