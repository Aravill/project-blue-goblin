package location

import (
	"errors"
)

var ActOneLocations []*Location

func registerLocation(location *Location) error {
	for _, loc := range ActOneLocations {
		if loc.Id == location.Id {
			return errors.New("Location with id " + location.Id + " has already been registered.")
		}
	}
	ActOneLocations = append(ActOneLocations, location)
	return nil
}
