package act

import (
	"blue-goblin/internal/location"
	"errors"
)

type Act struct {
	Id        string
	Locations map[string]location.Location
}

func (act *Act) AddLocation(newLocation location.Location) error {
	if _, exists := act.Locations[newLocation.Id]; exists {
		return errors.New("Location with ID " + newLocation.Id + " already exists.")
	}
	act.Locations[newLocation.Id] = newLocation
	return nil
}

func (act *Act) GetLocation(locationId string) *location.Location {
	if loc, exists := act.Locations[locationId]; exists {
		return &loc
	}
	return nil
}

func NewAct(actId string) Act {
	var act = Act{
		Id:        actId,
		Locations: make(map[string]location.Location),
	}
	return act
}

func (act *Act) VerifyLocations() error {
	act.VerifyExits()
	return nil
}

func (act *Act) VerifyExits() error {
	for _, location := range act.Locations {
		for _, exit := range location.Exits {
			if _, exists := act.Locations[exit.Destination]; exists {
				return errors.New("Exit '" + exit.Id + "' in location '" + location.Id + "' points to a non-existent location '" + exit.Destination + "' within act '" + act.Id + "'.")
			}
		}
	}
	return nil
}
