package cli

import "flag"

type Flags struct {
	Rover     *string
	NumOfDays *int
	ApiKey    *string
}

func (f *Flags) Get() {
	f.Rover = flag.String("rover", "curiosity", "rover name")
	f.NumOfDays = flag.Int("days", 10, "number of days to fetch image from")
	f.ApiKey = flag.String("api_key", "DEMO_KEY", "api key")
	flag.Parse()
}
