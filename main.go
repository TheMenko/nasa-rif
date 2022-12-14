package main

import (
	"fmt"
	api "github.com/TheMenko/nasa-rif/api"
	cli "github.com/TheMenko/nasa-rif/cli"
	"github.com/perimeterx/marshmallow"
)

func main() {

	flags := cli.Flags{}
	flags.Get()

	marshmallow.EnableCache()

	err := api.Initialize(*flags.Rover, *flags.NumOfDays, *flags.ApiKey)
	if err != nil {
		fmt.Printf("API Failed: %s\n", err.Error())
	}
}
